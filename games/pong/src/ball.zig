const std = @import("std");
const rl = @import("raylib");

const Ball = @This();

pos: rl.Vector2,
r: f32 = 16,
vel: rl.Vector2 = .{ .x = 0, .y = 0 },

pub fn render(self: Ball) void {
    rl.drawCircleV(self.pos, self.r, .white);
}
