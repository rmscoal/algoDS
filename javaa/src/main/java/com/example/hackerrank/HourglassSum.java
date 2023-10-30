package com.example.hackerrank;

import java.util.List;

public class HourglassSum {
  public static int hourglassSum(List<List<Integer>> arr) {
    int up = 0;
    int mid = 1;
    int down = 2;
    int max = Integer.MIN_VALUE;

    while (up < 4 && mid < 5 && down < 6) {
      int i = 0;
      int j = 1;

      while (i <= 3 && j <= 4) {
        int sum = 0;
        sum += arr.get(up).get(i) + arr.get(up).get(i + 1) + arr.get(up).get(i + 2)
            + arr.get(mid).get(j)
            + arr.get(down).get(i) + arr.get(down).get(i + 1) + arr.get(down).get(i + 2);

        if (max < sum) {
          max = sum;
        }

        i++;
        j++;
      }

      up++;
      mid++;
      down++;
    }

    return max;
  }
}
