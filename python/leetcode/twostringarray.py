class TwoStringArray:

  @classmethod
  def Solver(cls, word1, word2) -> bool:
    # Joining the elements in a list:
    # res_one = "".join(map(str, word1))
    # res_two = "".join(map(str, word2))
    # 
    # or
    word1 = "".join(word1)
    word2 = "".join(word2)

    return word1 == word2