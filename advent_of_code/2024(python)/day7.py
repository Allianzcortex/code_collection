

def _is_equal_part1(target:int, nums:list[int], prefix:int)->bool:
    if (not nums):
        return prefix == target
    
    is_equal = False
    
    # we can do `*` or `+` 2 operators
    prefix_mul = prefix*nums[0]
    is_equal = is_equal or _is_equal_part1(target,nums[1:],prefix_mul)

    prefix_add = prefix+nums[0]
    is_equal = is_equal or _is_equal_part1(target,nums[1:],prefix_add)

    return is_equal


def solve_day7_part1(matrix:list[str])->int:
    res = 0
    for line in matrix:
        expected_answer, given_input = line.split(":")
        target = int(expected_answer)
        nums = [int(num) for num in given_input.strip().split(" ")]
        if _is_equal_part1(target, nums[1:],nums[0]):
            res += target
    
    return res



def _is_equal_part2(target:int, nums:list[int], prefix:int)->bool:
    if (not nums):
        return prefix == target
    
    is_equal = False
    
    # we can do `*` or `+` or `||` 3 operators
    prefix_mul = prefix*nums[0]
    is_equal = is_equal or _is_equal_part2(target,nums[1:],prefix_mul)

    prefix_add = prefix+nums[0]
    is_equal = is_equal or _is_equal_part2(target,nums[1:],prefix_add)

    prefix_concat = int(str(prefix) + str(nums[0]))
    is_equal = is_equal or _is_equal_part2(target,nums[1:],prefix_concat)

    return is_equal


def solve_day7_part2(matrix:list[str])->int:
    res = 0
    for line in matrix:
        expected_answer, given_input = line.split(":")
        target = int(expected_answer)
        nums = [int(num) for num in given_input.strip().split(" ")]
        if _is_equal_part2(target, nums[1:],nums[0]):
            res += target

    return res