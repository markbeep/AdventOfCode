const std = @import("std");

pub fn main() !void {
    var gpa = std.heap.GeneralPurposeAllocator(.{}).init;
    defer _ = gpa.deinit();
    const alloc = gpa.allocator();

    const cwd = std.fs.cwd();
    const buf = try cwd.readFileAlloc("days/02.txt", alloc, .unlimited);
    defer alloc.free(buf);

    // try p1(buf);
    try p2(buf);
}

fn comp(x: []u8, y: []u8) bool {
    if (x.len != y.len) return false;
    for (x, y) |xc, yc| {
        if (xc != yc) return false;
    }
    return true;
}

fn isValid(x: u64) !bool {
    var buf: [20]u8 = undefined;
    const str = try std.fmt.bufPrint(&buf, "{d}", .{x});
    if (str.len % 2 == 1) return true;
    for (1..str.len) |i| {
        if (comp(str[0..i], str[i..])) return false;
    }
    return true;
}

fn p1(buf: []u8) !void {
    var iter = std.mem.splitScalar(u8, buf, ',');
    var c: u64 = 0;
    while (iter.next()) |elem| {
        var nums_iter = std.mem.splitScalar(u8, elem, '-');
        const first = nums_iter.next() orelse "0";
        const second = nums_iter.next() orelse "0";
        const from = try std.fmt.parseInt(u64, first, 10);
        const to = try std.fmt.parseInt(u64, second, 10);
        for (from..to + 1) |x| {
            if (!try isValid(x)) c += x;
        }
    }
    std.debug.print("p1: {d}\n", .{c});
}

fn isValid2(x: u64) !bool {
    var buf: [20]u8 = undefined;
    const str = try std.fmt.bufPrint(&buf, "{d}", .{x});

    for (1..str.len) |len| {
        if (str.len % len != 0) continue;
        var off = len;
        var all_invalid = true;
        while (off < str.len) : (off += len) {
            if (!comp(str[0..len], str[off .. off + len])) {
                all_invalid = false;
                break;
            }
        }
        if (all_invalid) return false;
    }
    return true;
}

fn p2(buf: []u8) !void {
    var iter = std.mem.splitScalar(u8, buf, ',');
    var c: u64 = 0;
    while (iter.next()) |elem| {
        var nums_iter = std.mem.splitScalar(u8, elem, '-');
        const first = nums_iter.next() orelse "0";
        const second = nums_iter.next() orelse "0";
        const from = try std.fmt.parseInt(u64, first, 10);
        const to = try std.fmt.parseInt(u64, second, 10);
        for (from..to + 1) |x| {
            if (!try isValid2(x)) c += x;
        }
    }
    std.debug.print("p2: {d}\n", .{c});
}
