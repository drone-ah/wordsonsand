const std = @import("std");
const rl = @import("raylib");

const Ball = @This();

pos: rl.Vector2,
r: f32 = 16,
vel: rl.Vector2 = .{ .x = 1, .y = 0 },

pub fn render(self: Ball) void {
    rl.drawCircleV(self.pos, self.r, .white);
}

pub fn update(self: *Ball) void {
    self.pos = rl.math.vector2Add(self.pos, self.vel);
}
