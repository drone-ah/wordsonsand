---
title:
  "Building Pong with Zig and Raylib - Part 4: Smarter Collisions, Cleaner Code"
date: 2025-07-10T15:33:34+01:00
categories:
  - shri-codes
tags:
  - shri-codes
  - zig
  - raylib
  - gamedev
  - pong
---

Change of Plans

I was going to dive into menus and UI, but after sharing the early version of
Pong on [ziggit.dev](https://ziggit.dev/), I got a bunch of helpful feedback.
The feedback was the kind that makes you stop and think, _ah, right... I should
probably fix that before carrying on._

So that's what this episode became: a collection of fixes, tweaks, and small
refactors that clean up the code and align things more closely with how things
_should_ be done in Zig (or at least, better than I had them before).

## 🧼 Naming Matters

I’d originally named my files `paddle.zig` and `ball.zig` - lowercase,
snake-case. I had tried to find the guidelines around this, but turns out I only
got half the story. If a file implicitly defines a struct via top-level fields,
it should be named in PascalCase. So, `Paddle.zig`, not `paddle.zig`.

It’s a small thing, but one that helps be a bit more idiomatic - and avoids
confusion when others are reading it.

(now to make the same change across many more files in
[triangle](../../endeavours/triangle.md))

## 🛠️ Default Field Initializers (and When _Not_ to Use Them)

Another thing I learned: structs that aren’t used as config objects shouldn’t
use default field values. Instead, they should have an `init` constant that
represents their starting state.

I’d missed this distinction, and both of my types were using default values
incorrectly. So I cleaned that up and added the values into the `init` method.
It’s a subtle change, but it keeps config objects and plain data objects
conceptually separate - and makes it clearer which parts of a struct are
supposed to be overridden.

## ✨ RLS, Please

One of the comments suggested I lean more into Zig’s Result Location Syntax -
where you define the type on the left-hand side and let Zig figure out the rest.

I’d been using a mix of styles. Nothing broke, but consistency helps. So I swept
through the code and updated those as well.

```zig
//main.zig
var left_paddle: Paddle = .init(Paddle.size.x * 0.5, .left, screen_height);
var right_paddle: Paddle = .init(screen_width - Paddle.size.x * 1.5, .right, screen_height);
var ball: Ball = .init(.{ .x = screen_width * 0.5, .y = screen_height * 0.5 });

```

## 🎯 Fixing Paddle Collisions on the Y Axis

Now for something more visible: my collision detection logic only checked the
center point of the ball along the y-axis. That meant if the ball clipped the
paddle at the edge, it was sneaking through.

The fix was simple: add/subtract the ball’s radius in the y-axis check. Much
better.

```zig
const colliding = ball.pos.y + ball.r >= self.pos.y and ball.pos.y - ball.r <= self.pos.y + size.y;
```

You can actually see the difference in-game - that satisfying little _thock_ now
triggers when it should, even on corner hits.

(now I just need add some sounds - I'd forgotten about that)

## 🧽 `isColliding` Should Only Collide

Previously, `isColliding` also handled coloring the paddle red when it detected
a hit - a debug leftover that had no place in the final function.

I stripped that out and left `isColliding` to do just one thing: return whether
there was a collision. If I want debug visuals again later, I’ll wrap this in
another function.

```zig
// Paddle.zig
pub fn isColliding(self: *const Paddle, ball: *const Ball) bool {
    // which edge do we need to check
    const crossing_x: bool = switch (self.which) {
        .right => ball.pos.x + ball.r >= self.pos.x,
        .left => ball.pos.x - ball.r <= self.pos.x + size.x,
    };

    if (!crossing_x) {
        return false;
    }

    const colliding = ball.pos.y + ball.r >= self.pos.y and ball.pos.y - ball.r <= self.pos.y + size.y;

    return colliding;
}
```

## 🔀 Movement Logic Encapsulation

Paddle movement was scattered, and the logic for moving up/down lived inside
`main.zig`. I pulled that out into proper `moveUp` and `moveDown` methods on
`Paddle`.

It reads cleaner now:

```zig
//Paddle.zig
pub fn moveUp(self: *Paddle, dt: f32) void {
    self.move(-100, dt);
}

pub fn moveDown(self: *Paddle, dt: f32) void {
    self.move(100, dt);
}
```

```zig
// main.zig
if (rl.isKeyDown(.w)) {
    left_paddle.moveUp(dt);
}

if (rl.isKeyDown(.s)) {
    left_paddle.moveDown(dt);
}

if (rl.isKeyDown(.e)) {
    right_paddle.moveUp(dt);
}

if (rl.isKeyDown(.d)) {
    right_paddle.moveDown(dt);
}

```

…and it keeps the input logic in `main`, but the movement logic in the paddle -
where it belongs.

## 📐 Resolution Independence

Some values were hardcoded (like setting `y = 200` for paddle start position),
while others used `getScreenHeight()` and `getScreenWidth()`. I refactored to
make everything use actual screen dimensions, converting `i32` screen height
values to `f32` where needed.

It was a bit fiddly, but worth it. Now Pong should behave properly regardless of
window size.

## Closing

So, yeah - not the flashiest episode, but a satisfying one. Lots of small things
that feel better now that they’re fixed. And it's a reminder that sharing early
(even rough work) is usually a good idea. You never know what you'll learn.

Next time, we _will_ get into UI. I’m planning to bring in
[DVUI](https://github.com/david-vanderson/dvui) and show a basic score display,
a pause menu, and maybe some options for reset and quit.

Until then, thanks for reading (and watching) - see you in the next one.

## Links

- [Watch Video](../../youtube/shri-codes/pong/pong-4.md)
- [Source Code (at this point)](../../../../games/pong/)
- Prev: [Ball Movement & Paddle Collisions](./pong-3.md)
- Next: Display Score, Menu and other UI.
