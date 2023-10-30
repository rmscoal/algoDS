package com.example.hackerrank;

import java.util.List;

public class ArrayManipulation {
  public static long solution(int n, List<List<Integer>> queries) {
    long max = -1;
    int[] arr = new int[n + 2];

    queries.forEach((query) -> {
      arr[query.get(0)] += query.get(2);
      arr[query.get(1) + 1] -= query.get(2);
    });

    long tmp = 0;

    for (int i = 0; i < arr.length; i++) {
      tmp += arr[i];
      max = Math.max(max, tmp);
    }

    return max;
  }

}
