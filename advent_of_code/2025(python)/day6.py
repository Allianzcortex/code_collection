

def solve_day6_part1(input_):
    res = 0

    section = len(input_[0].split())
    questions = [[] for _ in range(section)]

    for line in input_:
        for i, q in enumerate(line.split()):
            questions[i].append(q)

    for q in questions:
        match q[-1]:
            case '*':
                temp_res = 1
                for num in q[:-1]:
                    temp_res *= int(num)
            case '+':
                temp_res = 0
                for num in q[:-1]:
                    temp_res += int(num)

        res += temp_res

    return res


def solve_day6_part2(input_):
    res = 0

    section = len(input_[0].split())
    questions = [[] for _ in range(section)]

    for line in input_:
        for i, q in enumerate(line.split()):
            questions[i].append(q)

    aligned_questions = [[] for _ in range(section)]
    start_index = 0
    for i in range(section):
        max_length = len(sorted(questions[i], key=lambda x: len(x))[-1])
        aligned_questions[i] = [
            line[start_index:start_index+max_length] for line in input_]
        start_index += (max_length+1)

    for q in aligned_questions:
        match q[-1].strip():
            case '*':
                start_index = 0
                temp_res = 1
                while True:
                    is_last_digit_reached = False
                    curr_num = ""
                    for num in q[:-1]:
                        if start_index <= len(num)-1:
                            is_last_digit_reached = True
                            curr_num += num[start_index]

                    if is_last_digit_reached:
                        temp_res *= int(curr_num)
                    else:
                        break

                    start_index += 1
            case '+':
                start_index = 0
                temp_res = 0
                while True:
                    is_last_digit_reached = False
                    curr_num = ""
                    for num in q[:-1]:
                        if start_index <= len(num)-1:
                            is_last_digit_reached = True
                            curr_num += num[start_index]

                    if is_last_digit_reached:
                        temp_res += int(curr_num)
                    else:
                        break

                    start_index += 1
        res += temp_res

    return res
