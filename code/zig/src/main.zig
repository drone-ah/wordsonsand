//! By convention, main.zig is where your main function lives in the case that
//! you are building an executable. If you are making a library, the convention
//! is to delete this file and start with root.zig instead.

test "all" {
    const toml_with_defaults = @import("toml_with_defaults.zig");
    _ = toml_with_defaults;
}
