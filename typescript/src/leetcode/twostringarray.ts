/**
 * A class to solve "Check If Two String Array are Equivalent".
 *
 * Problem:
 * Given two string arrays word1 and word2, return true if the two arrays represent
 * the same string, and false otherwise. A string is represented by an array if the
 * array elements concatenated in order forms the string.
 */
export default class TwoStringArray {
  /**
   * @public
   * Solver
   *
   * Time-complexity: O(n)
   * Space-complexity: O(1)
   * @param word1
   * @param word2
   * @returns the equivalence of the two string array
   */
  public static Solver(word1: Array<string>, word2: Array<string>): boolean {
    const stringOne = word1.join('');
    const stringTwo = word2.join('');

    return stringOne === stringTwo;
  }
}
