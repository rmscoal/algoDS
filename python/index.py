from leetcode.pangramsentence import PangramSentence
from leetcode.twostringarray import TwoStringArray
from leetcode.stayingingrid import StayingInGrid

if __name__ == "__main__":
  print("Pangram Sentence Solver: ", PangramSentence.Solver("thequickbrownfoxjumpsoverthelazydog"))
  print("Equivalence of Two String Array: ", TwoStringArray.Solver(["ab", "c"], ["a", "bc"]))
  print("Staying In Grid: ", StayingInGrid.FastSolver(3, [0,1], "RRDDLU"))