---
title: 'shri codes Pong with Zig and Raylib #6: Font Size, Collision Bugs, and Refactors'
scheduledDate: "2025-07-15T09:45:00+01:00"
youtubeId: opuoMiMNkok
playlist: "PLX7QRDvlHn4J5uVbmVlkyDaGJ8utKR9K_"
tags:
  - ziglang
  - pongclone
  - codingvideo
  - devlog
  - raylib
  - dvui
  - game-ui
  - refactoring
  - pong
  - retrodev
  - gamedev
  - retrogames
  - opensource
  - indiegamedev
chapters:
  - "00:00 intro"
  - "02:23 feat: make font size for score bigger"
  - "08:35 fix: prevent ball getting stuck in the paddle"
  - "17:48 fix: score on edge of ball going past screen"
  - "20:80 refactor: show score based on paddle play area"
  - "37:31 refactor: move score to `Game` (HEAD -> pong/pt-6)"
  - "43:59 outro"
links:
  - title: Blog
    url: ../../../posts/shri-codes/pong/6-refactor.md
  - title: Feedback post on ziggit
    url: https://ziggit.dev/t/building-pong-in-zig-with-raylib-part-1-paddles-and-a-ball/10768/12
  - title: "prev: "
    url: ./pong-5.md
publishDate: ""
hashes:
  description: 2c27c234965add0b79625b131d4bba99
---

In this episode of building Pong with Zig and Raylib, we fix some key bugs,
improve the score display, and refactor our game logic for cleaner structure.

🛠️ **What we cover:**

- Fixing the paddle collision bug where the ball could get stuck
- Increasing the font size for the score using DVUI
- Improving score logic by moving it out of the paddle struct
- Cleaning up rendering based on paddle play area
- Discussing future plans for a pause menu

Thanks to feedback, especially from milogreg(from ziggit.dev), we’re tightening
up the gameplay and preparing for more features ahead!
