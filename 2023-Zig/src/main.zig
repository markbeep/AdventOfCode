const std = @import("std");
const d1 = @import("01.zig");

pub fn main() !void {
    try d1.main();
    std.debug.print("Nothing here in the main file\n", .{});
}
