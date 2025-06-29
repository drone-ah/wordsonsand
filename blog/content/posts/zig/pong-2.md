---
title:
  "Building Pong in Zig with Raylib – Part 2: Ball Movement & Paddle Collisions"
date: 2025-06-29T15:21:35+01:00
categories:
  - shri-codes
tags:
  - shri-codes
  - zig
  - raylib
  - gamedev
  - pong
---

In [Part 1](./pong-1.md) we set up the basics: a window, paddles, and a ball. In
this episode, we go one step further and get the ball moving 😉, add paddle
collisions, and make everything frame-rate independent.

## Ball Movement

The first step was to give the ball some velocity and update its position every
frame.

```zig
self.pos = rl.math.vector2Add(self.pos, self.vel);
```

We also fixed a small oversight: the ball and paddles were being recreated every
frame inside the game loop. Moving their initialization outside meant we could
actually observe state changes between frames.

## Frame-Rate Independence

Raylib provides `GetFrameTime()` which returns the time in seconds since the
last frame. Multiplying the velocity by this `dt` value ensures that the ball
movement stays consistent across different frame rates:

```zig
const vel_this_frame = rl.math.vector2Scale(self.vel, dt);
self.pos = rl.math.vector2Add(self.pos, vel_this_frame);
```

With that, the ball now moves at a steady speed, no matter the frame rate.

## Paddle Collision (X-Axis)

Next up: detecting collisions between the ball and paddles. I considered writing
a standalone collision checker, but ended up keeping the logic within the
`Paddle` struct itself.

To simplify which edge to check (left or right), I added a `which` field to
paddles - an enum with values `left` and `right`. That made the conditional
logic much cleaner.

To debug collisions, I added color switching: when a paddle detects a collision
on the x-axis, it flashes red.

```zig
pub fn isColliding(self: *Paddle, ball: *const Ball) bool {
    // which edge do we need to check
    const crossing_x: bool = switch (self.which) {
        .right => ball.pos.x + ball.r >= self.pos.x,
        .left => ball.pos.x - ball.r <= self.pos.x + size.x,
    };

    self.colour = if (crossing_x) .red else .white;
    return crossing_x;
}
```

## Paddle Collision (Y-Axis)

After confirming horizontal collision detection, I added vertical bounds
checking. This just involved verifying the ball's y-position is within the
paddle’s vertical range.

```zig
pub fn isColliding(self: *Paddle, ball: *const Ball) bool {
    // which edge do we need to check
    const crossing_x: bool = switch (self.which) {
        .right => ball.pos.x + ball.r >= self.pos.x,
        .left => ball.pos.x - ball.r <= self.pos.x + size.x,
    };

    if (!crossing_x) {
        self.colour = .white;
        return false;
    }

    const colliding = ball.pos.y >= self.pos.y and ball.pos.y <= self.pos.y + size.y;

    self.colour = if (colliding) .red else .white;
    return colliding;
}
```

## Bounce Logic

With detection in place, we added bounce logic to the ball. If a collision with
a paddle is detected, we flip the x-component of the velocity vector:

```zig
pub fn checkPaddleCollision(self: *Ball, paddle: *Paddle) void {
    if (paddle.isColliding(self)) {
        self.vel = rl.math.vector2Scale(self.vel, -1);
    }
}
```

## What’s Next

That wraps up part 2. In the next episode, we’ll handle edge collisions (top and
bottom), scoring, and input management.

Thanks for following along!

## Links

- [Watch Video](https://youtu.be/IoOLH1O_a7M)
- [Source Code (at this point)](https://github.com/drone-ah/wordsonsand/tree/de1c67e77812cb200e46ab9003d5ac3f2e2bf6ea/games/pong)
- Prev: [Place Paddles & Ball](./pong-1.md)
- Next: Edge Collisions, Scoring & Inputs
