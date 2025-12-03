const std = @import("std");
const expect = std.testing.expect;

pub fn main() !void {
    const cwd = std.fs.cwd();
    var buf: [30000]u8 = undefined;
    const content = try cwd.readFile("days/03.txt", &buf);
    // std.debug.print("p1: {d}\n", .{try p1(content)});
    std.debug.print("p2: {d}\n", .{try p2(content)});
}

fn p1(buf: []const u8) !u64 {
    var c: u64 = 0;

    var iter = std.mem.splitScalar(u8, buf, '\n');
    while (iter.next()) |bank| {
        if (bank.len == 0) continue;
        var x: [2]u64 = undefined;
        var afts: [1]u64 = undefined;
        for (0..2) |i| x[i] = @as(u64, bank[bank.len - 2 + i]) - 48;
        afts[0] = x[1];
        for (2..bank.len) |i| {
            const batt = @as(u64, bank[bank.len - 1 - i]) - 48;
            trickle(batt, &x, &afts);
        }
        c += x[0] * 10 + x[1];
    }

    return c;
}

fn adder(x: []u64) !u64 {
    var n: u64 = 0;
    for (x, 1..) |xv, i| {
        n += (xv) * try std.math.powi(u64, 10, x.len - i);
    }
    return n;
}

fn find_max(bank: []const u8, x: []u64) !u64 {
    if (bank.len == x.len) return try adder(bank);
    if (bank.len < x.len) return 0;
    if (x.len == 0) return 0;
    var mx: u64 = 0;
    for (0..bank.len - x.len + 1) |i| {
        const batt = @as(u64, bank[i]) - 48;
        if (batt >= x[0]) {
            mx = @max(mx, batt * try std.math.powi(u64, 10, x.len - 1) + try find_max(bank[i + 1 ..], x[1..]));
        }
    }
    return mx;
}

fn trickle(batt: u64, x: []u64, afts: []u64) void {
    if (x.len < 2 or afts.len == 0) return;
    const pre = x[0];
    const aft = x[1];
    const aft_max = afts[0];
    if (batt >= pre) {
        // element is shifted to the left. shift following elements over if additional space frees up
        x[0] = batt;
        if (pre >= aft) {
            // shift aft up. old aft is now free.
            x[1] = pre;
            trickle(pre, x[1..], afts[1..]);
        }
        if (afts[0] >= x[1]) {
            // trickle(x[1], x[1..], afts[1..]);
            x[1] = afts[0];
        }
    } else if (batt >= aft_max) {
        // found a value better than current best value
        // save for when the parent is eventually shifted to the left
        afts[0] = batt;
        trickle(aft_max, x[1..], afts[1..]);
    }
    // TODO: trickle down batt to lower nums??
}

fn pr(x: []u64) void {
    for (x) |v| {
        std.debug.print("{d}", .{v});
    }
    std.debug.print("\n", .{});
}

const Index = struct {
    val: u64,
    index: u64,
};

fn find_best(x: []const u8) Index {
    var ind: u64 = 0;
    var val: u64 = @as(u64, x[0]) - 48;
    for (x[1..], 1..) |c, i| {
        const cv = @as(u64, c) - 48;
        if (cv > val) {
            val = cv;
            ind = i;
        }
    }
    return Index{ .val = val, .index = ind };
}

fn p2(buf: []const u8) !u64 {
    var c: u64 = 0;

    var iter = std.mem.splitScalar(u8, buf, '\n');
    while (iter.next()) |bank| {
        if (bank.len == 0) continue;
        var x: [12]u64 = undefined;

        var last_index: u64 = 0;
        for (0..12) |i| {
            const ind = find_best(bank[last_index .. bank.len - 11 + i]);
            last_index = last_index + ind.index + 1;
            x[i] = ind.val;
        }
        c += try adder(&x);
    }

    return c;
}

test "p1" {
    try expect(try p1("987654321111111") == 98);
    try expect(try p1("811111111111119") == 89);
    try expect(try p1("234234234234278") == 78);
    try expect(try p1("818181911112111") == 92);

    const cwd = std.fs.cwd();
    var buf: [30000]u8 = undefined;
    const content = try cwd.readFile("days/03.txt", &buf);
    var iter = std.mem.splitScalar(u8, content, '\n');
    while (iter.next()) |bank| {
        if (bank.len == 0) continue;
        var x: [2]u64 = undefined;
        var afts: [1]u64 = undefined;
        for (0..2) |i| x[i] = @as(u64, bank[bank.len - 2 + i]) - 48;
        afts[0] = x[1];
        for (2..bank.len) |i| {
            const batt = @as(u64, bank[bank.len - 1 - i]) - 48;
            if (batt >= x[0]) {
                if (x[0] > x[1]) x[1] = x[0];
                if (afts[0] > x[1]) x[1] = afts[0];
                x[0] = batt;
            } else if (batt > x[1]) afts[0] = x[1];
        }
        try expect(try p1(bank) == x[0] * 10 + x[1]);
    }
}

test "p2" {
    try expect(try p2("987654321111111") == 987654321111);
    try expect(try p2("811111111111119") == 811111111119);
    try expect(try p2("234234234234278") == 434234234278);
    try expect(try p2("818181911112111") == 888911112111);
}
