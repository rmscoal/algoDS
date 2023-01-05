/**
 * A class to solve Split a String in Balanced Strings (Easy).
 *
 * Problem:
 * Balanced strings are those that have an equal quantity of
 * 'L' and 'R' characters. Given a balanced string s, split
 * it into some number of substrings such that:
 * - Each substring is balanced.
 * Return the maximum number of balanced strings you can obtain.
 */
export default class BalancedStrings {
    static Solver(s) {
        // Returned variable:
        let result = 0;
        // Holds the previous index for starting to balance.
        let prevIndex = 0;
        // Holds how many are to balanced out.
        let currentCount = 1;
        for (let i = 1; i < s.length; i += 1) {
            if (s[i] !== s[prevIndex]) {
                currentCount -= 1;
                if (currentCount === 0) {
                    result += 1;
                    prevIndex = i + 1;
                }
            }
            else {
                currentCount += 1;
            }
        }
        return result;
    }
}
//# sourceMappingURL=balancedstrings.js.map