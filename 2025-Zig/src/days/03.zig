const std = @import("std");

pub fn main() !void {
    const cwd = std.fs.cwd();
    var buf: [30000]u8 = undefined;
    const content = try cwd.readFile("days/03.txt", &buf);
    std.debug.print("p1: {d}\n", .{try p(2, content)});
    std.debug.print("p2: {d}\n", .{try p(12, content)});
}

fn find_best(x: []const u8) struct { u64, u64 } {
    var ind: usize = 0;
    var val: u8 = x[0];
    for (x[1..], 1..) |c, i| {
        if (c > val) {
            val = c;
            ind = i;
        }
    }
    return .{ ind + 1, @as(u64, val - 48) };
}

fn p(comptime size: usize, buf: []const u8) !u64 {
    var c: u64 = 0;
    var iter = std.mem.splitScalar(u8, buf, '\n');
    while (iter.next()) |bank| {
        if (bank.len == 0) continue;
        var last_index: u64 = 0;
        inline for (0..size) |i| {
            const ind = find_best(bank[last_index .. bank.len - (size - 1) + i]);
            last_index = last_index + ind[0];
            const pow = comptime try std.math.powi(u64, 10, size - i - 1);
            c += ind[1] * pow;
        }
    }
    return c;
}
