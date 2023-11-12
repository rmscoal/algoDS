package com.example.hackerrank;

import java.util.Arrays;

public class MinimumAbsoluteDiff {
  public static int minimumAbsoluteDiff(int[] list) {
    Arrays.sort(list);
    int min = Integer.MAX_VALUE;

    for (int i = 0; i < list.length - 1; i++) {
      min = Math.min(min, Math.abs(list[i] - list[i + 1]));
    }

    return min;
  }
}
