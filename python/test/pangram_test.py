import unittest
from leetcode import PangramSentence

class PangramSentenceTestCase(unittest.TestCase):
  def pangramTest(self):
    self.assertEqual(PangramSentence.Solver("thequickbrownfoxjumpsoverthelazydog"), true)

if __main__ == 'main':
  unittest.main()