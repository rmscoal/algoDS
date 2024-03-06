package com.example.leetcode.binary_search;

import java.util.List;

import static org.junit.jupiter.api.Assertions.assertEquals;

import java.util.ArrayList;

import org.junit.jupiter.api.Test;

class FindMinimumInRotatedTest {
  public int findMin(int[] nums) {
    int left = 0;
    int right = nums.length - 1;
    int result = nums[0];

    while (left <= right) {
      int mid = (left + right) / 2;
      if (nums[left] < nums[right]) {
        result = Math.min(result, nums[left]);
        return result;
      } else if (nums[mid] < nums[right]) {
        right = mid - 1;
        result = Math.min(result, nums[mid]);
      } else {
        left = mid + 1;
        result = Math.min(result, nums[mid]);
      }
    }

    return result;
  }

  class TestCase {
    int[] nums;
    int result;

    public TestCase(int[] nums, int result) {
      this.nums = nums;
      this.result = result;
    }
  }

  @Test
  void testFindMin() {
    List<TestCase> testcases = new ArrayList<>();
    testcases.add(new TestCase(new int[] { 4, 5, 0, 1, 2 }, 0));

    for (TestCase tc : testcases) {
      int result = new FindMinimumInRotatedTest().findMin(tc.nums);
      assertEquals(tc.result, result);
    }
  }
}
