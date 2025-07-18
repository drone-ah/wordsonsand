const std = @import("std");
const rl = @import("raylib");
const dvui = @import("dvui");

const Ball = @import("Ball.zig");

const Paddle = @This();

pub const Which = enum {
    left,
    right,
};

pos: rl.Vector2,
which: Which,
colour: rl.Color = .white,
play_area: dvui.Rect,

pub fn init(x: f32, which: Which, screen_width: f32, screen_height: f32) Paddle {
    const play_area: dvui.Rect = switch (which) {
        .left => .{ .x = 0, .y = 0, .w = screen_width * 0.5, .h = screen_height },
        .right => .{ .x = screen_width * 0.5, .y = 0, .w = screen_width * 0.5, .h = screen_height },
    };

    return .{
        .pos = .{
            .x = x,
            .y = screen_height * 0.5 - size.y * 0.5,
        },
        .which = which,

        .play_area = play_area,
    };
}

pub const size = rl.Vector2{ .x = 25, .y = 100 };

pub fn render(self: Paddle) void {
    rl.drawRectangleV(self.pos, size, self.colour);
}

pub fn isColliding(self: *const Paddle, ball: *const Ball) bool {
    // which edge do we need to check
    const crossing_x = switch (self.which) {
        .right => ball.vel.x > 0 and
            ball.pos.x + ball.r >= self.pos.x,
        .left => ball.vel.x < 0 and
            ball.pos.x - ball.r <= self.pos.x + size.x,
    };

    if (!crossing_x) {
        return false;
    }

    const colliding = ball.pos.y + ball.r >= self.pos.y and ball.pos.y - ball.r <= self.pos.y + size.y;

    return colliding;
}

pub fn moveUp(self: *Paddle, dt: f32) void {
    self.move(-100, dt);
}

pub fn moveDown(self: *Paddle, dt: f32) void {
    self.move(100, dt);
}

pub fn move(self: *Paddle, y: f32, dt: f32) void {
    self.pos.y += y * dt;
}
