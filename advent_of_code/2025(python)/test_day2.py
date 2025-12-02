import unittest

from day2 import solve_day2_part1_right_version

question_input = "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,\
1698522-1698528,446443-446449,38593856-38593862,565653-565659,\
824824821-824824827,2121212118-2121212124"


class TestSample(unittest.TestCase):
    def test_part1(self):
        self.assertEqual(solve_day2_part1_right_version(
            question_input), 1227775554)


if __name__ == "__main__":
    unittest.main()
