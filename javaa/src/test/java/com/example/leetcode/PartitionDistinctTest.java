package com.example.leetcode;

import static org.junit.jupiter.api.Assertions.assertEquals;

import org.junit.jupiter.api.Test;

public class PartitionDistinctTest {
  public int partitionDisjoint(int[] nums) {
    int i = 0;
    int currentMax = nums[i], traversedMax = nums[i];
    for (int j = 1; j < nums.length; j++) {
      if (nums[j] < currentMax) {
        i = j;
        currentMax = traversedMax;
      }

      traversedMax = Math.max(traversedMax, nums[j]);
    }

    return i + 1;
  }

  @Test
  public void partitionDistinctTest() {
    PartitionDistinctTest pd = new PartitionDistinctTest();

    assertEquals(pd.partitionDisjoint(new int[]{32, 45, 57, 24, 19, 0, 24, 49, 67, 87, 87}), 8);
  }
}
