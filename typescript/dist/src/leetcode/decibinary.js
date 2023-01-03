/**
 * A class to solve Decibinary
 *
 * Problem:
 * A decimal number is called deci-binary if each of
 * its digits is either 0 or 1 without any leading zeros.
 * For example, 101 and 1100 are deci-binary, while 112
 * and 3001 are not.
 *
 * @example
 * const d = new Decibinary("123")
 * d.Solver() # 3
 */
export default class Decibinary {
    /**
     * @constructor
     * @param s as the string input
     */
    constructor(s) {
        this.s = s;
        this.max = 0;
    }
    /**
     * @method
     * Iterate through the array and find the largest
     * integer in the string. It would be the minimum
     * decibinary partition.
     *
     * Time-complexity: O(n)
     * Space-complexity: O(1)
     *
     * @returns {number}
     */
    Solver() {
        for (let i = 0; i < this.s.length; i += 1) {
            const temp = parseInt(this.s[i], 10);
            if (this.max <= parseInt(this.s[i], 10)) {
                this.max = temp;
            }
        }
        return this.max;
    }
    /**
     * @method
     * Iterate from 9 until 0 descendingly and returns when
     * the iterator value is included in the string.
     *
     * Time-complexity: O(10)
     * Space-complexity: O(1)
     * @returns {number}
     */
    FastSolver() {
        for (let i = 9; i >= 0; i -= 1) {
            if (this.s.includes(i.toString())) {
                return i;
            }
        }
        return null;
    }
}
//# sourceMappingURL=decibinary.js.map