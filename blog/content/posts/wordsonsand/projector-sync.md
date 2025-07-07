---
title: "Projector: Keep YouTube Descriptions synced"
date: 2025-07-07T20:08:21+01:00
categories:
  - wordsonsand
tags:
  - wordsonsand
  - projector
  - youtube
  - hugo
  - golang
  - inscribe
  - oauth
  - automation
---

In my [previous post](./projector-Hugo.md), I used [hugo](https://gohugo.io/) to
generate correctly linked, always up to date descriptions for my YouTube Videos.

But if I'm generating the descriptions automatically... I'm hardly going to be
excited about copying and pasting them into YouTube - right? right!

Automating this process brings up a few design choices.

## Planning

### Which language

There were a few contenders, and here's how I thought them through:

#### Zig

I’m currently learning [Zig](https://ziglang.org/), and I love using it for my
game development. But it doesn’t yet have mature libraries for working with the
YouTube Data API - and I don’t feel like writing one. So, sadly, Zig’s out for
this one.

#### Python

I used Python for [despatches](./despatches.md) and it was the right fit there -
good libraries for BlueSky and Reddit.

However, I did not enjoy the experience:

- `bazel` was a constant struggle
- `poetry` is nice… but still a bit of a nightmare. It just makes the pain more
  structured
- Worst of all: blusky failed after reddit succeeded caused a _partial success_,
  which broke the Git commit and silently caused a post to be repeated
  (embarrassing!)

  That kind of problem _can_ happen in Go (nil pointer), though it wouldn’t in
  Zig. But at least with Go, most handleable errors _stay_ errors — they don’t
  crash the whole tool.

<!-- more -->

#### java

Sure, I could do this in Java - but I really don’t want to mess with the JVM.
And more importantly, I’m doing this for fun. Java doesn’t feel like that
anymore.

#### golang

Not quite my favourite any more, but still a close second. It's _fast_, has
YouTube libraries and it somehow seems fitting that Hugo is also a go baby.

Even though I’m not wiring the two directly, the ecosystem fit is nice.

## Overall Plan

- Let `Hugo` render the YouTube description as plain text
- Traverse the `youtube/*.md` files in the source directory

  - Skip videos that are too old to update (maybe older than 30 days?)
  - Hash the rendered output (title, description, tags, etc.)
  - Compare that hash with the one stored in the frontmatter
  - If it doesn’t match,
    - Update the metadata on YouTube
    - Update the hash
  - commit and push any updates (should be only hash changes)

## Validation

One thing worth being careful about is whether the metadata is valid. We do not
want the sync to fail during its scheduled run - when it won't have many choices
on how to resolve it.

In a bid to mitigate this, we'll add a command to validate the source and
rendered files.

The validation would expect the rendered files to be generated as well, which
seems reasonable since Hugo is probably running as `hugo serve` while the
content files are being updated.

```go
func validate(sourcePath string, renderedPath string) error {
	targetSourceDir, err := getTargetDir(sourcePath)
	if err != nil {
		return err
	}

	targetRenderedDir, err := getTargetDir(renderedPath)
	if err != nil {
		return nil
	}

	videos, err := findRecentVideos(targetSourceDir)
	for _, video := range videos {
		_, err := video.getDescription(targetRenderedDir)
		if err != nil {
			slog.Warn("unable to find rendered file", "file", video.renderedPath)
		}
	}
	return nil
}
```

The validate function will retrieve the relevant files and check that there is a
corresponding rendered description.

If it errors in that process, we know that it would error out in the sync.

We can't catch errors around the API though at this stage, and that's
unavoidable.

## Sync

### Hashing the Description

This part was surprisingly easy:

```go
bdesc, err := video.getDescription(targetRenderedDir)
if err != nil {
    slog.Warn("unable to find rendered file", "file", video.renderedPath)
}

// We want to hash the contents of description
// Check with the hash in the metadata to see if it matches
hash := md5.Sum(bdesc)
strHash := hex.EncodeToString(hash[:])
```

The challenge was trying to write the updated yaml frontmatter back. I was using
the `adrg/frontmatter` library to read the frontmatter, but it does not support
writing it back.

### Detour: Write a small frontmatter Library

I took a little detour to build
[inscribe, a little frontmatter library that supports reading and writing back in yaml](../golang/inscribe.md).

## Auth

We need the YouTube Client to have an OAuth Token, which we can retrieve by:

- [Create a new OAuth Client](https://console.cloud.google.com/auth/clients) -
  Type of desktop is probably the easiest
- add your user account to
  [test users](https://console.cloud.google.com/auth/audience):
- go to
  https://accounts.google.com/o/oauth2/v2/auth?client_id=YOUR_CLIENT_ID&redirect_uri=urn:ietf:wg:oauth:2.0:oob&response_type=code&scope=https://www.googleapis.com/auth/youtube
  - Remember to substitute your actual client_id
  - Add any other scopes you might want
  - Go through the flow steps - it'll warn you that the app is unreleased, which
    is expected
- Take the code that it provides
- Call the following curl command

```bash
curl -X POST https://oauth2.googleapis.com/token \
  -d client_id=YOUR_CLIENT_ID \
  -d client_secret=YOUR_CLIENT_SECRET \
  -d code=PASTE_THE_CODE_HERE \
  -d grant_type=authorization_code \
  -d redirect_uri=urn:ietf:wg:oauth:2.0:oob
```

You should finally get something like:

```json
{
  "access_token": "ya29...",
  "expires_in": 3599,
  "refresh_token": "1//0g...",
  "scope": "https://www.googleapis.com/auth/youtube",
  "token_type": "Bearer"
}
```

The `refresh_token` is what you want to save / use as the `access_token` will
expire (after an hour in this example).

The authentication was a bit more involved with a refresh token, but the
`oauth2` library helps us out:

```go
func NewYouTube(ClientId string, ClientSecret string, RefreshToken string) (YouTube, error) {

	conf := &oauth2.Config{
		ClientID:     ClientId,
		ClientSecret: ClientSecret,
		Endpoint:     google.Endpoint,
		Scopes:       []string{"https://www.googleapis.com/auth/youtube"},
	}

	// Construct a token from just the refresh token
	token := &oauth2.Token{RefreshToken: RefreshToken}

	ctx := context.Background()

	// Create an authenticated client
	httpClient := conf.Client(ctx, token)

	ytService, err := youtube.NewService(ctx, option.WithHTTPClient(httpClient))
	if err != nil {
		return YouTube{}, err
	}

	return YouTube{
		service: ytService,
	}, nil

}
```

## Updating the description

Setting the description is a little more complicated because you can't set just
the description.

Everything defined in the `VideoSnippet` gets updated.

To support this, what we need to do is get the current snippet for the video,
then update it:

```go
vListCall := ytService.Videos.List([]string{"snippet"})
vListCall = vListCall.Id(videoId)
res, err := vListCall.Do()
if err != nil {
    return err
}

if len(res.Items) != 1 {
    return fmt.Errorf("wrong number of videos returned: %d", len(res.Items))
}

ytVideo := res.Items[0]
ytVideo.Snippet.Description = desc

vUpdateCall := ytService.Videos.Update([]string{"snippet"}, ytVideo)
_, err = vUpdateCall.Do()
```

## GitHub Action

The GitHub Action is fairly straightforward, mostly a copy of the Hugo one,
then:

- Add Bazel
- Run projector sync
- Commit if changed

```yaml
- uses: bazel-contrib/setup-bazel@0.15.0
  with:
    # Avoid downloading Bazel every time.
    bazelisk-cache: true
    # Store build cache per workflow.
    disk-cache: ${{ github.workflow }}
    # Share repository cache between workflows.
    repository-cache: true
- name: Run projector sync
  env:
    GOOGLE_CLIENT_ID: ${{ secrets.PROJECTOR_GOOGLE_CLIENT_ID }}
    GOOGLE_CLIENT_SECRET: ${{ secrets.PROJECTOR_GOOGLE_CLIENT_SECRET }}
    GOOGLE_REFRESH_TOKEN: ${{ secrets.PROJECTOR_GOOGLE_REFRESH_TOKEN }}
  continue-on-error: true
  run:
    bazel run //tools/projector:projector -- sync -source blog/content/youtube
    -rendered blog/public/youtube
- name: Commit and push if changed
  run: |
    git config user.name "drone-ah bot"
    git config user.email "github.actions@drone-ah.com"

    if ! git diff --quiet; then
      git add -u
      git commit -m "auto: log youtube updates"
      git push
    else
      echo "No changes to commit"
    fi
```

We also needed to upgrade one permission - `contents`

```yaml
permissions:
  contents: write
```

## Conclusion

The Google/YouTube documentation was the hardest part here in that it was pretty
obtuse and hard to understand.

Writing a little frontmatter library was unexpected, but while it took a little
time, was straightforward.

Once I got a handle on that, the rest of it was pretty straightforward, partly
because I was reusing parts from before.

## Links

- [Part 1: Outputting YouTube Descriptions from Hugo](./projector-Hugo.md)
- [inscribe: simple frontmatter yaml library](../golang/inscribe.md)
