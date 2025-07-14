---
title: "Building Pong with Zig and Raylib #5: Show Score with dvui"
publishDate: "2025-07-15T08:00:00+01:00"
categories:
  - shri-codes
tags:
  - shri-codes
  - zig
  - raylib
  - gamedev
  - pong
---

In this episode, I finally add a score display to Pong using DVUI, a native Zig
UI framework. The scoring logic was already in place - now it's time to show it
on screen.

## Extract out `Game` struct

One of the structural changes I wanted to make to tidy up the code was to pull
out the game logic into its own struct. This change helps to declutter the
`main.zig` file, leading the way to add in the dvui scaffolding.

This refactor pulls the update and render logic into a `Game` struct, which
helps declutter `main.zig` and sets up a better foundation for UI work.

[Game.zig](../../../../../games/pong/src/Game.zig)

```zig
pub fn update(self: *Game, dt: f32) void {
    self.ball.checkEdgeCollisions(self.screen_height);
    self.ball.update(dt);
    self.ball.checkPaddleCollision(&self.left_paddle);
    self.ball.checkPaddleCollision(&self.right_paddle);
    if (self.ball.pos.x > self.screen_width) {
        self.left_paddle.score += 1;
        std.debug.print("scores: l: {d}, r: {d}\n", .{ self.left_paddle.score, self.right_paddle.score });
        self.ball.reset();
    }

    if (self.ball.pos.x < 0) {
        self.right_paddle.score += 1;
        std.debug.print("scores: l: {d}, r: {d}\n", .{ self.left_paddle.score, self.right_paddle.score });
        self.ball.reset();
    }

    if (rl.isKeyDown(.w)) {
        self.left_paddle.moveUp(dt);
    }

    if (rl.isKeyDown(.s)) {
        self.left_paddle.moveDown(dt);
    }

    if (rl.isKeyDown(.e)) {
        self.right_paddle.moveUp(dt);
    }

    if (rl.isKeyDown(.d)) {
        self.right_paddle.moveDown(dt);
    }
}

pub fn render(self: *const Game) void {
    self.left_paddle.render();
    self.right_paddle.render();
    self.ball.render();

    showScore(self.screen_width * 0.25, self.left_paddle.score);
    showScore(self.screen_width * 0.75, self.right_paddle.score);
}

```

## Add dvui dependency

Let's fetch the dependency, adding it to `build.zig.zon`:

`zig fetch --save git+https://github.com/david-vanderson/dvui.git`

Then, in `build.zig`, we also need to declare it:

```zig
const dvui_dep = b.dependency("dvui", .{
    .target = target,
    .optimize = optimize,
});

const dvui = dvui_dep.module("dvui_raylib");
```

and then add it as a dependency:

```zig
exe.root_module.addImport("dvui", dvui);
```

## Add `dvui` to game loop

We also need to initialise dvui in the main loop.

[main.zig](../../../../../games/pong/src/main.zig)

### Import

```zig
const dvui = @import("dvui");

const RaylibBackend = dvui.backend;
comptime {
    std.debug.assert(@hasDecl(RaylibBackend, "RaylibBackend"));
}
const ray = RaylibBackend.c;
```

### Initialise

```zig
var gpa: std.heap.GeneralPurposeAllocator(.{}) = .init;
const allocator = gpa.allocator();
defer _ = gpa.deinit();

//--------------------------------------------------------------------------
// init Raylib backend
// init() means the app owns the window (and must call CloseWindow itself)
var backend = RaylibBackend.init(allocator);
defer backend.deinit();
backend.log_events = true;

// init dvui Window (maps onto a single OS window)
// OS window is managed by raylib, not dvui
var win = try dvui.Window.init(@src(), allocator, backend.backend(), .{});
defer win.deinit();
```

### Pre-render

```zig
// marks the beginning of a frame for dvui, can call dvui functions after this
try win.begin(std.time.nanoTimestamp());

// send all Raylib events to dvui for processing
_ = try backend.addAllEvents(&win);
```

### Post-render

```zig
_ = try win.end(.{});

// cursor management
if (win.cursorRequestedFloating()) |cursor| {
    // cursor is over floating window, dvui sets it
    backend.setCursor(cursor);
} else {
    // cursor should be handled by application
    backend.setCursor(.arrow);
}
```

## Show Score

We can now show the score using `dvui.label` with positioning hardcoded to about
25% and 75% across the screen. There may be a better way to position it but it
works well enough for now.

I had to generate an id for the label so that they were unique. I generated it
the x-position using `@intFromFloat()`.

[Game.zig](../../../../../games/pong/src/Game.zig)

```zig
pub fn render(self: *const Game) void {
    self.left_paddle.render();
    self.right_paddle.render();
    self.ball.render();

    showScore(self.screen_width * 0.25, self.left_paddle.score);
    showScore(self.screen_width * 0.75, self.right_paddle.score);
}

fn showScore(xpos: f32, score: u8) void {
    const id: usize = @intFromFloat(xpos);
    var right = dvui.box(@src(), .horizontal, .{ .rect = .{ .x = xpos, .y = 50, .w = 50, .h = 50 }, .id_extra = id });
    defer right.deinit();

    dvui.label(@src(), "{d}", .{score}, .{ .color_text = .white, .font_style = .title });
}
```

## Closing

As always, things took a bit longer than expected, but by the end:

- The game's structure is cleaner
- DVUI is wired up properly
- Scores now show up on screen

Shoutout to [milo_greg](https://ziggit.dev/u/milogreg/summary) on
[ziggit.dev](https://ziggit.dev/) for loads of really valuable feedback and
tips. That kind of thoughtful review really helps.

## Links

- [Watch Video](../../../youtube/shri-codes/pong/pong-5.md)
- [Source Code (at this point)](../../../../games/pong/)
- Prev: [Smarter Collisions & Cleaner Code](./4-refactor.md)
- Next: More UI
