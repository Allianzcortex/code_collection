
from collections import defaultdict

def solve_day5_part1(rules:list[str], pages:list[str])->int:
    before_rules = defaultdict(list)
    for rule in rules:
        first, second = rule.split("|")
        # each item stores what item must be before current one
        # before the rule : 75 : [97]
        # after the rule :  97 : [75]
        before_rules[int(second)].append(int(first))
    
    res = 0
    for page in pages:
        nums = [int(num) for num in page.split(",")]
        is_ok = True
        for i in range(1, len(nums)):
            for j in range(i):
                # check whether it fits with the rule
                first_num, second_num = nums[j], nums[i]
                if second_num in before_rules[first_num]:
                    is_ok = False
                    break
            if not is_ok:
                break
        
        if is_ok:
            res += nums[len(nums)//2-1 if len(nums)%2==0 else len(nums)//2]

    return res
