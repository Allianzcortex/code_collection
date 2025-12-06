import unittest

from day6 import solve_day6_part1, solve_day6_part2

question_input = [
    "123 328  51 64",
    " 45 64  387 23",
    "  6 98  215 314",
    "*   +   *   +  ",
]


class TestSample(unittest.TestCase):

    def test_part1(self):
        self.assertEqual(solve_day6_part1(question_input), 4277556)

    def test_part2(self):
        self.assertEqual(solve_day6_part2(question_input), 3263827)


if __name__ == "__main__":
    unittest.main()
