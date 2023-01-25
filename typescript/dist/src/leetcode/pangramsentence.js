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
    /**
     * Time-complexity: O(n^2)
     * Space-complexity: O(n)
     * @param sentence contains only lowercase string English letters
     * without space
     * @example
     * Input: sentence = "thequickbrownfoxjumpsoverthelazydog"
     * Output: true
     * Explanation: sentence contains at least one of every letter of
     * the English alphabet.
     */
    static Solver(sentence) {
        const letters = [];
        for (let i = 0; i < sentence.length; i += 1) {
            if (!letters.includes(sentence[i])) {
                letters.push(sentence[i]);
            }
        }
        return letters.length === this.ALPHABET_LENGTH;
    }
    /**
     * OutOfTheBoxSolver
     *
     * Time-complexity: O(n)
     * Space-complexity: O(n)
     * @param sentence
     * @returns whether the sentence contains all 26 letter alphabet
     */
    static OutOfTheBoxSolver(sentence) {
        // If sentence.length is less than 26, return false immediately
        if (sentence.length < 26) {
            return false;
        }
        const sentenceAsCharCodeArray = sentence.split('').map((value) => value.charCodeAt(0)).sort();
        for (let i = 0, j = 1; j < sentenceAsCharCodeArray.length; i += 1, j += 1) {
            if (sentenceAsCharCodeArray[j] > sentenceAsCharCodeArray[i] + 2) {
                return false;
            }
        }
        return true;
    }
    static FastSolver(sentence) {
        if (sentence.length < 26)
            return false;
        const sentenceAsSet = new Set(sentence);
        return sentenceAsSet.size === 26;
    }
}
PangramSentence.ALPHABET_LENGTH = 26;
//# sourceMappingURL=pangramsentence.js.map