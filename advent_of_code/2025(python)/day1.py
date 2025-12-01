
def solve_day1_part1(actions):
    start_degree = 50
    res = 0

    for ac in actions:
        step = (int(ac[1:])) % 100
        match ac[0]:
            case 'L':
                start_degree -= step
                if start_degree < 0:
                    start_degree = 100 + start_degree
            case 'R':
                start_degree += step
                if start_degree >= 100:
                    start_degree = start_degree-100

        res += 1 if start_degree == 0 else 0

    return res


def solve_day1_part2(actions):
    start_degree = 50
    res = 0

    for ac in actions:
        step = (int(ac[1:]))
        res += step//100

        step = step % 100
        start_degree = start_degree % 100
        if step == 0:
            continue

        initial_start = start_degree
        match ac[0]:
            case 'L':
                start_degree -= step
                if start_degree <= 0:
                    start_degree = 100 + start_degree
                    if initial_start != 0 and start_degree != 0:
                        res += 1
            case 'R':
                start_degree += step
                if start_degree >= 100:
                    start_degree = start_degree-100
                    if initial_start != 0 and start_degree != 0:
                        res += 1
        res += 1 if start_degree == 0 else 0

    return res
