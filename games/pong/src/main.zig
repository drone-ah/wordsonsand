// raylib-zig (c) Nikolas Wipper 2023

const std = @import("std");
const rl = @import("raylib");
const Paddle = @import("Paddle.zig");
const Ball = @import("Ball.zig");
const dvui = @import("dvui");

const Game = @import("Game.zig");

const RaylibBackend = dvui.backend;
comptime {
    std.debug.assert(@hasDecl(RaylibBackend, "RaylibBackend"));
}
const ray = RaylibBackend.c;

pub fn main() anyerror!void {
    // Initialization
    //--------------------------------------------------------------------------------------
    rl.initWindow(800, 450, "raylib-zig [core] example - basic window");
    const screen_width: f32 = @floatFromInt(rl.getScreenWidth());
    const screen_height: f32 = @floatFromInt(rl.getScreenHeight());

    defer rl.closeWindow(); // Close window and OpenGL context

    rl.setTargetFPS(60); // Set our game to run at 60 frames-per-second
    //--------------------------------------------------------------------------------------
    var gpa: std.heap.GeneralPurposeAllocator(.{}) = .init;
    const allocator = gpa.allocator();
    defer {
        const deinit_status = gpa.deinit();
        //fail test; can't try in defer as defer is executed after we return
        if (deinit_status == .leak) @panic("MEM LEAK!!!");
    }

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

    var game: Game = .init(screen_width, screen_height);

    // Main game loop
    while (!rl.windowShouldClose()) { // Detect window close button or ESC key
        // Update
        //----------------------------------------------------------------------------------
        // TODO: Update your variables here
        //----------------------------------------------------------------------------------

        // Draw
        //----------------------------------------------------------------------------------

        const dt = rl.getFrameTime();

        rl.beginDrawing();
        defer rl.endDrawing();

        // marks the beginning of a frame for dvui, can call dvui functions after this
        try win.begin(std.time.nanoTimestamp());

        // send all Raylib events to dvui for processing
        _ = try backend.addAllEvents(&win);

        rl.clearBackground(.black);

        game.update(dt);
        game.render();

        //----------------------------------------------------------------------------------
        _ = try win.end(.{});

        // cursor management
        if (win.cursorRequestedFloating()) |cursor| {
            // cursor is over floating window, dvui sets it
            backend.setCursor(cursor);
        } else {
            // cursor should be handled by application
            backend.setCursor(.arrow);
        }
    }
}
