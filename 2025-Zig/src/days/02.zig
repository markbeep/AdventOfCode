const std = @import("std");

pub fn main() !void {
    const cwd = std.fs.cwd();
    var buf: [1024]u8 = undefined;
    const content = try cwd.readFile("days/02.txt", &buf);
    try p(content);
}

var c1: u64 = 0;
var c2: u64 = 0;

fn check(from: u64, to: u64) !void {
    for (from..to + 1) |x| {
        if (!try isValid(x)) _ = @atomicRmw(u64, &c1, .Add, x, .acq_rel);
        if (!try isValid2(x)) _ = @atomicRmw(u64, &c2, .Add, x, .acq_rel);
    }
}

fn p(buf: []u8) !void {
    var iter = std.mem.splitScalar(u8, buf, ',');
    var threads: [100]std.Thread = undefined;

    var i: u64 = 0;
    while (iter.next()) |elem| {
        var nums_iter = std.mem.splitScalar(u8, elem, '-');
        const from = try std.fmt.parseInt(u64, nums_iter.next().?, 10);
        const to = try std.fmt.parseInt(u64, nums_iter.next().?, 10);
        threads[i] = try std.Thread.spawn(.{}, check, .{ from, to });
        i += 1;
    }
    for (threads[0..i]) |thread| {
        thread.join();
    }
    std.debug.print("p1: {d}\np2: {d}\n", .{ c1, c2 });
}

fn isValid(x: u64) !bool {
    var buf: [20]u8 = undefined;
    const str = try std.fmt.bufPrint(&buf, "{d}", .{x});
    if (str.len % 2 == 1) return true;
    for (1..str.len) |i| {
        if (std.mem.eql(u8, str[0..i], str[i..])) return false;
    }
    return true;
}

fn isValid2(x: u64) !bool {
    var buf: [20]u8 = undefined;
    const str = try std.fmt.bufPrint(&buf, "{d}", .{x});
    for (1..str.len) |len| {
        if (str.len % len != 0) continue;
        var off = len;
        var all_invalid = true;
        while (off < str.len) : (off += len) {
            if (!std.mem.eql(u8, str[0..len], str[off .. off + len])) {
                all_invalid = false;
                break;
            }
        }
        if (all_invalid) return false;
    }
    return true;
}
