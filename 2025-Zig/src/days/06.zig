const std = @import("std");

pub fn main() !void {
    const cwd = std.fs.cwd();
    var buf: [30000]u8 = undefined;
    const content = try cwd.readFile("days/06.txt", &buf);
    try p1(1000, 4, content);
    try p2(content);
}

fn p1(comptime size: usize, comptime height: usize, buf: []const u8) !void {
    var c: u64 = 0;
    var nums: [height * size]u64 = undefined;
    var adds: [size]bool = undefined;

    // nums
    var iter = std.mem.splitScalar(u8, buf, '\n');
    var li: usize = 0;
    var ni: usize = 0;
    while (iter.next()) |line| {
        if (line.len == 0) continue;
        var iter_nums = std.mem.tokenizeScalar(u8, line, ' ');
        while (iter_nums.next()) |num| {
            nums[ni] = try std.fmt.parseInt(u64, num, 10);
            ni += 1;
        }

        li += 1;
        if (li == height) break;
    }
    // ops
    var iter_nums = std.mem.tokenizeScalar(u8, iter.next().?, ' ');
    ni = 0;
    while (iter_nums.next()) |num| {
        adds[ni] = num[0] == '+';
        ni += 1;
    }

    // add em up
    for (0..size) |i| {
        const _1 = nums[0 * size + i];
        const _2 = nums[1 * size + i];
        const _3 = nums[2 * size + i];
        const _4 = nums[3 * size + i];
        const add = adds[i];
        if (add) {
            c += _1 + _2 + _3 + _4;
        } else {
            c += _1 * _2 * _3 * _4;
        }
    }

    std.debug.print("p1: {d}\n", .{c});
}

fn p2(buf: []const u8) !void {
    var c: u64 = 0;
    var lines: [5][]const u8 = undefined;

    var iter = std.mem.splitScalar(u8, buf, '\n');
    var li: usize = 0;
    while (iter.next()) |line| {
        if (line.len == 0) continue;
        lines[li] = line;
        li += 1;
    }

    var i: usize = lines[0].len;
    var nums: [4]u64 = undefined;
    var nc: usize = 0;
    var skip_one = false;
    while (i > 0) {
        i -= 1;
        if (skip_one) {
            nc = 0;
            skip_one = false;
            continue;
        }

        var n: u64 = 0;
        if (lines[0][i] != ' ') n = lines[0][i] - 48;
        if (lines[1][i] != ' ') n = n * 10 + lines[1][i] - 48;
        if (lines[2][i] != ' ') n = n * 10 + lines[2][i] - 48;
        if (lines[3][i] != ' ') n = n * 10 + lines[3][i] - 48;
        nums[nc] = n;
        nc += 1;

        if (lines[4][i] == '+') {
            skip_one = true;
            for (nums[0..nc]) |val| {
                c += val;
            }
        } else if (lines[4][i] == '*') {
            skip_one = true;
            var t: u64 = 1;
            for (nums[0..nc]) |val| {
                t *= val;
            }
            c += t;
        }
    }

    std.debug.print("p2: {d}\n", .{c});
}
