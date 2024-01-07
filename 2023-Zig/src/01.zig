const std = @import("std");

const f = @embedFile("01.txt");
const s = std.mem.trim(u8, f, " \n");

pub fn main() !void {
    std.debug.print("p1 = {!}\n", .{p1(s)});
    std.debug.print("p2 = {!}\n", .{p2()});
}

fn p1(str: []const u8) !i64 {
    var it = std.mem.split(u8, str, "\n");
    // Part 1
    var total: i64 = 0;
    while (it.next()) |line| {
        var first: ?u32 = null;
        var last: ?u32 = null;
        for (line) |ch| {
            if ('0' <= ch and ch <= '9') {
                if (first == null) {
                    first = ch - '0';
                }
                last = ch - '0';
            }
        }
        total += 10 * first.? + last.?;
    }

    return total;
}

test "p1" {
    const v = try p1(s);
    try std.testing.expect(53921 == v);
}

fn p2() !i64 {
    const R = struct { n: []const u8, r: []const u8 };
    const replacement = [_]R{
        R{ .n = "one", .r = "one1one" },
        R{ .n = "two", .r = "two2two" },
        R{ .n = "three", .r = "three3three" },
        R{ .n = "four", .r = "four4four" },
        R{ .n = "five", .r = "five5five" },
        R{ .n = "six", .r = "six6six" },
        R{ .n = "seven", .r = "seven7seven" },
        R{ .n = "eight", .r = "eight8eight" },
        R{ .n = "nine", .r = "nine9nine" },
    };
    var arena = std.heap.ArenaAllocator.init(std.heap.page_allocator);
    defer arena.deinit();
    var output: []u8 = try arena.allocator().alloc(u8, s.len);
    @memcpy(output, s);
    for (replacement) |r| {
        const size = std.mem.replacementSize(u8, output, r.n, r.r);
        const old = output;
        output = try arena.allocator().alloc(u8, size);
        _ = std.mem.replace(u8, old, r.n, r.r, output);
    }

    return p1(output);
}

test "p2" {
    const v = try p2();
    try std.testing.expect(54676 == v);
}
