// raylib-zig (c) Nikolas Wipper 2023

const rl = @import("raylib");
const Paddle = @import("paddle.zig");
const Ball = @import("ball.zig");

pub fn main() anyerror!void {
    // Initialization
    //--------------------------------------------------------------------------------------
    const screenWidth = 800;
    const screenHeight = 450;

    rl.initWindow(screenWidth, screenHeight, "raylib-zig [core] example - basic window");
    defer rl.closeWindow(); // Close window and OpenGL context

    rl.setTargetFPS(60); // Set our game to run at 60 frames-per-second
    //--------------------------------------------------------------------------------------
    var left_paddle = Paddle.init(Paddle.size.x * 0.5, .left);
    var right_paddle = Paddle.init(screenWidth - Paddle.size.x * 1.5, .right);
    var ball = Ball{ .pos = .{ .x = screenWidth * 0.5, .y = screenHeight * 0.5 } };

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

        ball.update(dt);
        _ = left_paddle.isColliding(ball);
        _ = right_paddle.isColliding(ball);

        left_paddle.render();
        right_paddle.render();
        ball.render();

        //----------------------------------------------------------------------------------
    }
}
