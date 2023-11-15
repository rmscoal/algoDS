package com.example.random;

import java.util.Arrays;
import java.util.Comparator;

// https://youtu.be/bC7o8P_Ste4

public class DisjoinInterval {
  public static int disjointInterval(int[][] arr) {
    // Sort them by end interval
    // Given: [[1,4], [2,3], [8,9], [4,6]]
    Arrays.sort(arr, new Comparator<int[]>() {
      @Override
      public int compare(int[] a, int[] b) {
        return Integer.compare(a[1], b[1]);
      }
    });
    // We have: [[2,3], [1,4], [4,6], [8,9]]

    int count = 1;
    int prevEnd = arr[0][1];

    for (int[] el : arr) {
      if (el[0] <= prevEnd) {
        continue;
      } else {
        count++;
        prevEnd = el[1];
      }
    }

    return count;
  }
}
