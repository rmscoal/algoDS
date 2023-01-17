/**
 * A class to solve Number of Arithmetic Triplets Leetcode
 * Problem:
 * You are given a 0-indexed, strictly increasing integer
 * array nums and a positive integer diff. A triplet (i, j, k)
 * is an arithmetic triplet if the following conditions are met:
 * - i < j < k
 * - nums[j] - nums[i] = diff
 * - nums[k] - nums[j] = diff
 * @example
 * Input: nums = [0,1,4,6,7,10], diff = 3
 * Output: 2
 * Explanation:
 * (1, 2, 4) is an arithmetic triplet because both 7 - 4 == 3 and 4 - 1 == 3.
 * (2, 4, 5) is an arithmetic triplet because both 10 - 7 == 3 and 7 - 4 == 3.
 * @constraint
 * 3 <= nums.length <= 200,
 * 0 <= nums[i] <= 200,
 * 1 <= diff <= 50,
 * nums is strictly increasing.
 * 
 * Time-complexity = O(n)
 * Space-complexity = O(1)
 */
export default class ArithmeticTriplets {
  /**
   * @method Solver
   * @param {Array<number>} nums defines the input strictly increasing array
   * @param {int} diff defines the difference to obliged
   * @returns {number} the number of unique arithmetic triplets
   * 
   * Time-complexity = O(n),
   * Space-complexity = O(1)
   */
  static Solver(nums, diff) {
    // Returned variable
    let count = 0

    console.debug("Nums:", nums)

    // Loop through the array and find the unique triplets
    for (let i = 0; i < nums.length; i++) {
      console.debug("Current i:", i)
      console.debug("Current nums[i]:", nums[i])
      if (nums.includes(nums[i] + diff)) {
        console.debug("There includes in the array nums[i] + diff")
        let indexOfNewVar = nums.indexOf(nums[i] + diff, i)
        console.debug("The targeted index:", indexOfNewVar, "with value of:", nums[indexOfNewVar])
        if (nums.includes(nums[indexOfNewVar] + diff)) {
          console.debug("There includes in the array nums[indexOfNewVar] + diff")
          count++
        }
      }
    }

    return count
  }
}
