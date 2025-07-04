---
title: "Projector: Keep YouTube Descriptions synced"
date: 2025-07-03T16:52:21+01:00
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
