const std = @import("std");

pub fn main() !void {
    const cwd = std.fs.cwd();
    var buf: [30000]u8 = undefined;
    const content = try cwd.readFile("days/04.txt", &buf);
    try p(140, false, content);
    try p(140, true, content);
}

inline fn get(buf: []const u8, y: usize, x: usize) u8 {
    return buf[y * 141 + x]; // buf includes newlines
}

fn p(comptime size: usize, comptime p2: bool, buf: []u8) !void {
    var c: u64 = 0;
    var additional: u64 = 1;
    while (additional > 0) {
        additional = 0;
        for (0..size) |y| {
            for (0..size) |x| {
                if (get(buf, y, x) == '.') continue;
                var cur: u8 = 0;
                if (y > 0 and x > 0 and get(buf, y - 1, x - 1) == '@') cur += 1;
                if (y > 0 and get(buf, y - 1, x) == '@') cur += 1;
                if (y > 0 and x < size - 1 and get(buf, y - 1, x + 1) == '@') cur += 1;
                if (x > 0 and get(buf, y, x - 1) == '@') cur += 1;
                if (x < size - 1 and get(buf, y, x + 1) == '@') cur += 1;
                if (y < size - 1 and x > 0 and get(buf, y + 1, x - 1) == '@') cur += 1;
                if (y < size - 1 and get(buf, y + 1, x) == '@') cur += 1;
                if (y < size - 1 and x < size - 1 and get(buf, y + 1, x + 1) == '@') cur += 1;
                if (cur < 4) {
                    additional += 1;
                    if (p2) buf[y * 141 + x] = '.';
                }
            }
        }
        c += additional;
        if (!p2) additional = 0;
    }
    std.debug.print(if (p2) "p2: {d}\n" else "p1: {d}\n", .{c});
}
