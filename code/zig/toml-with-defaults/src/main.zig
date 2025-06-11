//! By convention, main.zig is where your main function lives in the case that
//! you are building an executable. If you are making a library, the convention
//! is to delete this file and start with root.zig instead.

const std = @import("std");

pub fn main() !void {
    std.debug.print("all in the tests", .{});
}

const Controls = struct {
    forward: []const u8 = "w",
    craft: []const u8 = "q",
    inventory: []const u8 = "e",
};

const User = struct {
    controls: Controls = .{},
};

test "load partial toml config" {
    const toml = @import("toml");
    const allocator = std.testing.allocator;
    var parser = toml.Parser(User).init(allocator);
    defer parser.deinit();

    const source =
        \\[controls]
        \\craft = "s"
    ;
    var result = try parser.parseString(source);
    defer result.deinit();

    const config = result.value;
    const default = User{};
    try std.testing.expectEqualStrings(default.controls.forward, config.controls.forward);
    try std.testing.expectEqualStrings("s", config.controls.craft);
}
