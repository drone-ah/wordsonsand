---
layout: post
title: Materials
date: 2025-05-13 12:07:00
type: post
published: true
status: publish
categories:
  - triangle
tags:
  - crafting
  - devlog
meta:
redirect_from: /s/tlog3
---

In triangle, everything should feel earned - so all items are crafted, not
found. I am also a big fan of factory / automation games. How much of a factory
sim can we squeeze into an ARPG (or on to a triangle for that matter)? Let's
find out!

## Drops

As a starting point, when asteroids are destroyed, they will drop materials.
I've also been pondering whether materials could drop when they split. It makes
logical sense, but I want to get a better sense of pacing before diving into
that question.

Currently, only iron drops. I started off by finding a sprite sheet with
reasonable looking sprites and then started integrating them in before I decided
to just stick with shapes — so iron drops are little blue circles.

```zig
var shape = Shape.initCircle(7);
shape.move(pos);
shape.colour = .blue;
```

PS: you haven't seen the `Shape` code, but you get the gist.

These drops are managed by a `MaterialField` that keeps track of all the dropped
materials. They are given a low linear velocity, and no rotational velocity
(it's a circle, so you wouldn't see it anyway).

They do not interact with collisions to avoid them being bumped off the screen
or you having to chase them down.

At some point, they could also be optimised to only update if they are on
screen. Assuming the player will pick up most of them, this may not be
necessary.

## Pickup

I also "installed an attractor" on the ship such that if it gets close enough to
a material, it'll pull it in. When it's close enough, it'll be "picked up."

```zig
pub fn pickup(self: *Ship, material: *Material.Drop) bool {
    const magnet_sq = self.magnet_range * self.magnet_range;
    const pickup_sq = self.pickup_range * self.pickup_range;

    const dist = self.pos().sub(material.body.pos());
    const dist_sq = dist.magSquared();

    if (dist_sq <= pickup_sq) {
        self.save.inventory.addMatDrop(material);
        return true;
    }

    if (dist_sq <= magnet_sq) {
        // pull it towards us
        material.body.applyForce(dist.scale(self.magnet_force).sub(material.body.lvel));
    }

    return false;
}
```

It is a physics interaction, so it is possible to slingshot the material and
have it fly off the screen (it's happened to me). It feels like a sun
interaction and you are in control of it, so I am inclined to leave that in.

Once it's close enough, it's picked up, and added to an inventory.

The attractor also feels like an upgrade that can be crafted later.

## Smelting

### Crafting Panel

For smelting, I started by sketching out a straightforward UI. I'll admit up
front that I don't enjoy UI work, but in earnest I started on it, and I got as
far as displaying the panel itself and I got tremendously bored.

![Crafting Panel Sketch](/assets/2025/05/craft-ui-sketch.png "Crafting Panel Sketch")

I decided to make material refinement prioritisation automatic, at least for the
time being. It would always prioritise the most valuable material that is in the
inventory.

Actually, that'll remove some factory micromanagement without removing the
feeling of agency from the player.

It'll mainly be a problem if the player has run out of lower tier materials and
they have a large amount of higher tier items in the inventory. I figure let's
wait until we have some playtesting before we worry about it. The system
currently only drops one material and it can be refined automatically.

### Refining

All materials and items will have a tier, starting with iron at 1. To be able to
smelt a material, you'll need a smelter at a minimum of that tier. For example
to craft iron, you need a minimum tier 1 smelter (which is fine since 1 is the
minimum tier anyway). If you wanted to smelt a tier 3 material, you'll need a
minimum of tier 3 smelter.

The ship starts with a Tier 1 smelter and will always have a tier 1 smelter.
This smelter cannot be modified, removed, or destroyed. The same is true of some
other factories, which will be covered in the next post.

The first iteration of refining took the total capacity of all smelters at each
tier, and would convert a portion of the materials from `ore` to `ingot`. It
used floating points to track progress. If it would take four seconds to convert
one ore to ingot, over one second, it'd transfer 0.25 from ore to ingot.

The display only showed integers, so there was a bit of `floor` and `ceil` to be
able to show consistent numbers.

I didn't like this way of doing things. It felt icky — hacky — but it was simple
enough and it works, for now.

## Other posts

- [Companion vlog for this post](https://youtu.be/8ct9aWNj3Zk)
- [Prev: Procedural Asteroid Fields]({% post_url triangle/2025-05-10-asteroid-field %})
- Next: [Refineries, Constructors and other other factory types in a
  world without
  conveyor belts.]({% post_url triangle/2025-05-20-crafting-machines %})
