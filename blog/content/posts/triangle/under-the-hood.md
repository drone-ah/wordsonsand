---
title: "Under the Hood of triangle"
date: 2025-06-27T08:01:28+01:00
categories:
  - shri-codes
tags:
  - shri-codes
  - triangle
  - zig
  - raylib
  - dvui
  - gamedev
  - deep-dive
---

In [this Let's Code video](https://youtu.be/8nA-a5Z1IDc) on
[**shri codes**](http://www.youtube.com/@ShriCodesHere), I walk through some of
the early systems in place for [_Triangle_](../../tags/triangle/) - my space
ARPG built in **Zig** with **raylib**.

For the broader story behind the project - what _Triangle_ is, where it's going,
and why I’m making it - check out the
[triangle endeavour](../../endeavours/triangle.md) or the
[devlog archive](../../tags/triangle/). This post is about the _how_.

### 🧩 Core Systems So Far

At this stage, Triangle supports:

- Basic player movement and shooting (Asteroids-style).
- Procedurally generated asteroid fields per _sector_.
- Material drops (iron ore) from destroyed asteroids.
- Automatic recipe unlocking when new items are picked up.
- Smelting system that converts ore to ingots.
- Unlocking and queuing up _constructors_ using the crafting menu.
- A basic UI for crafting and inventory.
- Configurable keyboard mappings using TOML.

That’s the visible layer. Below the surface, the codebase has carved out some
space for more ambitious systems.

### 🗂️ Code Structure Overview

Here’s a quick tour of the key directories and their purpose:

#### `blit/`

Handles all the 2D rendering abstractions - bounding boxes, shapes (including
triangle fans 😉 and circles), and canvas drawing. Wraps `raylib` for better
namespacing and type control.

#### `phys/`

Physics system. `body.zig` implements basic movement, collisions, and inertia.
Still minimal, but usable.

#### `combat/`

Still early - only `bullet` is implemented, and is used to destroy asteroids.
Placeholder for upcoming systems like health, damage, and targeting.

#### `items/`

This module forms the backbone of the crafting system. Items are tied to
recipes, and recipes are triggered via pickups or crafting actions. There's a
lot more to unpack here - likely its own devlog soon.

#### `notifier/`

Manages in-game UI feedback. When you pick something up or unlock a new recipe,
this is what shows it. Notifiers are scoped - one for inventory, one for
crafting - and show up on either side of the screen.

#### `ship/`

This module contains the ship related logic, including movement, and the logic
for constructing smelters and other buildables. Eventually this will handle
assembly chains.

#### `asteroid`

These files should be refactored into its own module, and should contain the
asteroid itself, as well as sector, a chunk of space - asteroids, entities,
bullets, etc. This is the closest thing to a “level” in Triangle. It’s
procedurally generated, and you can eventually travel from one sector to the
next.

#### `diag/`

Code to support diagnostics & logging

#### `game/`

This module should get a bit of refactoring as well, moving some of the files
from the root into its own module

- `config`: Takes care of user control bindings and input overrides via a TOML
  config file.
- `notifier`: supports collecting notifications across multiple channels.
- `save`: Just scaffolding for now - no save/load logic yet, but it’s a
  placeholder to group serializable components together.
- `context`: container struct to support dependency injection.

#### `ui/`

Holds in-game panels like crafting and inventory. Needs cleanup and
consolidation.

### 🧪 What’s Missing

A lot of the systems are stubbed out, but not yet wired in. For instance:

- `power` is empty (future energy system)
- The _canvas_ currently wraps `raylib` without adding much - that will likely
  change as UI complexity grows.

### 🧶 Threads I’ll Pull Next

These are the next things I’ll either build or record:

- Improving the coding structure. I'd like to improve organisation and am
  considering extracting a re-usable library out of it, mainly so I can tinker
  with other smaller game projects.
- A **rudimentary menu system**
- A **build+deploy system** to generate early test builds
- Possibly a **contact/feedback** form directly in-game or on the site

I’m releasing this in what I’m calling a **"Sprout" build** - a playable seed, a
form of extremely early access. Feedback and collaboration is really important
to me and I want to get this out there as soon as there is the tiniest spark of
"life."

If you’d like to follow the journey, you can
[watch the video devlog](https://youtu.be/8nA-a5Z1IDc)
