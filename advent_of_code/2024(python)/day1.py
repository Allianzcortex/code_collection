
from collections import Counter

def _process_input(inp:list[str]) -> tuple[list[int], list[int]]:
    left,right = [],[]
    for row in inp:
        left_item, right_item = row.split(",")
        left.append(int(left_item))
        right.append(int(right_item))
    
    return left, right

def solve_day1_part1(inp:list[str]) -> int:
    
    left_arr, right_arr = _process_input(inp)
    res = 0
    for left_item, right_item in zip(sorted(left_arr), sorted(right_arr)):
        res += abs(left_item-right_item)
    
    return res

def solve_day1_part2(inp:list[str]) -> int:

    left_arr, right_arr = _process_input(inp)
    right_counter = Counter(right_arr)

    res = 0
    for left_item in left_arr:
        res += left_item * right_counter.get(left_item, 0)
    
    return res
  