const std = @import("std");

pub fn main() !void {
    const cwd = std.fs.cwd();
    var buf: [30000]u8 = undefined;
    const content = try cwd.readFile("days/xx.txt", &buf);
    try p1(content);
    // try p2(content);
}

fn p1(buf: []const u8) !void {
    var c: u64 = 0;
    std.debug.print("p1: {d}\n", .{c});
}

fn p2(buf: []const u8) !void {
    var c: u64 = 0;
    std.debug.print("p2: {d}\n", .{c});
}
