---
title: 'shri codes Pong with Zig and Raylib - Part 3: Edge Collisions, Scoring & Player Input'
publishDate: "2025-07-05T09:45:00+01:00"
youtubeId: https://youtu.be/9TmoiLjtWrg
playlist: "PLX7QRDvlHn4J5uVbmVlkyDaGJ8utKR9K_"
tags:
  - ziglang
  - pongclone
  - codingvideo
  - devlog
  - raylib
  - pong
  - retrodev
  - gamedev
  - retrogames
  - opensource
  - indiegamedev
chapters:
  - 0:00 Intro
  - 1:16 Edge Collisions
  - 15:42 Scoring
  - 25:13 Move Paddles
  - 32:38 Outro
links:
  - title: Blog
    url: ../../../posts/zig/pong-2.md
  - title: Part 1
    url: ./pong-1.md
  - title: Part 2
    url: ./pong-2.md
hashes:
  description: de1cd101e372aef8630dc2861435cd7f
---

**In this episode of Building Pong in Zig with Raylib**, we make the game
actually _playable_! 🕹️

We wrap up core gameplay mechanics by adding:

- **Edge collisions** - the ball now bounces off the top and bottom of the
  screen
- **Scoring logic** - players score when the ball goes past their opponent's
  side
- **Player input** - paddles can now be moved using `W`/`S` and `E`/`D` keys
- Debug lines and position logging to help track down weird bugs
- Ball reset and direction handling after goals

By the end of this video, we’ve got a functioning, two-player Pong clone with
scoring and movement. Next up: showing the score, adding a menu, and
implementing win conditions.
