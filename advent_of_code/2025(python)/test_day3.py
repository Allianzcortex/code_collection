import unittest

from day3 import solve_day3_part1, solve_day3_part2

question_input = [
    "987654321111111",
    "811111111111119",
    "234234234234278",
    "818181911112111",
]


class TestSample(unittest.TestCase):

    def test_part1(self):
        self.assertEqual(solve_day3_part1(question_input), 357)

    def test_part2(self):
        self.assertEqual(solve_day3_part2(question_input), 3121910778619)


if __name__ == "__main__":
    unittest.main()
