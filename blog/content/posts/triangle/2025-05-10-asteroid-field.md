---
categories:
  - triangle
date: "2025-05-10T10:07:08Z"
meta: null
redirect_from:
  - /s/tlog2
  - /2025/05/08/asteroid-field/
slug: asteroid-field
status: publish
tags:
  - triangle
  - quadtree
  - procedural generation
  - devlog
title: Procedural Asteroid Field Generation
---

In this post, I am going to cover procedural asteroid field generation. At a
high level, I wanted:

- An asteroid field that feels infinite
- Natural-looking distribution and density
- A safe starting area for the player

## Spawning Asteroids

I explored strategies for spawning multiple asteroids on screen without overlaps
and avoiding the player’s starting zone. Initially, I considered a brute-force
collision check for each spawn candidate, but quickly realized it wouldn’t scale
well with many asteroids.

I decided to split the screen into a grid. If the maximum radius of an asteroid
is `r`, then each grid would be `2r x 2r`, and would be able to accommodate
asteroids at their full size. The asteroid size is determined randomly, with a
minimum size. It also offsets the center of the x by a random amount, up to a
maximum of `r`.

<!--more-->

![The grid visualised](/assets/2025/05/asteroid-map.png "The asteroid grid")

In theory, this means that if two adjacent cells spawn the full size asteroid
and offset in the "wrong" direction, they could start off overlapping. I've not
seen this happen yet, likely because the full size asteroids are not that
common.

The grid has an additional offset in the x axis for each row. With this offset,
even if every asteroid was placed at the x center of each grid, they would not
look uniform.

Finally, for each grid cell, there is some further randomisation that determines
whether the asteroid will get spawned at all.

In hindsight - there might be more randomisation than we need. It works well
though.

The asteroids are also given a low starting velocity, both linearly and
rotationally.

Even if asteroids do overlap on spawn, the physics will correct it through the
collision detection.

## Spawning the Ship

The code for this part is a little icky right now. It checks the y position of
the ship and avoids spawning asteroids in any grid cell that overlaps the
starting location.

Of course, this should only matter for the asteroids spawned on the starting
screen. The rest of the asteroids don't need this check. It works just fine
though - for now :)

## Making it "Infinite"

I pondered the best way to make it feel infinite without risk of slowdowns and
issues. A quadtree was an obvious answer. However, it felt too complicated since
I'd never written one before, and I was worried about how I would move each
asteroid between each section as it moved. (Spoiler alert: I should have just
used quadtrees, and intend to move to it later.)

In the meantime, I wanted to build something which felt simpler, which was
basically to create chunks of asteroids, with each chunk being a little larger
than the size of the screen.

The system could then, based on the current y position, track just the current,
previous, and next chunks, and process only these asteroids.

When I got it working, it was pretty ok. The main problem I had now was the
asteroids leaving the screen and never coming back.

After I'd implemented this, I ended up watching
[Coding Challenge #98: Quadtree](https://www.youtube.com/watch?v=OJxEcs0w_kE),
and boy did I feel silly about being scared about quadtrees. They were promptly
demystified and felt like a much simpler solution than what I have here - so
I'll add that to the list to change.

## Asteroids Drifting Offscreen

One of the problems was that the asteroids would drift offscreen. Because I
wasn't updating all the asteroids, they'd never come back on to the screen. In
fact, even if I updated _all_ the asteroids, the screen would eventually clear.
I believe this was because there is a lot of offscreen space (mainly on the left
and right of the screen) that the asteroids could fly off to.

I tried various ways of wrapping the asteroids around after they hit a few
hundred pixels past the edge. I didn't want it to look like a wraparound, and
pretending like the screen was about twice the width helped.

However, the problem was now with the asteroids slinking out the bottom of the
screen (and the top too), and there wasn't an easy way to tackle those.

In the end, I ended up
[using an attractor](https://www.youtube.com/watch?v=OAcXnzRNiCY), which I
placed on the center of each chunk.

This worked surprisingly well that I didn't even need the horizontal wraparound.

I had also ended up making the `Chunk` struct a little heavy, and I ended up
refactoring it. While refactoring, I overoptimized it to only update asteroids
on the screen. This "fix" meant that asteroids off-screen aren't attracted to
the center of the chunk anymore, and the screen ends up clearing again.
Something to fix later - maybe when I implement the quadtree.

## Debug Panel

There were various points when I was trying to track down issues that I reverted
to my habit of printing stuff out - but that was a nightmare because it was
printing the same thing in each frame.

Eventually, I got into the habit of displaying stuff on screen and drawing
different colours etc. I then wanted a way to see things easier and toggle some
of the debug features. To do this, I needed a GUI.

I looked at [raygui](https://github.com/raysan5/raygui), which looked good but
probably too basic for me. I don't enjoy GUI work and I figured something a
little more fleshed out would help.

I had already looked at [dvui](https://github.com/david-vanderson/dvui) before
and it looked like it could be a good candidate.

Integrating it into _triangle_ was pretty straightforward - at least once I
figured out how to add it to `build.zig`.

## Next Steps

Right now, the field feels big enough (around 15 chunks), but not truly
infinite. I'm thinking of implementing a circular buffer - say, 8 or 10 chunks -
where old ones are recycled as the player moves forward. If a player backtracks,
most asteroids will have drifted off anyway, and replacements will blend in just
fine.

I’ll probably tackle that - and maybe finally add a quadtree - soon.

**Next up:** crafting systems and how I’m handling material drops and recipes.

## Other posts

- [Companion vlog for this post](https://youtu.be/RXcBDC8Ki1w)
- [Basic Gameplay](./2025-05-08-basic-gameplay.md)
- [Next: Materials & Pickups](2025-05-13-materials.md)
