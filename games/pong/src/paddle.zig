const std = @import("std");
const rl = @import("raylib");

const Ball = @import("ball.zig");

const Paddle = @This();

pub const Which = enum {
    left,
    right,
};

pos: rl.Vector2,
which: Which,
colour: rl.Color = .white,
score: u8 = 0,

pub fn init(x: f32, which: Which) Paddle {
    return .{
        .pos = .{
            .x = x,
            .y = 200,
        },
        .which = which,
    };
}

pub const size = rl.Vector2{ .x = 25, .y = 100 };

pub fn render(self: Paddle) void {
    rl.drawRectangleV(self.pos, size, self.colour);
}

pub fn isColliding(self: *Paddle, ball: *const Ball) bool {
    // which edge do we need to check
    const crossing_x: bool = switch (self.which) {
        .right => ball.pos.x + ball.r >= self.pos.x,
        .left => ball.pos.x - ball.r <= self.pos.x + size.x,
    };

    if (!crossing_x) {
        self.colour = .white;
        return false;
    }

    const colliding = ball.pos.y >= self.pos.y and ball.pos.y <= self.pos.y + size.y;

    self.colour = if (colliding) .red else .white;
    return colliding;
}

pub fn move(self: *Paddle, y: f32, dt: f32) void {
    self.pos.y += y * dt;
}
