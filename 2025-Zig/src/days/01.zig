const std = @import("std");

pub fn main() !void {
    var gpa = std.heap.GeneralPurposeAllocator(.{}).init;
    defer _ = gpa.deinit();
    const alloc = gpa.allocator();

    const cwd = std.fs.cwd();
    const buf = try cwd.readFileAlloc("days/01.txt", alloc, .unlimited);
    defer alloc.free(buf);

    try p1(buf);
    try p2(buf);
}

fn p1(buf: []u8) !void {
    var pw: u32 = 0;
    var c: i32 = 50;
    var split_iterator = std.mem.splitScalar(u8, buf, '\n');
    while (split_iterator.next()) |line| {
        if (line.len == 0) continue;
        const move = try std.fmt.parseInt(i32, line[1..], 10);
        c = @mod(if (line[0] == 'L') c - move else c + move, 100);
        pw += @intFromBool(c == 0);
    }
    std.debug.print("p1: {d}\n", .{pw});
}

fn p2(buf: []u8) !void {
    var pw: i32 = 0;
    var c: i32 = 50;
    var split_iterator = std.mem.splitScalar(u8, buf, '\n');
    while (split_iterator.next()) |line| {
        if (line.len == 0) continue;
        const move = try std.fmt.parseInt(i32, line[1..], 10);
        if (line[0] == 'L') {
            if (move >= c) {
                pw += @intFromBool(c != 0);
                pw += @divFloor(move - c, 100);
            }
            c = @mod(c - move, 100);
        } else {
            if (move >= 100 - c) {
                pw += 1 + @divFloor(move - (100 - c), 100);
            }
            c = @mod(c + move, 100);
        }
    }
    std.debug.print("p2: {d}\n", .{pw});
}
