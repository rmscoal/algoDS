package com.example.hackerrank;

import java.util.HashMap;
import java.util.List;

public class SalesByMatch {
  public static int sockMerchant(int n, List<Integer> arr) {
    HashMap<Integer, Integer> countMap = new HashMap<Integer, Integer>();

    int pairs = 0;

    arr.forEach((value) -> {
      if (countMap.get(value) == null) {
        countMap.put(value, 1);
      } else {
        countMap.put(value, countMap.get(value) + 1);
      }
    });

    for (int value : countMap.values()) {
      pairs += value / 2;
    }

    return pairs;
  }
}
