import unittest

from day4 import solve_day4_part1, solve_day4_part2


matrix_day4 = [
"MMMSXXMASM",
"MSAMXMSMSA",
"AMXSXMAAMM",
"MSAMASMSMX",
"XMASAMXAMM",
"XXAMMXXAMA",
"SMSMSASXSS",
"SAXAMASAAA",
"MAMMMXMMMM",
"MXMXAXMASX",
]



class TestDay4(unittest.TestCase):

    def test_part1(self):
        self.assertEqual(solve_day4_part1(matrix_day4), 18)

    def test_part2(self):
        self.assertEqual(solve_day4_part2(matrix_day4), 9)

if __name__ == "__main__":
    unittest.main()
