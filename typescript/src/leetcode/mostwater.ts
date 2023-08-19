/* eslint-disable no-plusplus */
/* eslint-disable max-len */
export default class MostWater {
  public static Solver(heights: number[]): number {
    let max = 0;
    for (let i = 0; i < heights.length; i += 1) {
      for (let j = i + 1; j < heights.length; j += 1) {
        const tmp = heights[i] >= heights[j] ? heights[j] * (j - i) : heights[i] * (j - i);
        if (tmp > max) {
          max = tmp;
        }
      }
    }

    return max;
  }

  public static PerfSolver(heights: number[]): number {
    let max = 0;
    let i = 0;
    let j = heights.length - 1;

    while (i < j) {
      const tmp = Math.min(heights[i], heights[j]) * (j - i);
      max = tmp > max ? tmp : max;

      if (heights[i] < heights[j]) {
        i++;
      } else if (heights[i] > heights[j]) {
        j--;
      } else {
        i++;
      }
    }

    return max;
  }
}
