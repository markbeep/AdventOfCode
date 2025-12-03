const std = @import("std");

pub fn main() !void {
    const cwd = std.fs.cwd();
    var buf: [30000]u8 = undefined;
    const content = try cwd.readFile("days/03.txt", &buf);
    // try p1(content);
    try p2(content);
}

fn p1(buf: []u8) !void {
    var c: u64 = 0;

    var iter = std.mem.splitScalar(u8, buf, '\n');
    while (iter.next()) |bank| {
        if (bank.len == 0) continue;
        var pre = @as(u64, bank[bank.len - 2]) - 48;
        var aft = @as(u64, bank[bank.len - 1]) - 48;
        var aft_max = aft;
        for (2..bank.len) |i| {
            const batt = @as(u64, bank[bank.len - 1 - i]) - 48;
            if (batt >= pre) {
                if (pre > aft) aft = pre;
                if (aft_max > aft) aft = aft_max;
                pre = batt;
            } else if (batt > aft) aft_max = aft;
        }
        c += pre * 10 + aft;
    }

    std.debug.print("p1: {d}\n", .{c});
}

fn adder(x: []const u8) !u64 {
    var n: u64 = 0;
    for (x, 1..) |xv, i| {
        n += (xv - 48) * try std.math.powi(u64, 10, x.len - i);
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

fn p2(buf: []u8) !void {
    var c: u64 = 0;

    var iter = std.mem.splitScalar(u8, buf, '\n');
    while (iter.next()) |bank| {
        if (bank.len == 0) continue;
        var x: [12]u64 = undefined;
        for (0..12) |i| x[i] = @as(u64, bank[i]) - 48;
        const mx: u64 = try find_max(bank, &x);
        c += mx;
        std.debug.print("bank: {d}\n", .{mx});
    }

    std.debug.print("p2: {d}\n", .{c});
}
