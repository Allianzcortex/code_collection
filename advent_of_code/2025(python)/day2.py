def solve_day2_part1_wrong_version(input_):
    # only work for test case
    res = 0
    for ranges in input_.split(","):
        left, right = ranges.split("-")

        invalid_left = invalid_right = ""
        if len(left) % 2 == 0:
            left_left_half = left[:len(left)//2]
            left_right_half = left[len(left)//2:]
            if int(left_left_half) >= int(left_right_half):
                invalid_left = left_left_half+left_left_half

        if len(right) % 2 == 0:
            right_left_half = right[:len(right)//2]
            right_right_half = right[len(right)//2:]
            if int(right_left_half) <= int(right_right_half) and int(right_left_half) >= int(left_right_half):
                invalid_right = right_left_half + right_left_half
            if int(left_left_half) > int(right_right_half):
                invalid_left = ""

        res += int(invalid_left) if invalid_left != "" else 0
        res += int(invalid_right) if invalid_right != "" and invalid_left != invalid_right else 0

    return res


def solve_day2_part1_right_version(input_):
    res = 0
    for ranges in input_.split(","):
        left, right = ranges.split("-")
        # both are odds
        if len(left) % 2 == 1 and len(right) % 2 == 1:
            continue

        left_val, right_val = int(left), int(right)
        # if left is even
        if len(left) % 2 == 0:
            left_left_half = int(left[:len(left)//2])
            while True:
                next_val = f"{left_left_half}{left_left_half}"
                if left_val <= int(next_val) <= right_val:
                    res += int(next_val)
                elif int(next_val) > right_val:
                    break

                left_left_half += 1

        # if right is even
        elif len(right) % 2 == 0:
            right_left_half = int(right[:len(right)//2])
            while True:
                next_val = f"{right_left_half}{right_left_half}"
                if left_val <= int(next_val) <= right_val:
                    res += int(next_val)
                elif int(next_val) < left_val:
                    break

                right_left_half -= 1

    return res
