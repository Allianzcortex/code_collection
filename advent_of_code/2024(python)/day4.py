
prefix_dir = {
    "up": [-1, 0],
    "down": [1, 0],
    "left": [0, -1],
    "right": [0, 1],
    "up-left": [-1, -1],
    "up-right": [-1, 1],
    "down-left": [1, -1],
    "down-right": [1, 1],
}

def _solve_matrix_part1(matrix:list[str], target:str, x:int, y:int, dir:str) -> bool:

    if x<0 or x>len(matrix)-1 or y<0 or y>len(matrix[0])-1:
        return False
    if matrix[x][y] != target[0]:
        return False
    if len(target) == 1:
        return True
    return _solve_matrix_part1(matrix,target[1:],x+prefix_dir[dir][0], y+prefix_dir[dir][1], dir)

def solve_day4_part1(matrix:list[str]) -> int:

    target = "XMAS"
    res = 0
    for i in range(len(matrix)):
        for j in range(len(matrix[0])):
            for dir in prefix_dir.keys():
                res += int(_solve_matrix_part1(matrix, target, i, j, dir))
    
    return res

def _solve_matrix_part2(matrix:list[str], x:int, y:int) -> bool:

    if x+2>len(matrix)-1 or y+2>len(matrix[0])-1:
        return False
    x1_char = matrix[x][y] + matrix[x+1][y+1] + matrix[x+2][y+2]
    x2_char = matrix[x][y+2] + matrix[x+1][y+1] + matrix[x+2][y]

    targets = ("MAS", "SAM")
    if x1_char in targets and x2_char in targets:
        return True
    
    return False

def solve_day4_part2(matrix:list[str]) -> int:

    res = 0
    for i in range(len(matrix)):
        for j in range(len(matrix[0])):
            res += _solve_matrix_part2(matrix, i, j)
    return res


