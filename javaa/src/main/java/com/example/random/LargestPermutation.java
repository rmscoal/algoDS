package com.example.random;

import java.util.HashMap;

public class LargestPermutation {
  public static int[] largestPermutation(int[] arr, int swaps) {
    // The array is random permutation of number from 1 to N
    // and we are guaranteed that there are no duplicates of
    // integers in arr. So to swap and give the higher result
    // we can swap it with numbers starting from N = len(arr)
    // till our swaps is finished and then decrement that N.
    int idx = 0;
    int N = arr.length;

    // We use hashmap to keep the index of the numbers.
    HashMap<Integer, Integer> map = new HashMap<>();

    for (int i = 0; i < arr.length; i++) {
      map.put(arr[i], i);
    }

    while (idx < arr.length && swaps > 0) {
      if (arr[idx] < N) {
        int tmp = arr[idx];
        arr[idx] = arr[map.get(N)];
        arr[map.get(N)] = tmp;
        N--;
        swaps--;
      } else {
        idx++;
      }
    }

    return arr;
  }
}
