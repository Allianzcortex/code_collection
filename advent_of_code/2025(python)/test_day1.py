import unittest

from day1 import solve_day1_part1, solve_day1_part2

question_input = [
    "L68",
    "L30",
    "R48",
    "L5",
    "R60",
    "L55",
    "L1",
    "L99",
    "R14",
    "L82",
]


class TestSample(unittest.TestCase):

    def test_part1(self):
        self.assertEqual(solve_day1_part1(question_input), 3)

    def test_part2(self):
        self.assertEqual(solve_day1_part2(question_input), 6)


if __name__ == "__main__":
    unittest.main()
