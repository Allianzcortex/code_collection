import unittest

from day4 import solve_day4_part1, solve_day4_part2

question_input = [
    "..@@.@@@@.",
    "@@@.@.@.@@",
    "@@@@@.@.@@",
    "@.@@@@..@.",
    "@@.@@@@.@@",
    ".@@@@@@@.@",
    ".@.@.@.@@@",
    "@.@@@.@@@@",
    ".@@@@@@@@.",
    "@.@.@@@.@.",
]


class TestSample(unittest.TestCase):

    def test_part1(self):
        self.assertEqual(solve_day4_part1(question_input), 13)

    def test_part2(self):
        self.assertEqual(solve_day4_part2(question_input), 43)


if __name__ == "__main__":
    unittest.main()
