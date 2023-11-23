package com.example.leetcode;

import java.util.HashMap;
import java.util.Map;

public class ContainsDuplicate {
  public static boolean containsDuplicateHash(int[] nums) {
    Map<Integer, Integer> map = new HashMap<Integer, Integer>();

    for (int i = 0; i < nums.length; i++) {
      if (map.get(nums[i]) != null) {
        return true;
      }

      map.put(nums[i], 1);
    }

    return false;
  }

  public static boolean containsDuplicate(int[] nums) {
    for (int i = 0; i < nums.length; i++) {
      int current = nums[i];
      int j = i - 1;

      while (j >= 0 && nums[j] > current) {
        nums[j + 1] = nums[j];
        j--;
      }

      if (j >= 0 && nums[j] == current)
        return true;

      nums[j + 1] = current;
    }

    return false;
  }

  public static boolean containsNearbyDuplicate(int[] nums, int k) {
    // Sliding window
    // 1. Make the window first
    // 2. Then walk through the nums keeping the window length.

    Map<Integer, Integer> map = new HashMap<Integer, Integer>();

    for (int i = 0; i < Math.min(nums.length, k + 1); i++) {
      map.put(nums[i], map.get(nums[i]) == null ? 1 : map.get(nums[i]) + 1);
      if (map.get(nums[i]) > 1) {
        return true;
      }
    }

    map.put(nums[0], map.get(nums[0]) - 1);
    int left = 1;
    int right = k + 1;

    while (right < nums.length) {
      map.put(nums[right], map.get(nums[right]) == null ? 1 : map.get(nums[right]) + 1);
      if (map.get(nums[right]) > 1) {
        return true;
      }

      map.put(nums[left], map.get(nums[left]) - 1);
      left++;
      right++;
    }

    return false;
  }

  public static boolean containsNearbyDuplicateV2(int[] nums, int k) {
    Map<Integer, Integer> map = new HashMap<Integer, Integer>();

    for (int i = 0; i < nums.length; i++) {
      map.put(nums[i], map.get(nums[i]) == null ? 1 : map.get(nums[i]) + 1);
      if (map.get(nums[i]) > 1) {
        return true;
      }

      if (i - k >= 0) {
        map.put(nums[i - k], map.get(nums[i - k]) - 1);
      }
    }

    return false;
  }

  public static boolean containsNearbyDuplicateV3(int[] nums, int k) {
    Map<Integer, Integer> map = new HashMap<Integer, Integer>();

    for (int i = 0; i < nums.length; i++) {
      if (map.get(nums[i]) != null) {
        int index = map.get(nums[i]);
        if (i - index <= k) {
          return true;
        }
      }

      map.put(nums[i], i);
    }

    return false;
  }
}
