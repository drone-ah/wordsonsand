---
title: 'shri codes triangle in #1.1 | The Pause Menu (zig + raylib + dvui)'
publishDate: "2025-07-09T20:53:55+01:00"
youtubeId: ETrRyTCmcPs
playlist: PLX7QRDvlHn4KlzfvZ66GPKQl7HEamwM-v
categoryId: 20
pinnedComment: ""
tags:
  - gamedev
  - ziglang
  - ziggamedev
  - indiegame
  - indiegamedev
  - devlog
  - raylib
  - trianglegame
  - dvui
  - ziggameui
chapters:
  - "0:0:00 Intro"
  - "0:0:44 Planning"
  - "0:15:11 feat: q to quit"
  - "0:17:04 I got distracted"
  - "0:18:18 feat: q to quit: cont..."
  - "0:22:18 feat: close panels on esc"
  - "0:31:22 refactor: move menu logic to main"
  - "0:41:55 feat: pause game on menu display"
  - "0:53:08 feat: resume button"
  - "1:06:24 feat: quit button"
  - "1:11:05 Outro"
links:
  - title: Blog
    url: ../../../posts/triangle/menu.md
  - title: Intro to triangle codebase
    url: intro.md
  - title: dvui 0.15.x issues
    url: https://github.com/david-vanderson/dvui/issues/382
hashes:
  description: d278ff680dda485ce869e068628de669
---

In this episode, I work on building a simple pause menu for triangle, my ARPG
factory-building game set in space - inspired by asteroids and shaped by
recovery.

I sketch out what the menu should look like, think through some basic UX (where
it should sit, what it should include), and start wiring it up using DVUI and
Raylib in Zig. There’s some back and forth with version issues - Zig 0.15 has
some nice improvements, but DVUI isn’t quite ready yet - so I revert to 0.14 and
clean up the code to match.

I get the basics working: Escape to open, Resume and Quit buttons, and proper
panel-closing logic. Next time, I’ll work on improving the layout and theming.

If you’re building your own UI in Zig, trying DVUI, or just want to follow the
journey of triangle, hope this is useful - or at least mildly interesting.
Questions and thoughts welcome as always.
