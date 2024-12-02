import unittest

from day2 import solve_day2_part1, solve_day2_part2

question2_input = [
"7 6 4 2 1",
"1 2 7 8 9",
"9 7 6 2 1",
"1 3 2 4 5",
"8 6 4 4 1",
"1 3 6 7 9",
]


class TestDay2(unittest.TestCase):

    def test_part1(self):
        self.assertEqual(solve_day2_part1(question2_input), 2)

    def test_part2(self):
        self.assertEqual(solve_day2_part2(question2_input), 4)

if __name__ == "__main__":
    unittest.main()
