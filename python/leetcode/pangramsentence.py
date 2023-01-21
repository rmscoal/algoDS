class PangramSentence:

  LENGTH_OF_ALPHABET = 26

  @classmethod
  def Solver(cls, sentence):
    list_of_diff_alphabets_in_sentence = []

    for i in range(0, len(sentence)):
      if sentence[i] not in list_of_diff_alphabets_in_sentence:
        list_of_diff_alphabets_in_sentence.append(sentence[i])

    return len(list_of_diff_alphabets_in_sentence) == PangramSentence.LENGTH_OF_ALPHABET