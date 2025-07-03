---
title:
  "shri codes Pong with Zig and Raylib - Part 2: Ball Movement & Paddle
  Collision"
publishDate: 2025-07-08T09:45:00+01:00
youtubeId: "IoOLH1O_a7M"
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
  - 0:50 Move the ball
  - 3:48 Framerate Independence
  - 7:25 Paddle Collision X
  - 23:26 Paddle Collision Y
  - 25:30 Bug find & fix
  - 28:37 Paddle Collision Y (cont...)
  - 30:35 Ball Bounce
  - 41:08 Outro
links:
  - title: Blog
    url: ../../../posts/zig/pong-2.md
  - title: Prev
    url: ./pong-1.md
  - title: Next
    url: ./pong-3.md
outputs: ["plain"]
_build:
  list: never
  render: always
  publishResources: false
sitemap: false
---

**In this episode of Building Pong in Zig with Raylib**, we get the ball moving
and dive into collision detection! 🚀

We pick up from where we left off — with two paddles and a stationary ball — and
start adding real gameplay:

- Add ball movement using velocity
- Make movement frame rate independent
- Implement basic x and y axis collision with paddles
- Get the ball bouncing off paddles using direction reversal
- Debug with paddle color changes to show collision points
- Tidy up edge detection for accurate hitboxes

This episode is all about laying the foundation for proper gameplay physics.
It’s unscripted, exploratory, and focuses on getting a working loop — no
perfection, just progress.
