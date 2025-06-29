const std = @import("std");
const rl = @import("raylib");

const Paddle = @import("paddle.zig");

const Ball = @This();

pos: rl.Vector2,
r: f32 = 16,
vel: rl.Vector2 = .{ .x = 250, .y = 300 },

pub fn render(self: Ball) void {
    rl.drawCircleV(self.pos, self.r, .white);
}

pub fn update(self: *Ball, dt: f32) void {
    const vel_this_frame = rl.math.vector2Scale(self.vel, dt);
    self.pos = rl.math.vector2Add(self.pos, vel_this_frame);
}

pub fn checkPaddleCollision(self: *Ball, paddle: *Paddle) void {
    if (paddle.isColliding(self)) {
        self.vel = rl.math.vector2Scale(self.vel, -1);
    }
}

pub fn checkEdgeCollisions(self: *Ball, screen_height: f32) void {
    if (self.pos.y < self.r or self.pos.y > screen_height - self.r) {
        self.vel.y *= -1;
    }
}
