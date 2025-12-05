
def solve_day5_part1(input_):

    ranges = []
    cnt = 0

    for line in input_:
        if line.strip() == "":
            merged_ranges = _merge_intervals(ranges)
            continue
        if "-" in line:
            start, end = map(int, line.split("-"))
            # Process range
            ranges.append([start, end])
        else:
            number = int(line)
            # Process ingredients
            for range in merged_ranges:
                if range[0] <= number <= range[1]:
                    cnt += 1

    return cnt


def _merge_intervals(intervals):
    if not intervals:
        return []

    intervals.sort(key=lambda x: x[0])
    res = [intervals[0]]
    prev_end = intervals[0][1]
    for current in intervals[1:]:
        # no overlap
        if current[0] > prev_end:
            res.append(current)
            prev_end = current[1]
        else:
        # with overlap, choose maximum end as new end
            res[-1][1] = max(prev_end, current[1])
            prev_end = res[-1][1]
    return res


def solve_day5_part2(input_):
    ranges = []
    cnt = 0

    for line in input_:

        cnt = 0
        if "-" in line:
            start, end = map(int, line.split("-"))
            # Process range
            ranges.append([start, end])

        if line.strip() == "":
            merged_ranges = _merge_intervals(ranges)
            for range in merged_ranges:
                cnt += (range[1]-range[0]+1)

            return cnt
