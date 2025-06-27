---
title: "init & deinit in Zig"
date: 2025-04-27T17:01:57+01:00
categories:
  - learning-zig
tags:
  - learning-zig
  - zig
  - quote
  - memory-management
  - nostalgia
  - learning-to-code
---

When I was ten, my grandmother passed away. As was custom where we lived, my
family moved into her house - a full household with uncles and aunts. I stayed
there for about nine months.

My youngest uncle was a Computing Studies teacher at a local college. He also
taught private classes at home. I couldn’t get enough.

I don’t know what most ten-year-olds dream about, but my dream was to learn C. I
saved ₹300 - a lot of money in 1994 - to buy a book called _Encyclopedia C_. I
read it cover to cover, understood maybe 10% of it, and reread it years later to
pick up more.

But I never actually programmed in C.

<!--more-->

## The Long Detour

Life took me through a range of other languages instead: Prolog (strangely, what
my school machines had), Visual Basic, ASP, PHP, Java, Python, JavaScript, Go,
Rust. I tinkered with C++ when messing with game engines, but never quite got
around to C.

Some of the early fears carried through - memory management was intimidating:
`malloc`, `free`, and in C++, `new` and `delete`.

Over time, I began to overcome the fear of systems languages - first through
Java, then Go. But I never revisited C, and I never quite got over the fear of
manual memory management.

In hindsight, these concepts weren’t really designed for a ten-year-old to
grasp. It makes sense that they felt out of reach.

## Allocator Confusion

Zig had me curious for a while, and some recent health issues gave me the space
to explore it. I’ve been building a small game in Zig, and it’s been feeding
that original desire to learn C - just with a bit more clarity.

I was still intimidated by memory allocation though, and went as far as I could
without using an allocator.

Eventually, I had to use one. I understood the convention of `init` and
`deinit`, and the idea of allocators in general.

But I was confused about how `deinit` worked in the context of `ArenaAllocator`.

> A little learning is a dangerous thing
>
> -- Alexander Pope

If the arena frees memory for you, wouldn’t calling `deinit` on individual
objects inside it risk double-freeing?

If I skip calling `deinit` on a struct that normally needs it, will I miss other
cleanup tasks?

And does this mean the arena allocator isn’t a true drop-in replacement if I
have to skip `deinit`?

At the time, the answer wasn’t clear - and I couldn’t find documentation that
resolved it.

## Clarity

Thankfully, the
[super friendly folk at Ziggit](https://ziggit.dev/t/deinit-and-arena-allocator/9856)
helped clarify things.

### `ArenaAllocator` handles it

`ArenaAllocator` _is_ a drop-in replacement for other allocators. If you
`deinit` objects using memory from the arena, and those objects try to free that
memory, the operation is effectively (though not exactly) a no-op. There’s no
risk of double-free.

> It’s always safe to free / destroy memory from an arena, as long as you treat
> it as freed of course. But it won’t actually be released until the arena is
> deinit’ed.
>
> -- [mnemnion](https://ziggit.dev/u/mnemnion/summary)

Even if some memory is freed manually before the arena is cleared, it’s handled
safely.

I was overthinking it.

### `deinit` should still be called

`deinit` exists for cleanup logic - not just memory. It should always be called
where appropriate, regardless of the allocator. That part doesn’t change.

### Further learning

Zig also encourages a light touch. One of the impacts of this is that the object
itself should not hold on to the allocator in the `init` and use it in `deinit`.
It could, but the better way would be to accept the `allocator` in both the
`init` and the `deinit`.

This convention is further reinforced by
[zig deprecating the `managed` variants of collections](https://ziglang.org/download/0.14.0/release-notes.html#Embracing-Unmanaged-Style-Containers)

## Thanks

In many ways, this was about cleaning up more than just memory.

Special thanks to the helpful folks at [ziggit.dev](https://ziggit.dev/).

## References

- [Full discussion on ziggit.dev](https://ziggit.dev/t/deinit-and-arena-allocator/9856)
- [Implementation of `ArenaAllocator`](https://github.com/ziglang/zig/blob/d92649da80a526f2e2b2f220c05b81becf4fa627/lib/std/heap/arena_allocator.zig#L253-L267)
