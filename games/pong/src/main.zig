// raylib-zig (c) Nikolas Wipper 2023

const std = @import("std");
const rl = @import("raylib");
const Paddle = @import("Paddle.zig");
const Ball = @import("Ball.zig");

const Game = @import("Game.zig");

pub fn main() anyerror!void {
    // Initialization
    //--------------------------------------------------------------------------------------
    rl.initWindow(800, 450, "raylib-zig [core] example - basic window");
    const screen_width: f32 = @floatFromInt(rl.getScreenWidth());
    const screen_height: f32 = @floatFromInt(rl.getScreenHeight());

    defer rl.closeWindow(); // Close window and OpenGL context

    rl.setTargetFPS(60); // Set our game to run at 60 frames-per-second
    //--------------------------------------------------------------------------------------

    var game: Game = .{
        .left_paddle = .init(Paddle.size.x * 0.5, .left, screen_height),
        .right_paddle = .init(screen_width - Paddle.size.x * 1.5, .right, screen_height),
        .ball = .init(.{ .x = screen_width * 0.5, .y = screen_height * 0.5 }),
        .screen_height = screen_height,
        .screen_width = screen_width,
    };

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

        rl.clearBackground(.black);

        game.update(dt);
        game.render();

        //----------------------------------------------------------------------------------
    }
}
