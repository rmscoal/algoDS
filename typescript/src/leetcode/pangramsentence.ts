/**
 * A class to solve Check if the Sentence is Pangram in Leetcode (easy)
 *
 * Problem:
 * A pangram is a sentence where every letter of the English alphabet
 * appears at least once. Given a string sentence containing only
 * lowercase English letters, return true if sentence is a pangram,
 * or false otherwise.
 */
export default class PangramSentence {
  private static ALPHABET_LENGTH = 26;

  /**
   * Time-complexity: O(n)
   * Space-complexity: O(n)
   * @param sentence contains only lowercase string English letters
   * without space
   * @example
   * Input: sentence = "thequickbrownfoxjumpsoverthelazydog"
   * Output: true
   * Explanation: sentence contains at least one of every letter of
   * the English alphabet.
   */
  public static Solver(sentence: string): boolean {
    const letters: Array<string> = [];

    for (let i = 0; i < sentence.length; i += 1) {
      if (!letters.includes(sentence[i])) {
        letters.push(sentence[i]);
      }
    }

    return letters.length === this.ALPHABET_LENGTH;
  }
}
