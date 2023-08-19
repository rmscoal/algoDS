export default class FirstAndLast {
  public static Solver(nums: number[], target: number): number[] {
    const res = [];
    for (let i = 0, j = nums.length - 1; i <= j; i += 1, j -= 1) {
      if (nums[i] === target) res.unshift(i);
      if (nums[j] === target) res.unshift(j);
    }

    if (res.length >= 2) {
      return [res[0], res[res.length - 1]];
    }

    if (res.length === 1) {
      return [res[0], res[0]];
    }

    return [-1, -1];
  }
}
