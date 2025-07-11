// raylib-zig (c) Nikolas Wipper 2023

const std = @import("std");
const rl = @import("raylib");
const Paddle = @import("Paddle.zig");
const Ball = @import("Ball.zig");

pub fn main() anyerror!void {
    // Initialization
    //--------------------------------------------------------------------------------------
    rl.initWindow(800, 450, "raylib-zig [core] example - basic window");
    const screen_width: f32 = @floatFromInt(rl.getScreenWidth());
    const screen_height: f32 = @floatFromInt(rl.getScreenHeight());

    defer rl.closeWindow(); // Close window and OpenGL context

    rl.setTargetFPS(60); // Set our game to run at 60 frames-per-second
    //--------------------------------------------------------------------------------------

    var left_paddle: Paddle = .init(Paddle.size.x * 0.5, .left, screen_height);
    var right_paddle: Paddle = .init(screen_width - Paddle.size.x * 1.5, .right, screen_height);
    var ball: Ball = .init(.{ .x = screen_width * 0.5, .y = screen_height * 0.5 });

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
        if (ball.pos.x > screen_width) {
            left_paddle.score += 1;
            std.debug.print("scores: l: {d}, r: {d}\n", .{ left_paddle.score, right_paddle.score });
            ball.reset();
        }

        if (ball.pos.x < 0) {
            right_paddle.score += 1;
            std.debug.print("scores: l: {d}, r: {d}\n", .{ left_paddle.score, right_paddle.score });
            ball.reset();
        }

        if (rl.isKeyDown(.w)) {
            left_paddle.moveUp(dt);
        }

        if (rl.isKeyDown(.s)) {
            left_paddle.moveDown(dt);
        }

        if (rl.isKeyDown(.e)) {
            right_paddle.moveUp(dt);
        }

        if (rl.isKeyDown(.d)) {
            right_paddle.moveDown(dt);
        }

        left_paddle.render();
        right_paddle.render();
        ball.render();

        //----------------------------------------------------------------------------------
    }
}
