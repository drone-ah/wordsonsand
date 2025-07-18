---
title:
  "Building Pong with Zig and Raylib #6: Font Size, Collision Bugs, and
  Refactors"
publishDate: "2025-07-19T08:00:00+01:00"
categories:
  - shri-codes
tags:
  - shri-codes
  - zig
  - raylib
  - gamedev
  - pong
---

In this one, we’ll tidy up our Pong implementation by addressing three key
issues:

1.  The score font is too small
2.  The ball can get stuck inside a paddle
3.  Trigger score on edge (not the middle) of the ball going past.

We’ll also refactor score tracking to better match game logic structure.

## 🖋️ Make the Score Font More Readable

The score text was a little too small. If you're using `dvui`, here's how to
adjust it:

```zig
const font_size: f32 = 64;
var label_options: dvui.Options = .{
    .color_text = .white,
    .font_style = .title,
};
label_options.font = label_options.fontGet().resize(font_size);
dvui.label(@src(), "{d}", .{score}, label_options);
```

Thanks to
[code sample from milogreg](https://ziggit.dev/t/building-pong-in-zig-with-raylib-part-1-paddles-and-a-ball/10768/12)

You may also need to adjust the width and height of the container if the larger
font gets clipped. In our case, font size of 64 felt about right.

## 🧱 Fix the Ball Getting Stuck in the Paddle

When the ball moves too far in one frame, it can land _inside_ the paddle and
bounce back and forth infinitely.

This trap happens because we just multiply the x-velocity with `-1` to reverse
direction.

One way to fix this is to only trigger the bounce if the ball is moving _toward_
the paddle. For example:

```zig
const crossing_x = switch (self.which) {
    .right => ball.vel.x > 0 and
        ball.pos.x + ball.r >= self.pos.x,
    .left => ball.vel.x < 0 and
        ball.pos.x - ball.r <= self.pos.x + size.x,
};
```

Thanks again to
[code sample from milogreg](https://ziggit.dev/t/building-pong-in-zig-with-raylib-part-1-paddles-and-a-ball/10768/12)

This ensures we don’t apply the bounce logic when the ball is already inside the
paddle.

Another (potential) way to fix this would be to change the collision logic to
fix the x direction based on whether it's the left or right paddle.

## 🧮 Trigger a Score When the Ball’s Edge Crosses the Screen

Previously, we checked whether the ball’s **center** (`ball.x`) crossed the
screen edge. Particularly with the ball being bigger than the paddle, this
caused issues when the top/bottom of the paddle hit the ball.

[Game.zig](../../../../../games/pong/src/Game.zig)

```zig
if (self.ball.pos.x + self.ball.r > self.screen_width) {
    self.left_score += 1;
    self.ball.reset();
}

if (self.ball.pos.x < self.ball.r) {
    self.right_score += 1;
    self.ball.reset();
}
```

This triggers the score as soon as the edge of the ball crosses the screen
bounds.

## 🔄 Refactor: Move Scores Out of the Paddle Struct

Storing the score inside the `Paddle` struct is convenient but semantically odd

- paddles shouldn’t own scores. Instead, let's move them into your `Game`
  struct:

Then, pass the score explicitly to the rendering function.

You can find the updated code in
[Game.zig](../../../../../games/pong/src/Game.zig).

## ✅ Bonus Fix: Make Score Display Use Paddle Play Area

We now compute each paddle’s _play area_ (a `dvui.Rect`) and use it to position
the score label, keeping layout logic more re-usable.

We add a `play_area` field or method to our `Paddle` struct that returns its
side of the screen. This makes the rendering logic clearer and more flexible.

## ⏭️ What’s Next?

Next episode: adding a **pause menu** to Pong, based on what I just built for my
other game, _triangle_. We’ll add basic Resume/Quit options and freeze game
state mid-play.

## Links

- [Watch Video](../../../youtube/shri-codes/pong/pong-6.md)
- [Source Code (at this point)](../../../../games/pong/)
- Prev: [Smarter Collisions & Cleaner Code](./5-ui.md)
- Next: Pause Menu
