import unittest

from day6 import solve_day6_part1


matrix_day6 = [
"....#.....",
".........#",
"..........",
"..#.......",
".......#..",
"..........",
".#..^.....",
"........#.",
"#.........",
"......#...",
]



class TestDay4(unittest.TestCase):

    def test_part1(self):
        self.assertEqual(solve_day6_part1(matrix_day6), 18)

    # def test_part2(self):
    #     self.assertEqual(solve_day4_part2(matrix_day4), 9)

if __name__ == "__main__":
    unittest.main()
