
def _is_single_record_safe(record:list[int]) -> bool:
    print(record)
    if not 1 <= abs(record[0] - record[1]) <= 3:
        return False
    is_increasing = record[0] < record[1]

    for i in range(1, len(record)-1):
        diff = record[i] - record[i+1]
        if (is_increasing and diff>0) or (not is_increasing and diff<0):
            return False
        if not 1 <= abs(diff) <= 3:
            return False

    return True

def _is_single_comprehensive(record:list[int]) -> bool:
    if not 1 <= abs(record[0] - record[1]) <= 3:
        # so we can remove either first or second
        print(record)
        remove_1st_item_record = record[1:]
        remove_2nd_item_record = record[:1] + record[2:]
        print("==fist==")
        print(remove_1st_item_record)
        print(remove_2nd_item_record)
        return _is_single_record_safe(remove_1st_item_record) or _is_single_record_safe(remove_2nd_item_record)
    
    is_increasing = record[0] < record[1]

    for i in range(1, len(record)-1):
        diff = record[i] - record[i+1]
        if (is_increasing and diff>0) or (not is_increasing and diff<0) or (not 1 <= abs(diff) <= 3):
            # either remove ith or (i+1)th item
            remove_ith_item_record = record[:i] + record[i+1:]
            remove_i1th_item_record = record[:i+1] + record[i+2:]
            print(record)
            print(i)
            print("==second==")
            print(remove_ith_item_record)
            print(remove_i1th_item_record)
            return _is_single_record_safe(remove_ith_item_record) or _is_single_record_safe(remove_i1th_item_record)
        
    return True


def solve_day2_part1(inp:list[str]) -> int:

    res = 0
    for report in inp:
        record = [int(item) for item in report.split()]
        # check first gap

        res += int(_is_single_record_safe(record))

    return res

def solve_day2_part2(inp:list[str]) -> int:

    res = 0
    for report in inp:
        record = [int(item) for item in report.split()]
        res += int(_is_single_comprehensive(record))

    return res