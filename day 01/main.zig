const std = @import("std");
const eql = std.mem.eql;
const ArrayList = std.ArrayList;
const allocator = std.heap.page_allocator;

pub fn main() !void {
    var file = try std.fs.cwd().openFile("input.txt", .{ .mode = .read_only, .lock = .none });
    defer file.close();

    var list1 = ArrayList(i32).init(allocator);
    var list2 = ArrayList(i32).init(allocator);

    var buf_reader = std.io.bufferedReader(file.reader());
    var in_stream = buf_reader.reader();

    var buf: [1024]u8 = undefined;
    while (try in_stream.readUntilDelimiterOrEof(&buf, '\n')) |line| {
        var tokens = std.mem.tokenizeSequence(u8, line, "   ");
        var i: u8 = 0;
        while (tokens.next()) |token| : (i += 1) {
            const t = std.mem.trim(u8, token, "\r");
            const num = try std.fmt.parseInt(i32, t, 10);
            if (i == 0) {
                try list1.append(num);
            } else {
                try list2.append(num);
            }
        }
    }

    const slice1 = try list1.toOwnedSlice();
    std.mem.sort(i32, slice1, {}, comptime std.sort.asc(i32));

    const slice2 = try list2.toOwnedSlice();
    std.mem.sort(i32, slice2, {}, comptime std.sort.asc(i32));

    var total: i32 = 0;
    for (slice1) |value1| {
        var cnt: i32 = 0;
        for (slice2) |value2| {
            if (value1 == value2) {
                cnt += 1;
            }
        }
        total += value1 * cnt;
        //const diff = @abs(slice1[index] - slice2[index]);
        //total += diff;
    }
    std.debug.print("total: {d}\n", .{total});
}
