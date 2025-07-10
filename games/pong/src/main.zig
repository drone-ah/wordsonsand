// raylib-zig (c) Nikolas Wipper 2023

const std = @import("std");
const rl = @import("raylib");
const Paddle = @import("Paddle.zig");
const Ball = @import("Ball.zig");

pub fn main() anyerror!void {
    // Initialization
    //--------------------------------------------------------------------------------------
    const screenWidth = 800;
    const screenHeight = 450;

    rl.initWindow(screenWidth, screenHeight, "raylib-zig [core] example - basic window");
    defer rl.closeWindow(); // Close window and OpenGL context

    rl.setTargetFPS(60); // Set our game to run at 60 frames-per-second
    //--------------------------------------------------------------------------------------
    var left_paddle: Paddle = .init(Paddle.size.x * 0.5, .left);
    var right_paddle: Paddle = .init(screenWidth - Paddle.size.x * 1.5, .right);
    var ball: Ball = .init(.{ .x = screenWidth * 0.5, .y = screenHeight * 0.5 });

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

        ball.checkEdgeCollisions(@floatFromInt(rl.getScreenHeight()));
        ball.update(dt);
        ball.checkPaddleCollision(&left_paddle);
        ball.checkPaddleCollision(&right_paddle);
        if (ball.pos.x > screenWidth) {
            left_paddle.score += 1;
            std.debug.print("scores: l: {d}, r: {d}", .{ left_paddle.score, right_paddle.score });
            ball.reset();
        }

        if (ball.pos.x < 0) {
            right_paddle.score += 1;
            std.debug.print("scores: l: {d}, r: {d}", .{ left_paddle.score, right_paddle.score });
            ball.reset();
        }

        if (rl.isKeyDown(.w)) {
            left_paddle.move(-100, dt);
        }

        if (rl.isKeyDown(.s)) {
            left_paddle.move(100, dt);
        }

        if (rl.isKeyDown(.e)) {
            right_paddle.move(-100, dt);
        }

        if (rl.isKeyDown(.d)) {
            right_paddle.move(100, dt);
        }

        left_paddle.render();
        right_paddle.render();
        ball.render();

        //----------------------------------------------------------------------------------
    }
}
