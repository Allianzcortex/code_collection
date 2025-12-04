
dirs = [
    [-1, -1],
    [-1, 0],
    [-1, 1],
    [0, -1],
    [0, 1],
    [1, -1],
    [1, 0],
    [1, 1]
]


def solve_day4_part1(grid):
    res = 0
    row, column = len(grid), len(grid[0])

    for i in range(row):
        for j in range(column):
            if grid[i][j] == '.':
                continue
            roll_cnt = 0
            for dir in dirs:
                x, y = i+dir[0], j+dir[1]
                if x < 0 or x > row-1 or y < 0 or y > column-1:
                    continue
                if grid[x][y] == '@':
                    roll_cnt += 1

            if roll_cnt < 4:
                res += 1

    return res


def solve_day4_part2(grid):
    res = 0
    row, column = len(grid), len(grid[0])

    while True:
        is_roll_moved = False
        for i in range(row):
            for j in range(column):
                if grid[i][j] == '.':
                    continue
                roll_cnt = 0
                for dir in dirs:
                    x, y = i+dir[0], j+dir[1]
                    if x < 0 or x > row-1 or y < 0 or y > column-1:
                        continue
                    if grid[x][y] == '@':
                        roll_cnt += 1

                if roll_cnt < 4:
                    is_roll_moved = True
                    grid[i] = grid[i][:j]+'.'+grid[i][j+1:]
                    res += 1

        if not is_roll_moved:
            break
    return res
