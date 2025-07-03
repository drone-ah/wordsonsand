---
title: "Building Pong in Zig with Raylib – Part 1: Setup, Paddles, and Ball"
description:
  "A devlog-style walkthrough of setting up a Pong clone using Zig and raylib,
  covering setup and basic rendering."
categories:
  - shri-codes
tags:
  - shri-codes
  - zig
  - raylib
  - gamedev
  - pong
date: 2025-06-28T20:38:29+01:00
---

Before continuing development on Triangle - my larger, arcade ARPG, factory
game, I wanted to take a step back and build something small and familiar. Pong
felt like the perfect choice: quick to prototype, easy to understand, and a good
warm-up before diving deeper into raylib again.

This post goes alongside [my video on YouTube](https://youtu.be/ICq2D_na6zc),
and walks through the early steps of the project: setting up the game, getting
something on screen, and implementing the paddles and the ball.

## Why Pong?

Sometimes, before getting deeper into a project, it helps to do something simple
just to get your hands moving again. I hadn’t written Zig in a couple of weeks,
and wanted a fast feedback loop to:

- Get back into using `raylib-zig`
- Have a bit of fun
- Remind myself why I made certain decisions in the first place

## Project Setup

I'm using [`raylib-zig`](https://github.com/Not-Nik/raylib-zig) for this. You
can scaffold a new project with:

```bash
./project_setup.sh <project-name>
```

You could also use `zig init` and then add the dependency in manually if you
prefer.

## Drawing the Playground

The Pong "arena" is simple:

- Two paddles
- A ball (currently static)
- A center dividing line
- Placeholder for scores

## Implementing Paddles

The paddles have:

- A position (top-left corner)
- A fixed size
- A simple render function

```zig
const std = @import("std");
const rl = @import("raylib");

const Paddle = @This();

pos: rl.Vector2,

pub fn init(x: f32) Paddle {
    return .{ .pos = .{
        .x = x,
        .y = 200,
    } };
}

pub const size = rl.Vector2{ .x = 25, .y = 100 };

pub fn render(self: Paddle) void {
    rl.drawRectangleV(self.pos, size, .white);
}
```

It might have been nice to be able to have the paddle calculate more of its own
values, but the more I think about it, the more it makes sense to keep the logic
in paddle simpler.

## Ball Placeholder

To wrap things up for this session, I added a static ball in the center of the
screen. It has a radius, position, and velocity fields ready to go. Rendering is
straightforward:

```zig
const std = @import("std");
const rl = @import("raylib");

const Ball = @This();

pos: rl.Vector2,
r: f32 = 16,
vel: rl.Vector2 = .{ .x = 0, .y = 0 },

pub fn render(self: Ball) void {
    rl.drawCircleV(self.pos, self.r, .white);
}
```

## What's Next

This was mostly about warming up, but the next episode will tackle:

- Adding velocity to the ball
- Basic collision detection with paddles

I'll try and think about simplifying paddle logic, especially the awkward
symmetry between left and right.

You can find the full source code
[on GitHub](https://github.com/drone-ah/wordsonsand/tree/main/games/pong).

See you in part 2!

## Links

- [YouTube Video](../../youtube/shri-codes/pong/pong-1.md)
- [Full Source Code (at this point)](https://github.com/drone-ah/wordsonsand/tree/shri-codes/pong/part-1/games/pong)
- Next: [Ball Movement & Paddle Collisions](./pong-2.md)
