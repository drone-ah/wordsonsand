---
title: "Projector: Keep YouTube Descriptions synced"
date: 2025-07-04T16:52:21+01:00
categories:
  - wordsonsand
tags:
  - wordsonsand
  - projector
  - youtube
  - hugo
draft: true
---

In my [previous post](./projector-hugo.md), I used [hugo](https://gohugo.io/) to
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

I used python for [despatches](./despatches.md) and it was the right fit there -
good libraries for BlueSky and Reddit.

However, I did not enjoy the experience:

- `bazel` was a constant struggle
- `poetry` is nice… but still a bit of a nightmare. It just makes the pain more
  structured
- Worst of all: one error caused a _partial success_, which broke the Git commit
  and silently caused a post to be repeated (embarrassing!)

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
YouTube libraries and it somehow seems fitting that hugo is also a go baby.

Even though I’m not wiring the two directly, the ecosystem fit is nice.

## Overall Plan

- Let `hugo` render the YouTube description as plain text
- Traverse the `youtube/*.md` files in the source directory

  - Skip videos that are too old to update (maybe older than 30 days?)
  - Hash the rendered output (title, description, tags, etc.)
  - Compare that hash with the one stored in the frontmatter
  - If it doesn’t match,
    - Update the metadata on YouTube
    - Update the hash

## Validation

One thing worth being careful about is whether the metadata is valid. We do not
want the sync to fail during its scheduled run - when it won't have many choices
on how to resolve it.

In a bid to mitigate this, we'll add a command to validate the source and
rendered files.

The validation would expect the rendered files to be generated as well, which
seems reasonable since hugo is probably running as `hugo serve` while the
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
