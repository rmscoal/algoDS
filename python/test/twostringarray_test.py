import unittest
from leetcode import TwoStringArray

class TwoStringArrayTestCase(unittest.TestCase):
  def twostringarrayTest(self):
    self.assertEqual(TwoStringArray.Solver(["ab", "c"], ["a", "bc"]), true)
    self.assertEqual(TwoStringArray.Solver(["a", "cb"], ["ab", "c"]), false)
    self.assertEqual(TwoStringArray.Solver(["abc", "d", "defg"], ["abcddefg"]), true)

if __main__ == 'main':
  unittest.main()