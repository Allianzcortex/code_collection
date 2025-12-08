
def solve_day7_part1(input_):
    input_ = [list(line) for line in input_]
    split_cnt = 0
    for i, line in enumerate(input_):
        if i == 0:
            line[line.index('S')] = '|'
            continue

        for j, ch in enumerate(line):
            if input_[i-1][j] == '|':
                if ch == '^':
                    split_cnt += 1
                    if j > 0:
                        line[j-1] = '|'
                    if j < len(line)-1:
                        line[j+1] = '|'
                else:
                    line[j] = '|'

    return split_cnt
