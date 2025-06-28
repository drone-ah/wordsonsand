const std = @import("std");
const rl = @import("raylib");

const Paddle = @This();

pos: rl.Vector2,

pub fn init(x: f32) Paddle {
    return .{ .pos = .{
        .x = x,
        .y = 200,
    } };
}

pub const size = rl.Vector2{ .x = 25, .y = 100 };

pub fn render(self: Paddle) void {
    rl.drawRectangleV(self.pos, size, .white);
}
