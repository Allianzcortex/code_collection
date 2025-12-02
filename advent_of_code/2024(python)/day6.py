
from day4 import prefix_dir

# def _traverse_matrix_part1(matrix:list[str], x:int, y:int, s:set, dir:str):

#     if matrix[x][y] != 'X':
#         s.add((x,y))
    
#     print((x,y))
#     # check new position
#     new_x, new_y = x+prefix_dir[dir][0], y+prefix_dir[dir][1]
#     if new_x<0 or new_x>len(matrix)-1 or new_y<0 or new_y>len(matrix[0])-1:
#         # have been out of matrix
#         return

#     new_dir = dir
#     if matrix[new_x][new_y] == '.':
#         _traverse_matrix_part1(matrix,new_x,new_y,s,new_dir)
#         return

#     if matrix[new_x][new_y] == '#':
#         # meet obstacle, turning 90 degrees
#         if(dir=="up"):
#             new_dir = "right"
#         elif(dir=="right"):
#             new_dir = "down"
#         elif(dir=="down"):
#             new_dir = "left"
#         else:
#             new_dir = "up"
    
#     new_x, new_y = x+prefix_dir[new_dir][0], y+prefix_dir[new_dir][1]
#     if new_x<0 or new_x>len(matrix)-1 or new_y<0 or new_y>len(matrix[0])-1:
#         return
    
#     _traverse_matrix_part1(matrix,new_x,new_y,s,new_dir)

def _traverse_matrix_part1(matrix:list[str], x:int, y:int, s:set, dir:str):

    while True:

        if matrix[x][y] != 'X':
            s.add((x,y))
    
        print((x,y))
    # check new position
        new_x, new_y = x+prefix_dir[dir][0], y+prefix_dir[dir][1]
        if new_x<0 or new_x>len(matrix)-1 or new_y<0 or new_y>len(matrix[0])-1:
            # have been out of matrix
            break

        
        if matrix[new_x][new_y] == '.' or matrix[new_x][new_y]=='^':
            x,y = new_x, new_y
        else:
            # meet obstacle, turning 90 degrees
            new_dir = dir
            if(dir=="up"):
                new_dir = "right"
            elif(dir=="right"):
                new_dir = "down"
            elif(dir=="down"):
                new_dir = "left"
            else:
                new_dir = "up"
    
            new_x, new_y = x+prefix_dir[new_dir][0], y+prefix_dir[new_dir][1]
            if new_x<0 or new_x>len(matrix)-1 or new_y<0 or new_y>len(matrix[0])-1:
                break
    
            x,y = new_x, new_y
            dir = new_dir

def solve_day6_part1(matrix:list[str]) -> int:
    visited_set = set()
    start_x, start_y = 0, 0
    
    # firstly decide where the guard is
    for i in range(len(matrix)):
        for j in range(len(matrix[0])):
            if matrix[i][j] == '^':
                start_x, start_y = i, j
    
    # next iterate through the whole matrix until out of matrix
    _traverse_matrix_part1(matrix,start_x,start_y,visited_set,"up")

    return len(visited_set)

