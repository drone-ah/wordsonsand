---
categories:
  - triangle
date: "2025-04-26T10:07:08Z"
meta: null
slug: a-lonely-triangle
status: publish
tags:
  - triangle
  - devlog
title: A lone triangle vs the universe
---

I've been unwell for a few months which has left me with a great deal of fatigue
and limited ability to be productive.

One of the activities to pass time while healing was playing a lot of ARPG's
recently, particularly [PoE 2](https://pathofexile2.com/home) and the
[Last Epoch](https://lastepoch.com/). I'd even had a few ideas about how I might
do things a little differently. Since my brain capacity was pretty low, these
games with their repetitive nature and easy dopamine hits were really enjoyable.

I found the gameplay a bit fragmented though, and having to follow guides was
annoying. I had been mulling over the idea of making something which flows, and
supports your decision making. It's not about giving the player answers, but
about providing the player the most relevant information in an easy to digest
format. They still get to make the choices. I struggled to identify what
mattered - guides were easier.

I did wonder if there was an easy way to just say - "use this build from this
guide" so that it would save me from all the tedious clicking :/

All this left me wondering if I could build something that captured the fun
while keeping you in the flow. I even looked at how complicated it would be to
make an isometric game. Too complicated for me was the answer.

Factory games, one of my other favourite genres felt a little too mentally
taxing. Otherwise, I would have dived into
[Space Age](https://factorio.com/buy-space-age).

I have also been bingeing on youtube videos and discovered a fantastic channel -
the [Coding Train](https://www.youtube.com/@TheCodingTrain).

I've been almost religiously watching the
[Coding Challenges](https://www.youtube.com/watch?v=17WoOqgXsRM&list=PLRqwX-V7Uu6ZiZxtDDRCi6uhfTH4FilpH)
playlist and I have been loving it. I love how he makes 2d graphics programming
seems so easy and approachable. I used to be so intimidated by all the math.

There was a slow confidence building in me that maybe - just maybe I could build
a game.

<!--more-->

## Coding Challenge 46

Last night, I was watching the one about remaking
[Asteroids](https://www.youtube.com/watch?v=hacZU523FyM). He made it look so
easy and fun. How he puts those asteroids together felt so clever - neat! I
could do that. Actually, I could make that really interesting.

For a little while now, I've been wanting to try out one of the challenges, tied
in with playing with [zig](https://ziglang.org/), which has also been on my
radar. I've largely worked with higher level languages like Java (and golang
more recently). I've always been fascinated by C, and C++, but never quite got
into it. Zig represented a more fun and modern version of C that I could enjoy.
I also really liked how explicit and simple zig feels.

It feels like a perfect opportunity to feed two birds with one scone. I could
try and replicate what he did in that video in zig.

A little bit of research on graphics/game libraries in zig revealed
[raylib](https://www.raylib.com/) to be the best candidate for my needs. While a
C library, it has zig bindings, there is plenty of documentation, albeit not
necessarily in zig - but that was ok.

I could see it in my mind - a field of asteroids, _pew pew pew_, asteroids break
apart. Destroyed asteroids would drop resources. Oh you could have factories on
the ship and build your armoury - weapons, armour, scanners - everything. That
would be an interesting twist - you don't pick up loot - you build it.

Each item would still have randomised mods, but you could extract them like in
Last Epoch and use them in other ones - I have a lot of ideas.

## Story {#story}

In terms of the story, I've been pondering the destructive nature of capitalism
a lot lately and that felt like a good metaphor to bring in here - it just
connected.

The planet where the triangle belongs has been destroyed by a corporation - not
necessarily maliciously, but in a bid to extract it of all valuable resource.

I have this image in my mind of something called a "Mill" that the planet is put
through. The remnants create an asteroid field in which you find yourself.

The game essentially starts off with seeking vengeance. Victory - at least over
the corporation - is not possible. You can get stronger, much stronger, and you
can get to a place where you can beat the waves with relative ease - but the
waves never end!

I have been pondering a softer ending, a choice that the player can make, when
they've had enough - but I'm reluctant to give that away.

On pondering this. I realised this could also be a quiet metaphor for
depression. Having been through it in my life, I could feel a connection. What
if **the triangle is indestructible**?

Actually, that would improve the flow - no game over screen, and no respawn. All
the machinery and installations on the ship could be destroyed, but the core of
the ship - that always remains. You can always keep fighting.

## Gameplay

### Scavenge

Starting off in an asteroid field, you're able to scavenge materials. The ship
starts off with a smelter, a factory, and some basic power. You can refine
materials and construct buildings to get stronger.

Each sector will get more difficult. I thought about being able to travel
vertically and horizontally, but in the end, I decided to make it a vertical
scroller. A two dimensional map might be more interesting, but it also brings
annoying choices, particularly in terms of wondering whether you should have
gone another way.

By putting the game on some form of rails, you have a smaller context to think
about - and you can focus on fun. I am reminded of
[Raptor: Call of the Shadows](https://en.wikipedia.org/wiki/Raptor:_Call_of_the_Shadows),
a game that I loved in the 90's.

Raptor had shops, but I want to do a super simplified factory aspect. Every
upgrade should feel earned — constructed, not found. I want to feel that
progress is made entirely by the actions of the player.

There will still be randomisation, but it will be limited to the mods that are
spawned on the items that are constructed. The player gets to decide what items
to build.

I want the crafting to be as deterministic as possible, but while also keeping
it interesting.

### Figuring it out

There are no (traditional) tutorials or help sections. This decision will put a
lot of pressure on getting the UI/UX intuitive and easy to use. However, it will
also help maintain flow and keep every bit of progress feeling earned.

Items will have descriptions and details so that the user can make choices, but
it is about experimentation and exploration. I would also like a component,
something like a computer that can be installed. The idea is that this can do
analysis of items and mods and give you details of clear and easy to understand
benefits and disadvantages. Not just what attributes change, but how it will
impact gameplay. I can't be perfect, but it must be accurate, and be able to
provide enough information to be able to decide whether it's worth a try.

The intention is that this computer can act a bit like the guide that you would
otherwise google.

Each sector increases difficulty, but is infinite in itself. After traveling for
a bit, you will be offered an exit. If you keep going, exits will be offered at
regular intervals - but you can just keep going.

You are dropped into a universe and every step is up to you to figure it out.

## Aesthetic

I can hear the synthwave tracks that play in the background already (it could be
because I've been listening to a lot of synthwave recently). It feels right and
thematic to the original Asteroids.

The graphics, at least initially is going to be simple and shape based, again
thematic. I've been pondering whether and how the graphics would get updated as
you get through the sectors. The start of it is going to be black and white with
little hints of colour. As you progress, you'll get more and more colour and the
vibrancy increases, sometimes in leaps and bounds. It could be a metaphor for
the journey of recovery.

The sound effects should also have a strong 80s vibe. I want it to feel like
you've been transported to the 80s but if it happened in 2025.

## triangle

As a name, at least for the time being, I've settled on **triangle**. It is a
passion project - I have a vision and a strong idea of how I want to feel
playing it.

At its heart, triangle is about (re)building - a game where every piece of
progress feels personal.

I would like to share it with the world and am looking for feedback, thoughts
ideas and supporters.

If the idea or any of the concepts resonate with you, I’d love to hear from you.

## Other Posts

- [Next: Basic Gameplay](/2025/05/08/basic-gameplay/)

## Updates

2025-05-31: Added [devlog #0 on youtube](https://youtu.be/8pBPQbJtIJk)\
2025-05-21: Added an [itch.io page](https://droneah.itch.io/triangle)
