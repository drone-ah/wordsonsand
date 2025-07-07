---
title:
  "Building Pong in Zig with Raylib – Part 3: Edge Collisions, Scoring & Player
  Input"
publishDate: 2025-07-07T12:01:00+01:00
categories:
  - shri-codes
tags:
  - shri-codes
  - zig
  - raylib
  - gamedev
  - pong
---

In this part, we round out the core gameplay of our Pong clone in Zig using
raylib. By the end, we've got a working game - with a bouncing ball, paddle
controls, and basic scoring.

## 🧱 Edge Collisions

We already had paddle collision working, but we needed the ball to bounce off
the top and bottom edges of the screen. That meant checking if the ball's
y-position was outside the visible range and inverting its vertical velocity
accordingly:

```zig
pub fn checkEdgeCollisions(self: *Ball, screen_height: f32) void {
    if (self.pos.y < self.r or self.pos.y > screen_height - self.r) {
        self.vel.y *= -1;
    }
}
```

There was a small detour where I realised `screen_height` wasn’t giving the
expected value - using `GetScreenHeight()` directly from raylib helped resolve
that.

## 🏁 Scoring System

Once edge collisions were in place, it was time to detect goals. If the ball
passed the left or right edge of the screen, the opposing player got a point. We
also need a `reset()` method to the `Ball` struct to return it to the centre
after each goal.

```zig
pub fn reset(self: *Ball) void {
    self.pos = self.home;
    self.vel = .{ .x = 250, .y = -50 };
}
```

For now, let's print the scores via `std.debug.print`, but this sets the stage
for UI integration.

```zig
if (ball.pos.x > screenWidth) {
    left_paddle.score += 1;
    std.debug.print("scores: l: {d}, r: {d}", .{ left_paddle.score, right_paddle.score });
    ball.reset();
}

if (ball.pos.x < 0) {
    right_paddle.score += 1;
    std.debug.print("scores: l: {d}, r: {d}", .{ left_paddle.score, right_paddle.score });
    ball.reset();
}
```

## ⌨️ Player Input

Finally, we wired up keyboard input to allow paddle movement. We used
`IsKeyDown` from raylib, and made sure movement was frame-rate independent by
scaling it with `dt`:

```zig
pub fn move(self: *Paddle, y: f32, dt: f32) void {
    self.pos.y += y * dt;
}
```

Left paddle uses `W/S`, right paddle uses `E/D`. Simple and responsive.

```zig
if (rl.isKeyDown(.w)) {
    left_paddle.move(-100, dt);
}

if (rl.isKeyDown(.s)) {
    left_paddle.move(100, dt);
}

if (rl.isKeyDown(.e)) {
    right_paddle.move(-100, dt);
}

if (rl.isKeyDown(.d)) {
    right_paddle.move(100, dt);
}
```

## ✅ What's Working

By the end of this episode, we’ve got:

- Ball bounces off top and bottom edges
- Scoring when the ball passes a paddle
- Ball reset after each goal
- Both paddles are fully controllable

All that’s missing now is a visible score, a win condition, and maybe a simple
menu.

## Links

- [Watch Video](https://youtu.be/IoOLH1O_a7M)
- [Source Code (at this point)](https://github.com/drone-ah/wordsonsand/tree/shri-codes/pong/part-3/games/pong)
- Prev: [Ball Movement & Paddle Collisions](./pong-2.md)
- Next: Code Improvements from Review
