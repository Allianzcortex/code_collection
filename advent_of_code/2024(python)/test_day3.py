import unittest

from day3 import solve_day3_part1, solve_day3_part2


instruction1 = "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"

instruction2 = "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"


class TestDay3(unittest.TestCase):

    def test_part1(self):
        self.assertEqual(solve_day3_part1(instruction1), 161)

    def test_part2(self):
        self.assertEqual(solve_day3_part2(instruction2), 48)

if __name__ == "__main__":
    unittest.main()
