const std = @import("std");

pub fn main() !void {
    const cwd = std.fs.cwd();
    var buf: [30000]u8 = undefined;
    const content = try cwd.readFile("days/04.txt", &buf);
    try p1(content);
    try p2(content);
}

fn p1(buf: []const u8) !void {
    var c: u64 = undefined;
    std.debug.print("p1: {d}\n", .{c});
}

fn p2(buf: []const u8) !void {
    var c: u64 = undefined;
    std.debug.print("p2: {d}\n", .{c});
}
