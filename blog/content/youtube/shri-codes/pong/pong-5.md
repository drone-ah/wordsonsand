---
title: "shri codes Pong with Zig and Raylib #5: Show the Scores with dvui"
scheduledDate: "2025-07-15T09:45:00+01:00"
youtubeId: laPvy6CVUx0
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
  - "01:25 build: add `dvui` dependency"
  - "07:35 refactor: extract `Game` struct"
  - "18:00 refactor: improve encapsulation of `Game` init logic"
  - "22:50 feat: integrate `dvui` into main loop"
  - "29:15 feat: show scores"
  - "38:59 refactor: improve code reuse & readability"
  - "45:54 prepare PR"
  - "50:23 review feedback"
  - "54:36 git fixup"
  - "56:13 outro"
links:
  - title: Blog
    url: ../../../posts/shri-codes/pong/5-ui.md
  - title: Feedback post on ziggit
    url: https://ziggit.dev/t/building-pong-in-zig-with-raylib-part-1-paddles-and-a-ball/10768/12
  - title: "prev: "
    url: ./pong-4.md
  - title: "next: "
    url: ./pong-6.md
hashes:
  description: ff5de54b61d9e4473df33d9205893c29
---

In this episode, we're diving into **UI work in Zig** using **DVUI**, adding
**score display** to our little Pong clone.

We’ve already got scoring logic working behind the scenes — now it's time to
show it on screen. Along the way, I:

- Integrate DVUI into the game loop
- Set up score rendering for both paddles
- Refactor the game into its own struct with proper init and render/update
  methods
- Tidy up the build setup and improve code reuse
- Discuss feedback from Ziggit (shoutout to milo_greg!) around allocators,
  rendering quirks, and edge collision bugs

This is **part 5** of the Pong series, and while I had hopes of doing more, as
always… things take longer than you think.

I’ll cover the remaining polish and next steps — like handling sound, fixing
paddle-bounce bugs, and maybe input config — in a future video.

🔧 Built with **Zig**, **Raylib**, and a good dose of elbow grease.

Thanks for watching — feedback always welcome!
