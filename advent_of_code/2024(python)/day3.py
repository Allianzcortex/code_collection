import re

def solve_day3_part1(instruction:str) -> int:

    res = 0
    match = re.findall(r'mul\((\d+),(\d+)\)', instruction)
    for (item1, item2) in match:
        res += int(item1) * int(item2)
    
    return res

def _is_do_command(instruction:str, i:int) -> bool:
    return (i+3)<len(instruction) and instruction[i:i+4] == "do()"

def _is_dont_command(instruction:str, i:int) -> bool:
    return (i+6)<len(instruction) and instruction[i:i+7] == "don't()"

def solve_day3_part2(instruction:str) -> int:

    res = 0
    index = 0

    new_instruction = "do()" + instruction + "don't()"
    while index < len(new_instruction):
        # !^don't()
        if(index==len(new_instruction)-7) and new_instruction[index:] == "don't()":
            break

        left_index = right_index = 0
        while not _is_do_command(new_instruction, index):
            left_index = index
            index += 1
        
        while not _is_dont_command(new_instruction, index):
            right_index = index
            index += 1

        res += solve_day3_part1(new_instruction[left_index:right_index+1])

    return res
