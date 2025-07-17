const std = @import("std");
const rl = @import("raylib");
const dvui = @import("dvui");

const Ball = @import("Ball.zig");
const Paddle = @import("Paddle.zig");

const Game = @This();

left_paddle: Paddle,
right_paddle: Paddle,
ball: Ball,
screen_height: f32,
screen_width: f32,

pub fn init(screen_width: f32, screen_height: f32) Game {
    return .{
        .left_paddle = .init(Paddle.size.x * 0.5, .left, screen_width, screen_height),
        .right_paddle = .init(screen_width - Paddle.size.x * 1.5, .right, screen_width, screen_height),
        .ball = .init(.{ .x = screen_width * 0.5, .y = screen_height * 0.5 }),
        .screen_height = screen_height,
        .screen_width = screen_width,
    };
}

pub fn update(self: *Game, dt: f32) void {
    self.ball.checkEdgeCollisions(self.screen_height);
    self.ball.update(dt);
    self.ball.checkPaddleCollision(&self.left_paddle);
    self.ball.checkPaddleCollision(&self.right_paddle);
    if (self.ball.pos.x + self.ball.r > self.screen_width) {
        self.left_paddle.score += 1;
        std.debug.print("scores: l: {d}, r: {d}\n", .{ self.left_paddle.score, self.right_paddle.score });
        self.ball.reset();
    }

    if (self.ball.pos.x < self.ball.r) {
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

    showScore(self.left_paddle);
    showScore(self.right_paddle);
}

fn showScore(paddle: Paddle) void {
    const xpos = paddle.play_area.x + (paddle.play_area.w * 0.5);

    const id: usize = @intFromFloat(xpos);
    var right = dvui.box(@src(), .horizontal, .{ .rect = .{ .x = xpos, .y = 50, .w = 150, .h = 150 }, .id_extra = id });
    defer right.deinit();

    const font_size: f32 = 64;
    var label_options: dvui.Options = .{
        .color_text = .white,
        .font_style = .title,
    };
    label_options.font = label_options.fontGet().resize(font_size);
    dvui.label(@src(), "{d}", .{paddle.score}, label_options);
}
