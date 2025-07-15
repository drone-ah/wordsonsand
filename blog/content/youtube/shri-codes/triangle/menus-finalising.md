---
title: "shri codes triangle in #1.3 | The Pause Menu (zig + raylib + dvui)"
scheduledDate: "2025-07-17T10:00:00+01:00"
youtubeId: VgI1qpDZ-vc
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
  - uidev
  - ziggameui
chapters:
  - "0:00:00 Intro"
  - "0:00:42 wire up form"
  - "0:17:03 remove logging button"
  - "0:17:43 make config path uneditable"
  - "0:28:20 outro"
links:
  - title: Blog
    url: ../../../posts/triangle/menu/index.md
  - title: Intro to triangle codebase
    url: intro.md
  - title: triangle menu #1.1
    url: ./menus-start.md
  - title: triangle menu #1.2
    url: ./menus-theming.md
---

In this devlog, I wrap up the core functionality of triangle’s pause menu. We
now have a working contact button that opens a browser, a clean way to show the
config path (copyable but not editable), and a fully functional, styled UI built
with Zig, Raylib, and DVUI.

This episode includes:

- Making the config path copyable
- Wiring up the **Contact** button to open a feedback form in the browser
- Dropping the **Logging** toggle (for now) after exploring custom log handlers
- Final cleanup and styling tweaks
- Merge into main — the pause menu is officially done

This completes the pause menu milestone for the Seedling release. Next up: a
friendly notice screen with a version-aware changelog and roadmap.

📁 Game: **triangle** — a synthwave ARPG/factory game in space  
🛠️ Tech: Zig, Raylib, DVUI  
🌱 Release: Seedling milestone
