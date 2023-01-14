/**
 * A class to solve Sorting the Sentence (string, easy)
 *
 * Problem:
 * A sentence is a list of words that are separated by a
 * single space with no leading or trailing spaces. Each
 * word consists of lowercase and uppercase English letters.
 *
 * A sentence can be shuffled by appending the 1-indexed word
 * position to each word then rearranging the words in the sentence.
 *
 * Example:
 * the sentence "This is a sentence" can be shuffled as "sentence4
 * a3 is2 This1" or "is2 sentence4 This1 a3".
 *
 * Given a shuffled sentence s containing no more than 9 words,
 * reconstruct and return the original sentence.
 */
export default class SentenceSorting {
    /**
     * Solver
     * Split the input string into an array. Then create an empty
     * fixed length array of string. Do a foor loop at the splitted
     * input string and insert according to the index attached to
     * the string.
     *
     * Time-complexity: O(n)
     * Space-complexity: O(n)
     * @param s the input string
     * @returns reconstruted string
     * @example
     * s = "is2 sentence4 This1 a3"
     * result = "This is a sentence"
     */
    static Solver(s) {
        const sAsArray = s.split(' ');
        const resultAsArray = new Array(sAsArray.length);
        for (let i = 0; i < sAsArray.length; i += 1) {
            const index = Number(sAsArray[i].charAt(sAsArray[i].length - 1));
            resultAsArray[index - 1] = sAsArray[i].substring(0, sAsArray[i].length - 1);
        }
        return resultAsArray.join(' ');
    }
}
//# sourceMappingURL=sentencesorting.js.map