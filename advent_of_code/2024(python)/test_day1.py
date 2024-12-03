import unittest

from day1 import solve_day1_part1, solve_day1_part2

question_input = [
"3, 4",
"4, 3",
"2, 5",
"1, 3",
"3, 9",
"3, 3",]

class TestSample(unittest.TestCase):

    def test_part1(self):
        self.assertEqual(solve_day1_part1(question_input), 11)

    def test_part2(self):
        self.assertEqual(solve_day1_part2(question_input), 31)

if __name__ == "__main__":
    unittest.main()
