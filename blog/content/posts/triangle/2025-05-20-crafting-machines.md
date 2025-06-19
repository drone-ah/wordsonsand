---
categories:
  - triangle
date: "2025-05-20T14:27:00Z"
meta: null
redirect_from: /s/tcrafting
slug: crafting-machines
status: publish
tags:
  - triangle
  - crafting
  - devlog
title: Crafting, Mods, and Machines
---

In this post, I want to explore the crafting system in triangle, especially how
mods, tiers, and machine interactions could create depth without overwhelming
complexity.

Crafting in triangle should be a mix of random number generation (rng) and
deterministic. Crafted items will have a random set of mods. However, like in
[Last Epoch](https://lastepoch.com/), you can extract the mods of the items.
These can then be placed into another item.

It feels like mod crafting would be an interesting way to approach crafting in
general. Unlike in the Last Epoch, the extracted mods aren't just combined back
together.

What if each mod retains its values? What if you could combine two mods together
to upgrade or reroll them?

When you insert a mod into an item, what if it destroyed the mod that you were
replacing?

## Tiers

All materials, items, and mods have tiers. Factories will only process items of
their tier or lower. A Tier 1 smelter cannot process a Tier 2 material and a
Tier 1 constructor cannot craft a Tier 2 item.

I am currently mulling over the idea of having nine tiers.

## Items (a.k.a Buildings? Machines?)

I'm toying with the idea of the tier determining how many mod slots it could
have. A Tier 2 item would have two mod slots and a Tier 5 item would have 5
slots.

This provides a natural power curve, even if the mods themselves do not have
tiers.

All items will also have up to one implicit mod. This mod will be item type
specific. e.g., for smelters, it will be `smelting rate`

From a concept perspective, there are probably buildings, or some kind of
structure that can be slotted on to the ship.

Perhaps buildings are a better name for them than items.

### Slots / Hardpoints

I see two kinds of slots on the ship

- _Interior slot_: Machines like the smelters will go in here
- _Exterior Slot_: Weapons, Armour, Hull Extensions etc. could go on to these.

## Mods

Currently mods also support tiers, but I am wondering whether the power creep
could become exponential, combined with the additional mod slots. There may be
good ways to mitigate it though, which will only become apparent once some of
these mechanics have been implemented.

Another reason to skip tiers on mods is if that means that the player ends up
with far too many mods in their inventory.

There are a couple of key areas in which I am thinking about deviating from
other games (like the Diablo series, PoE 1/2, Last Epoch etc.)

### Local only

Each mod applies only to the item it is on. A firing rate increase on your
smelters won't do anything - since they don't fire. A smelting speed increase on
your weapon will likewise have no effect. It would make sense to extract or
replace the ineffective mods.

If you have two weapons attached, and one of them has a mod that increases
firing rate, that will impact only that weapon.

### No Prefix / Suffix mods

There is no classification of mods as prefix/suffix or limits of each type. You
can have full attack mods on if you like. Considering the previous restriction,
it makes more sense to stack attack on weapons and defense on armour.

## Machines

There are a few factory based machines I have in mind so far.

All these machines will have a rate at which they will complete their work. I
have been pondering whether higher tier machines will complete lower tier work
faster.

## Smelters

The standard refinery. It will convert ore and scrap (obtained when enemy ships
are destroyed, and from scrapping items) into refined materials.

## Constructor

Converts refined materials into items of Tier 1. For Tier 2 onwards, I am
considering taking items from the previous tier plus a refined material of the
new tier. E.g. to craft a Tier 2 factory, it could take 2 x Factory Mk. I + 30
Tier 2 materials (whatever that turns out to be)

This could keep the recipes easy to understand and reason about. It also
provides one use for poorly rolled items.

## Disassemblers

These function similar to the Foundry in Last Epoch, at least with regards to
extracting all the mods from the item.

I haven't decided how deterministic it will be. Will it extract all items or
could some be damaged in the process and maybe turned to scrap?

## Scrappers

These are a bit like the trash bin. If crafting requires lower tier items, these
might not be necessary. TBD!

## Foundry

This would be the machine that could do mod crafting. I am still a little foggy
about how or if this machine could work.

Should mod crafting involve risk? Like combining two mods to maybe create a
higher tier one - or failing and getting scrap?

What about rerolling a mod by sacrificing another mod?

All ideas accepted :)

## Augmentor

The Augmentor could be the final step in the pipeline - letting you insert or
fine-tune mods once you’ve refined or crafted them.

This machine will work instantly so that it doesn't tie up the items currently
in use. It would allow you to modify items up to its own tier.

## triangle

With machines slotted directly onto the ship, triangle becomes a living,
drifting factory - salvaging, upgrading, refining, evolving, fighting, and
surviving!

## Other posts

- [Companion vlog for this post](https://youtu.be/livphL9lOxo)
- [Prev: Materials & Pickups ](2025-05-13-materials.md)
- Next: Coming Soon
