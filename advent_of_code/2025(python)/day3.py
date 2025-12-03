import heapq


def solve_day3_part1(input_):
    res = 0
    for line in input_:
        rel = {}
        max_ch, second_max_ch = '', ''
        for i, ch in enumerate(line):
            rel[ch] = i
            if max_ch == '' or (ch > max_ch and i < len(line)-1):
                second_max_ch = ''
                max_ch = ch
            elif second_max_ch == '' or ch > second_max_ch:
                second_max_ch = ch

        if rel[max_ch] < rel[second_max_ch]:
            res += int(f"{max_ch}{second_max_ch}")
        else:
            res += int(f"{second_max_ch}{max_ch}")

    return res


def solve_day3_part2(input_):
    res = 0
    for line in input_:
        # construct a priority-queue, sort based on (largest, left-most) order
        pq = []
        for i, ch in enumerate(line):
            pq.append((-1*int(ch), -1*(len(line)-i)))

        numbers = 12
        powers = ""
        prev_index = float('inf')
        heapq.heapify(pq)
        while numbers > 0:
            invalid_item = []
            while True:
                curr = heapq.heappop(pq)
                if -1*curr[1] >= numbers and -1*(curr[1]) < prev_index:
                    powers += str(-1*curr[0])
                    prev_index = -1*curr[1]
                    break
                else:
                    invalid_item.append(curr)

            # add original back
            for item in invalid_item:
                heapq.heappush(pq, item)
            numbers -= 1

        res += int(powers)

    return res
