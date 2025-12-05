import unittest

from day5 import solve_day5_part1, solve_day5_part2

question_input = [
    "3-5",
    "10-14",
    "16-20",
    "12-18",
    "",
    "1",
    "5",
    "8",
    "11",
    "17",
    "32",
]


class TestSample(unittest.TestCase):

    def test_part1(self):
        self.assertEqual(solve_day5_part1(question_input), 3)

    def test_part2(self):
        self.assertEqual(solve_day5_part2(question_input), 14)


if __name__ == "__main__":
    unittest.main()
