---
title:
  "shri codes Pong with Zig and Raylib - Part 4: Smarter Collisions & Cleaner
  Code"
publishDate: "2025-07-12T09:45:00+01:00"
youtubeId: rnOdZyiP7Bg
playlist: "PLX7QRDvlHn4J5uVbmVlkyDaGJ8utKR9K_"
tags:
  - ziglang
  - pongclone
  - codingvideo
  - devlog
  - raylib
  - refactoring
  - pong
  - retrodev
  - gamedev
  - retrogames
  - opensource
  - indiegamedev
chapters:
  - "00:00 Intro"
  - "03:41 style: improve file naming"
  - "06:10 refactor: remove default values from non-config structs"
  - "09:36 refactor: switch to rls"
  - "12:24 fix: improve paddle y-axis collisions"
  - "17:15 fix: ensure `isColliding` does only one thing"
  - "18:40 fix: make pong resolution independent"
  - "30:46 Outro"
links:
  - title: Blog
    url: ../../../posts/shri-codes/pong/5-ui.md
  - title: Feedback post on ziggit
    url: https://ziggit.dev/t/building-pong-in-zig-with-raylib-part-1-paddles-and-a-ball/10768/1
  - title: Part 1
    url: ./pong-1.md
  - title: Part 2
    url: ./pong-2.md
  - title: Part 3
    url: ./pong-3.md
  - title: "next: Part 5"
    url: ./pong-5.md
hashes:
  description: e6ea30da35774089ea86759996de49d9
---

In this fourth part of my Pong series in Zig + Raylib, I take a detour from UI
work to fix some important issues raised by the community (shoutout to
`milogreg` at ziggit.dev 🙌).

Here’s what we tackle:

- ✅ File naming: renaming `paddle.z` and `ball.z` to match Zig conventions
- ✅ Removing default initializers from non-config structs
- ✅ Using `rls` (result location syntax) more idiomatically
- ✅ Fixing Y-axis collision detection for paddles
- ✅ Refactoring `isColliding` to do just one thing
- ✅ Encapsulating paddle movement logic (`move_up`, `move_down`)
- ✅ Making the game resolution-independent

This kind of feedback loop is what makes sharing code early so valuable - I
learned a lot, and made the game better. If you're working in Zig or interested
in game dev from scratch with Raylib, why not have a look?

🛠️ Code walkthroughs, fixes, refactors  
🧠 Real-time thinking, not tutorials  
👾 Game dev in Zig with no engine, no fluff

Next time: UI with DVUI - score display, pause menu, and more.
