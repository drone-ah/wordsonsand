---
layout: post
title: Basic Gameplay
date: 2025-05-08 10:07:08
type: post
published: true
status: publish
categories:
  - triangle
tags:
  - zig
  - raylib
meta:
---

The first goal was to get it working as far as the coding challenge itself.

## The Ship

Spawning the ship itself was relatively straightforward. I only needed three
points and `drawTriangle` from `raylib`.

In the coding challenge, the ship was rotated with the keyboard, but I wanted
the ship to point to the mouse so that aiming was straightforward. Rotating to
the mouse was trickier — it involved `atan2` and ChatGPT got me started.

The coding challenge video did not worry about time lapsed between each frame,
but I wanted triangle to be framerate independent. That involved a bit of
jiggery pokery to get working, including determining how much the ship can move
within a time frame.

Moving the ship was a bit easier, but based on the many videos of the Coding
Train that implemented force, velocity and dampening, it was pretty
straightforward. Dampening took a bit of trial and error. ChatGPT sent me down
the garden path initially with an over-complicated formula, but I was able to
simplify it later.

Integrating the physics for linear momentum with the one for rotational momentum
was also quite nice — it meant that the ship could overshoot the aim when
rotating and will correct back.

## The Asteroids

The coding challenge worked with a fixed number of on screen asteroids and they
wrapped around. I needed to expand this to:

- A potentially infinite vertical scroller.
- Collisions between the asteroids (currently the same as from the challenge,
  only checks full circle collision)
- Figure out how to handle asteroids moving out of the screen

The above will be covered in a bit more depth in the next devlog.

## The Camera

I also needed a follow camera. Unlike the original asteroids, the ship can move
up/down and the camera should follow it.

The current code looks something like this:

```zig
/// ship_y: ship's current y position
/// ship_vy: ship's y velocity
/// dt: time elapsed since last update
/// margin: the zone out of which camera is moved
/// speed: maximum camera movement speed
fn updateCameraY(self: *Self, ship_y: f32, ship_vy: f32, dt: f32, margin: f32, speed: f32) void {
    const cam_y = self.camera.target.y;
    const inner_margin = margin * 0.8;

    // Predict future ship position
    // TODO: at high speeds, when the clamp goes off, the camera does a little shuffle
    // It would be nice to fix that
    const screen_height: f32 = @floatFromInt(self.screen.height);
    const predicted_y = ship_y + std.math.clamp(ship_vy * 30.0, -screen_height, screen_height);

    var target_y = self.camera.target.y;
    var target_clr: rl.Color = .red;

    const tmargin = cam_y - inner_margin;
    const bmargin = cam_y + inner_margin;
    if (predicted_y < tmargin) {
        target_y = predicted_y + margin;
        target_clr = .blue;
    } else if (predicted_y > bmargin) {
        target_y = predicted_y - margin;
        target_clr = .green;
    }

    if (target_y > screen_height / 2) target_y = screen_height / 2;
    const t = std.math.clamp(dt * speed, 0.0, 1.0);
    self.camera.target.y = lerp(cam_y, target_y, t);
}
```

It predicts the position of the ship, based on its y velocity. If that's outside
the middle zone, it'll move the camera at up to the maximum speed. There is some
clamping to prevent things flying off at high speeds. It also prevents the
camera from overtaking the ship.

This code is mostly hacked together with some help from ChatGPT. It'll be
cleaned up later, once I have a better indication of possible ship speeds, and
how/if we'll zoom in and out as well.

Another important facet to consider with the camera was what to do with the
edges. Out of the four edges, only one is infinite.

The current solution is that the camera stops tracking when it gets to the
"bottom." It also does not move to the left or right. The ship, on the other
hand is free to move out of the range of the camera. Currently, nothing changes
except that the camera does not follow.

Since the thrust works in the direction of the mouse, it's pretty easy to bring
the ship back on to the screen. This feels like the world is big and still out
there, we just don't track what happens out there.

I considered wrapping around on the left and right, but that felt more like the
ship was trapped in that zone. I want the feeling of
[being trapped in vengeance](/2025/04/26/a-lonely-triangle/#story) to be more
subtle ;)

## Combat

Combat is pretty much a mirror image of what happens in the coding challenge,
though I didn't have some of the helper functions. I learned some math :)

Initially, the update loop only handled asteroid collisions. I added bullets as
a separate field in the update loop. It then checks each bullet with every
asteroid in the active chunks (covered in devlog #2).

If there is collision, the asteroid is split into two, moved apart a bit, and
given opposite linear momentum. The bullet also takes "damage" at this point,
and can be removed. The damage system is designed to support penetration, which
is currently not active.

If the spawned asteroids would be too small, nothing is spawned. Later on, this
would be the trigger to spawn a material drop.

## Zig / Raylib

Learning [zig](https://ziglang.org) and [raylib](https://www.raylib.com/) was a
big part of this. Fortunately, both were a lot of fun to learn and work with.

One thing I found annoying was that the vector functionality in raylib was
scattered around as individual functions instead of on the vector struct. While
this was understandable, with raylib being written in C, I found it a bit
frustrating.

I ended up writing my own Vector struct in zig and the functions that I used as
methods on that struct. It was an opportunity for me to learn some vector math
as well.

I also encapsulated raylib inside a `Canvas` struct. I probably still have some
`rl.` calls other places in the code, but the idea is that a canvas is passed
into any bits of code that needs it. The main help is that it'll convert our
version of `Vector2` to the one that raylib wants.

I am also thinking about how I want input handling to work. I would like to
encapsulate that into a separate struct. Right now, I mainly access raylib
directly.

## Manual testing

While I am a big proponent and fan of Test Driven Development, I was happy
enough with manual testing for triangle. I found joy in seeing it work each time
I manage to get something working.

I do end up writing tests later, particularly when I got to bits of code that
was harder to test manually.

## Feel

triangle already feels fun and light. There are some clunky elements to be
ironed out, but so far, it feels good :)
