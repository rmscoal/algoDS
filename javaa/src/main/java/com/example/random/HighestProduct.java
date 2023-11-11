package com.example.random;

import java.util.Arrays;

public class HighestProduct {
  public static int highestProductSlow(int[] nums) {
    int max = Integer.MIN_VALUE;
    for (int i = 0; i < nums.length; i++) {
      for (int j = i + 1; j < nums.length; j++) {
        for (int k = j + 1; k < nums.length; k++) {
          if (nums[i] * nums[j] * nums[k] > max) {
            max = nums[i] * nums[j] * nums[k];
          }
        }
      }
    }

    return max;
  }

  public static int highestProduct(int[] nums) {
    Arrays.sort(nums);

    int twoLowestOneHighest = nums[0] * nums[1] * nums[nums.length - 1];
    int threeHighest = nums[nums.length - 1] * nums[nums.length - 2] * nums[nums.length - 3];

    return Math.max(twoLowestOneHighest, threeHighest);
  }
}
