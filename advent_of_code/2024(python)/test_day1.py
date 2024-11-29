import unittest

from day1 import add

class TestSample(unittest.TestCase):

    def test_add(self):
        self.assertEqual(add(2,4), 6)

if __name__ == "__main__":
    unittest.main()
