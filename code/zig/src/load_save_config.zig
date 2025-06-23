const std = @import("std");

const toml = @import("toml");

const Self = @This();

const log = std.log.scoped(.config);

const config_path = "/tmp/wordsonsand";

// mocking out the library for this sample code
const MockKownFolders = struct {
    fn getPath(_: *const MockKownFolders, allocator: anytype, _: anytype) !?[]u8 {
        return try std.fmt.allocPrint(allocator, "{s}", .{config_path});
    }
};

const known_folders = MockKownFolders{};

const Game = struct {
    news_read: i64 = 0, // seconds since epoch
};

const ConfigError = error{
    UnableToDetermineConfigLocation,
};

allocator: std.mem.Allocator,

game_path: []const u8,
game: Game,

pub fn init(allocator: std.mem.Allocator) ConfigError!Self {
    const maybe_config = known_folders.getPath(allocator, .roaming_configuration) catch {
        return ConfigError.UnableToDetermineConfigLocation;
    };
    if (maybe_config) |config| {
        defer allocator.free(config);

        // game config path
        const game_path = std.fmt.allocPrint(allocator, "{s}/triangle/game.toml", .{config}) catch {
            std.debug.panic("oom", .{});
        };

        return .{
            .allocator = allocator,

            .game_path = game_path,
            .game = loadConfig(allocator, Game, game_path),
        };
    }

    return ConfigError.UnableToDetermineConfigLocation;
}

pub fn markNewsAsRead(self: *Self) void {
    self.game.news_read = std.time.timestamp();
    saveConfig(self.allocator, self.game, self.game_path);
}

pub fn deinit(self: *Self, allocator: std.mem.Allocator) void {
    allocator.free(self.game_path);
}

fn loadConfig(allocator: std.mem.Allocator, ConfigType: type, path: []const u8) ConfigType {
    var parser = toml.Parser(ConfigType).init(allocator);
    defer parser.deinit();

    var result = parser.parseFile(path) catch |err| {
        log.warn("unable to read config file: {s}, err: {any}", .{ path, err });
        return .{};
    };
    defer result.deinit();

    return result.value;
}

fn saveConfig(allocator: std.mem.Allocator, Config: anytype, full_path: []const u8) void {
    const path = std.fs.path.dirname(full_path) orelse ".";
    std.fs.cwd().makePath(path) catch |err| { // creates parent dirs if needed
        log.warn("unable to save: {any}", .{err});
        return;
    };

    var file = std.fs.cwd().createFile(full_path, .{
        .read = false,
        .truncate = true,
    }) catch |err| {
        log.warn("unable to save: {any}", .{err});
        return;
    };
    defer file.close();

    var writer = file.writer().any();
    toml.serialize(allocator, Config, &writer) catch |err| {
        log.warn("unable to write to config file: {any}", .{err});
    };
}

test "save and load config" {
    const allocator = std.testing.allocator;

    var config = try init(allocator);
    defer config.deinit(allocator);

    try std.testing.expectEqual(0, config.game.news_read);

    config.markNewsAsRead();

    var new_config = try init(allocator);
    defer new_config.deinit(allocator);

    try std.testing.expectEqual(config.game.news_read, new_config.game.news_read);

    // bazel will take care of this, but let's be polite
    try std.fs.deleteTreeAbsolute(config_path);
}
