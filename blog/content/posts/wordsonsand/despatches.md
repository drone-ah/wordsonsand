---
title: "Automated Posting to BlueSky & Reddit"
date: 2025-07-01T10:09:47+01:00
categories:
  - wordsonsand
tags:
  - wordsonsand
  - despatches
  - automation
  - hugo
  - reddit
  - bluesky
  - python
---

I tend to be pretty impatient and when I'm doing something, I want to just
finish it off. Unfortunately, the world works better for me when I work to its
schedule.

Every time I finish a video for [shri codes](../../endeavours/shri-codes.md),
while I am still in the zone, I want to post to all the places (YouTube, BlueSky
and Reddit). However, this is usually the worst time to share these if I want to
get some decent traffic and raise awareness.

I've been remembering to post on the relevant days at reasonable times, but this
process is annoying at best, interrupts flow and takes up cognitive load.

I wanted to automate it. I've got to say that automating these two seemingly
simple tasks were rife with unexpected complexity.

My first challenge was trying to get `rule_python` to work, which, in the end, I
did not succeed and gave up.

Getting `pylyzer` to work in neovim was also a challenge - another one that I
gave up on.

<!-- more -->

I briefly gave up on python altogether and went with go, and I made stellar
progress until I got to the bit about actually posting to `BlueSky` - ChatGPT
had (once again) lied to me (shame on me not verifying their claims). The
library it wanted me to use was a hallucination, and did not exist. I then
realised that there was no real library for reddit integration either.

Back to python, and trusty `poetry` to see me through.

## Scheduled Elements

There are four elements to getting scheduling to work

### YouTube Scheduling

This part was the easiest. The platform is kind enough to provide an option to
schedule release of videos, and we'll use that!

### Scheduled Publish for Blog

`hugo` supports this out of the box. The bigger challenge was how to get GitHub
Actions to regenerate the site when relevant. In the end, I identified the
window during the week when I want to be publishing.

10am - 4pm Mon - Fri seemed like a decent slot. GitHub Actions though does not
support summer time. I opted for 10am - 3pm, which seemed the better option.

My GitHub action for publishing takes one minute to execute. If I run the action
every 30 minutes, for three hours five days a week:

`2 * 5 * 5 * 4 = 200`

[.github/workflows/hugo.yaml](https://github.com/drone-ah/wordsonsand/blob/main/.github/workflows/hugo.yaml)

```yaml
on:
  schedule:
    - cron: "*/30 12-15 * * 1-5"
  # Runs on pushes targeting the default branch
  push:
    branches:
      - main
```

I will need to run a second one for the despatches (below) as well, which would
mean around 400 minutes each month - while there are no limits for public
repos - it felt a little abusive to run it every minute.

Once this has been running safely for a while, I'll consider bumping the
cadence.

#### Cron is Unreliable on GitHub Actions

After I got this all ready with the two workflows set up to run on GitHub
Actions, I waited, and waited, and nothing happened.

[GA schedule doc](https://docs.github.com/en/actions/reference/events-that-trigger-workflows#schedule)
states:

> The schedule event can be delayed during periods of high loads of GitHub
> Actions workflow runs. High load times include the start of every hour. If the
> load is sufficiently high enough, some queued jobs may be dropped. To decrease
> the chance of delay, schedule your workflow to run at a different time of the
> hour.

The
[upptime post about GitHub Actions schedule not working](https://upptime.js.org/blog/2021/01/22/github-actions-schedule-not-working/)
includes some suggested workarounds, namely:

- [IFTTT](https://ifttt.com/) - seems to be limited to a maximum of hourly
- [Google Cloud Scheduler](https://cloud.google.com/scheduler/docs/) - could be
  a good solution but a bit of a sledgehammer
- [Cronhub](https://cronhub.io/) - starts at $19/mo

I also discovered:

- [cron-job.org](https://cron-job.org/) - haven't tried this yet, but looks
  viable

I was going to try out cron-job.org when ChatGPT suggested a simpler
alternative - a simple workflow that only triggered the relevant workflows.

According to ChatGPT, the more complex a workflow, the more likely it is to be
dropped. It makes sense, of course, and while I wasn't fully convinced, I
decided to
[give it a go](https://github.com/drone-ah/wordsonsand/blob/main/.github/workflows/cron.yaml).

It's only been 10 minutes, but it has completed one run already - which is
promising, but the original run also ran once.

I'll have to keep an eye on the reliability of this.

#### Switched to `cron-job.org`

While the above strategy was OK, I wanted something more reliable, so I switched
to [cron-job.org](https://console.cron-job.org)

I created a new access token, restricted to the repo and with two additional
permissions:

- actions: read & write
- contents: read (to read the workflow file, ChatGPT suggests)

I then set up a http call to:

`https://api.github.com/repos/<gh-username>/<repo-name>/actions/workflows/<workflow-filename>/dispatches`

- `<gh-username>`: use your github username from the url
- `<repo-name>`: name of your repo, again from the url
- `<workflow-filename>`: The filename of the workflow you want to trigger

To triger my hugo run, I used:
`https://api.github.com/repos/drone-ah/wordsonsand/actions/workflows/hugo.yaml/dispatches`

Under advanced, I set the following Headers:

- `Accept`: `application/vnd.github+json`
- `Authorization`: `token <personal-access-token>`
- `Content-Type`: `application/json`
- `User-Agent`: `cronjob`

Set `Request method` to `POST`

`Request body`:

```json
{
  "ref": "main"
}
```

### BlueSky

This one - posting to BlueSky was far more complicated than I anticipated. All
the complexity was around its requirement to separate the post out into facets.
I recognise and value the semantic content such a process would output. However,
I could not find an algorithm or any details on how to extract the facets from
some text - e.g. markdown.

I referenced some code from a couple of sources for a stopgap solution to
address urls and hashtags.

And then, I found [blueskysocial](https://github.com/dmoggles/blueskysocial)

### Reddit

You first need to [register an app](https://www.reddit.com/prefs/apps/) on
reddit, from a page I don't seem to be able to get from anywhere except a direct
link.

Once I registered a `personal script`, which will let any of the developers
registered on that client to post, I got to try and login and was faced with:

`prawcore.exceptions.OAuthException: invalid_grant error processing request`

#### Red Herrings

I tried directly with curl:

```bash
curl -u "$CLIENT_ID:$CLIENT_SECRET" \
  -d "grant_type=password&username=$USERNAME&password=$PASSWORD" \
  -A "$APP_NAME" \
  https://www.reddit.com/api/v1/access_token
```

and I got a similar error:

`{"error": "invalid_grant"}`

After stumbling around for a while, verifying and re-verifying the credentials,
I also set up a brand new account using password auth (mine was originally
oauth). It also returned the same error.

Some resources that I followed:

- [redditdev](https://www.reddit.com/r/redditdev)
- [OAuth2 Quick Start Example](https://github.com/reddit/reddit/wiki/OAuth2-Quick-Start-Example)

While I was lookin around, I noticed in tiny little letters on the page to
[register an app](https://www.reddit.com/prefs/apps/), when you create a new
app:

> By creating an app, you agree to Reddit's Developer Terms and Data Api Terms.
> **You must also
> [register to use the API](https://www.reddit.com/r/reddit.com/wiki/api/#wiki_read_the_full_api_terms_and_sign_up_for_usage).**

(Emphasis mine)

I followed the instructions on that page, which felt more like red tape, but
easy enough for an app that is only intended to post on a schedule.

Alas, this too did not help!

#### Final Solution

Perhaps not surprisingly, the final solution was to not use the password, but
get a refresh token instead.

You can do this manually on the browser. Start by going to the following URL:

`https://www.reddit.com/api/v1/authorize?client_id=YOUR_CLIENT_ID&response_type=code&state=xyz&redirect_uri=http://localhost&duration=permanent&scope=identity,submit,read`

- `YOUR_CLIENT_ID`: Replace this with the client id from your reddit app
- `redirect_uri`: This value has (`http://localhost` in the example) has to
  match the `redirect_uri` setting in your app
- `scope`: Update to the scopes you are looking for.
  [/api/vi/scopes](https://www.reddit.com/api/v1/scopes) will return the list of
  valid scopes and their descriptions.
- `state`: can be any value. It's supposed to match in the next step

The browser will then ask your permission (of the scopes you defined). If you
approve, the browser will redirect to localhost (or whatever url you define for
the redirect above).

This redirect will likely fail, but that's ok. There is one parameter in the URL
that you are looking for - `code`

In my case, I got something like:

`http://localhost:8080/?state=xyz&code=RilF7XDhRTr7o7B-iov2gpdDgum5pA#_`

(don't worry - that code isn't the actual one)

You want to take the code, but without the `#_` at the end and substitute it in
the following:

```bash
curl -X POST -A "despatcher" --user "$CLIENT_ID:$CLIENT_SECRET" \
  --data "grant_type=authorization_code&code=$CODE&redirect_uri=http://localhost:8080" \
  https://www.reddit.com/api/v1/access_token
```

- `CLIENT_ID`: The app id from your app settings page (again)
- `CLIENT_SECRET`: The secret from you app settings page
- `CODE`: The code that was in the URL above
- `redirect_uri`: exactly the same `redirect_uri` as above, and in the app
  settings

```json
{
  "access_token": "<access-token>",
  "token_type": "bearer",
  "expires_in": 86400,
  "refresh_token": "<refresh_token>",
  "scope": "read submit identity"
}
```

- `access_token`: You can use this to auth, but not so useful for long term use
  as it will expire
- `refresh_token`: more useful as it can be used to get a new access token. Pass
  to `praw`

[tools/despatcher/despatch.py](https://github.com/drone-ah/wordsonsand/tree/main/tools/despatcher/despatch.py)

```python
client_id = os.environ.get("APP_REDDIT_CLIENT_ID")
client_secret = os.environ.get("APP_REDDIT_CLIENT_SECRET")
refresh_token = os.environ.get("APP_REDDIT_REFRESH_TOKEN")

reddit = praw.Reddit(
    client_id=client_id,
    client_secret=client_secret,
    refresh_token=refresh_token,
    user_agent="despatcher",
)

print(reddit.user.me())

```

## Posting & Tracking

Once a post has been submitted, it is important that we log it somehow.
Otherwise, we'll end up posting it again (and again (and again)).

The cleanest solution I could think of was to update the markdown file, then
commit and push the change. This will also help to keep a log of it.

[.github/workflows/despatcher.yaml](https://github.com/drone-ah/wordsonsand/blob/main/.github/workflows/despatcher.yaml)

```yaml
- name: Commit and push if changed
  run: |
    git config user.name "drone-ah bot"
    git config user.email "github.actions@drone-ah.com"

    if ! git diff --quiet; then
      git add -u
      git commit -m "auto: log post submissions"
      git push
    else
      echo "No changes to commit"
    fi
```

## Partial Successes

Now, I thought I'd covered the worst offenders for risk of repeated posting, but
I'd missed one case.

What happens when something gets posted, then the script errors?

Well, the git commit won't happen - and sadly this happened to my. My apologies
to the nice folks at [r/selfhosted](https://www.reddit.com/r/selfhosted/) who
got a handful of my posts about automated posting - eek :(

Embarrassment aside, it identified at least one fix - probably two. Extra
embarrassing because something like this has happened to me before - many years
ago - but you live!

The first update is to get GitHub Actions to carry on even if there is an error:

[.github/workflows/despatcher.yaml](https://github.com/drone-ah/wordsonsand/blob/main/.github/workflows/despatcher.yaml)

```yaml
- name: Run despatcher script
  working-directory: tools/despatcher
  continue-on-error: true
  run: poetry run ./despatch.py ../../despatches/
```

The second fix it to catch any errors from the dispatchers.

[tools/despatcher/despatch.py](https://github.com/drone-ah/wordsonsand/tree/main/tools/despatcher/despatch.py)

```python
try:
    ptype = p.get("type")
    if ptype == "bluesky":
        url = post_bluesky(p)

    if ptype == "reddit":
        url = post_reddit(p)
except Exception as e:
    print(f"[ERROR] Failed to post to {ptype} for {path}: {e}")
    continue  # Skip to the next file

```

## Wrap Up

In the end, what I thought was a two hour job took me two days, but such is the
life of a software engineer (probably everyone).

I am looking forward to see how it works, and a little scared if it'll go off
and do random things in my name - but we'll see

## Links

- [Code Repo](../../../../tools/despatcher/)

## Updates

- 2025-07-08: Switch to `cron-job.org`
- 2025-07-02: Add note about GA cron unreliability
- 2025-07-02: Add details of handling partial success
