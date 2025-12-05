const std = @import("std");

pub fn main() !void {
    const cwd = std.fs.cwd();
    var buf: [30000]u8 = undefined;
    const content = try cwd.readFile("days/05.txt", &buf);
    try p(content);
}

const Range = struct { from: u64, to: u64, inv: bool = false };

fn p(buf: []const u8) !void {
    var ranges: [200]Range = undefined;
    var rc: usize = 0;

    var iter = std.mem.splitScalar(u8, buf, '\n');
    while (iter.next()) |line| {
        if (line.len == 0) break;
        var nums = std.mem.splitScalar(u8, line, '-');
        ranges[rc] = .{
            .from = try std.fmt.parseInt(u64, nums.next().?, 10),
            .to = try std.fmt.parseInt(u64, nums.next().?, 10),
        };
        rc += 1;
    }

    // part 1
    var c1: u64 = 0;
    while (iter.next()) |line| {
        if (line.len == 0) continue;
        const val = try std.fmt.parseInt(u64, line, 10);
        for (ranges[0..rc]) |range| {
            if (range.from <= val and val <= range.to) {
                c1 += 1;
                break;
            }
        }
    }

    // part 2
    var changed: bool = true; // lol
    while (changed) {
        changed = false;
        for (0..rc, ranges[0..rc]) |a, *ra| {
            for (ranges[a + 1 .. rc]) |*rb| {
                // (): a, []: b
                if (ra.inv or rb.inv) continue;
                // (  []  )
                if (ra.from <= rb.from and ra.to >= rb.to) {
                    rb.inv = true;
                    changed = true;
                } // (   [ ) ]
                else if (ra.from <= rb.from and ra.to <= rb.to and ra.to >= rb.from) {
                    ra.to = rb.to;
                    rb.inv = true;
                    changed = true;
                } // [  ()  ]
                else if (ra.from >= rb.from and ra.to <= rb.to) {
                    ra.inv = true;
                    changed = true;
                } // [ ( ] )
                else if (ra.from >= rb.from and ra.to >= rb.to and ra.from <= rb.to) {
                    ra.from = rb.from;
                    rb.inv = true;
                    changed = true;
                }
            }
        }
    }

    var c2: u64 = 0;
    for (ranges[0..rc]) |r| {
        if (!r.inv) {
            c2 += r.to - r.from + 1;
        }
    }
    std.debug.print("p1: {d}\n", .{c1});
    std.debug.print("p2: {d}\n", .{c2});
}
