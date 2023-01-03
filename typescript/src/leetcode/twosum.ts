/**
 * A class to solve Two Sum.
 *
 * Problem:
 * Given an array of integers nums and an integer target,
 * return indices of the two numbers such that they add up
 * to target.
 *
 * You may assume that each input would have exactly one
 * solution, and you may not use the same element twice.
 *
 * You can return the answer in any order.
 * @example
 * nums = [2,7,11,15], target = 9 // => [0,1]
 * Because nums[0] + nums[1] == 9, we return [0, 1].
 */
export default class TwoSum {
  /** @private */
  nums: Array<number>;

  /** @private */
  target: number;

  constructor(nums: Array<number>, target: number) {
    this.nums = nums;
    this.target = target;
  }

  /**
     * @method One-pass-Hash-Table:
     * Creates a hash map from the array where the keys are the value
     * in constructor's array and the value of the hash map is index
     * in constructor's array. This is done by iterating through the
     * array while also find whether the target subtracted by the
     * hash map key's. If there exist then return their indexes.
     *
     * Time-complexity: O(n).
     * Space-complexity: O(n).
     * @returns {Array<number>}
     */
  Solver(): Array<number> {
    const obj: Object = {};

    for (let i = 0; i < this.nums.length; i += 1) {
      if (Object.prototype.hasOwnProperty.call(obj, this.target - this.nums[i])) {
        return [obj[this.target - this.nums[i]], i];
      }

      obj[this.nums[i]] = i;
    }

    return null;
  }

  /**
     * @method Brute-force:
     * Iterate a double loop to find the all the values in the array
     * that when sum is equal to the target.
     *
     * Time-complexity: O(n^2).
     * Space-complexity: O(1).
     * @returns {Array<number>}
     */
  SlowSolver(): Array<number> {
    let pairOne: number = -1;
    let pairTwo: number = -1;

    for (let i = 0; i < this.nums.length; i += 1) {
      for (let j = 0; j < this.nums.length; j += 1) {
        if (this.nums[i] + this.nums[j] === this.target) {
          pairOne = i;
          pairTwo = j;
          break;
        }
      }
    }

    return [pairOne, pairTwo];
  }

  /**
   * @method
   * Iterate through the params and finds the two values
   * whose sum is equal to target.
   *
   * Time-complexity: O(n)
   * Space-complexity: O(1)
   * @returns {Array<number>}
   */
  FastSolver(): Array<number> {
    for (let i = 0; i < this.nums.length; i += 1) {
      if (this.nums.includes(this.target - this.nums[i])) {
        const tempIndex: number = this.nums.indexOf(this.target - this.nums[i]);
        if (tempIndex !== i) {
          return [tempIndex, i];
        }
      }

      if (this.nums.includes(this.nums[i] - this.target)) {
        const tempIndex: number = this.nums.indexOf(this.nums[i] - this.target);
        if (tempIndex !== i) {
          return [tempIndex, i];
        }
      }
    }

    return null;
  }
}
