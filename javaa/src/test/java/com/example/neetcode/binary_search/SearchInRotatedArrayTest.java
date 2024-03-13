package com.example.neetcode.binary_search;

import static org.junit.jupiter.api.Assertions.assertEquals;

import java.util.ArrayList;
import java.util.List;

import org.junit.jupiter.api.Test;

class SearchInRotatedArrayTest {
  public int search(int[] nums, int target) {
    int left = 0;
    int right = nums.length - 1;

    while (left <= right) {
      int mid = (left + right) / 2;

      if (nums[mid] == target) { // If they match, return immediately
        return mid;
      } else if (target > nums[mid] && target <= nums[right]) {
        System.out.println("I AM GOING RIGHT");
        left = mid + 1;
      } else if (target > nums[mid] && nums[mid] > nums[right]) {
        System.out.println("I AM GOING RIGHT");
        left = mid + 1;
      } else if (target > nums[mid] && nums[mid] < nums[left]) {
        System.out.println("I AM GOING LEFT");
        right = mid - 1;
      } else if (target < nums[mid] && target >= nums[left]) {
        System.out.println("I AM GOING LEFT");
        right = mid - 1;
      } else if (target < nums[mid] && nums[mid] < nums[right]) {
        System.out.println("I AM GOING LEFT");
        right = mid - 1;
      } else {
        System.out.println("I AM GOING RIGHT");
        left = mid + 1;
      }

      System.out.println("LEFT: " + left + " RIGHT: " + right);
    }

    return -1;
  }

  class TestCase {
    int[] nums;
    int target;
    int result;

    public TestCase(int[] nums, int target, int result) {
      this.nums = nums;
      this.target = target;
      this.result = result;
    }
  }

  @Test
  void testSearch() {
    List<TestCase> testcases = new ArrayList<>();
    testcases.add(new TestCase(
        new int[] { 4, 5, 6, 7, 0, 1, 2, 3 },
        0,
        4));
    testcases.add(new TestCase(
        new int[] { 4, 5, 6, 7, 8, 1, 2, 3 },
        8,
        4));
    testcases.add(new TestCase(
        new int[] { 7, 8, 4, 5, 6 },
        8,
        1));
    testcases.add(new TestCase(
        new int[] { 4, 5, 6, 7, 0, 1, 2, 3 },
        8,
        -1));
    testcases.add(new TestCase(
        new int[] { 3, 1 },
        1,
        1));
    testcases.add(new TestCase(
        new int[] { 10, 1, 2, 3, 4, 5, 6, 7 },
        5,
        5));
    testcases.add(new TestCase(
        new int[] { 10, 1, 2, 3, 4, 5, 6, 7 },
        1,
        1));

    for (TestCase tc : testcases) {
      int result = new SearchInRotatedArrayTest().search(tc.nums, tc.target);
      assertEquals(tc.result, result);
    }
  }
}
