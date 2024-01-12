package com.example.leetcode;

import static org.junit.jupiter.api.Assertions.assertArrayEquals;

import org.junit.jupiter.api.Test;

import java.util.ArrayList;
import java.util.List;

class ProductExceptSelfTest {

  class TestCase {
    public int[] input;
    public int[] result;

    public TestCase(int[] input, int[] result) {
      this.input = input;
      this.result = result;
    }
  }

  public int[] productExceptSelf_TwoArrays(int[] nums) {
    // Example:
    // Input: [1,2,3,4]
    // Output: [24, 12, 8, 6]
    //
    // Solution:
    // prefix = [1, 0, ..., 0]
    // suffix = [0, 0, ..., 1]
    //
    //
    // Pre-Sum like Approach
    // Input: [1, 2, 3, 4]
    //
    // Iterate through nums from the second index to the last
    // Prefix: [1, 1, 2, 6]
    //
    // Iterate through nums from the second last index to the first
    // Suffix: [24, 12, 4, 1]
    //

    int[] prefix = new int[nums.length];
    int[] suffix = new int[nums.length];
    int[] result = new int[nums.length];
    prefix[0] = 1;
    suffix[nums.length - 1] = 1;

    for (int i = 1; i < nums.length; i++) {
      prefix[i] = prefix[i - 1] * nums[i - 1];
    }

    for (int i = nums.length - 2; i >= 0; i--) {
      suffix[i] = suffix[i + 1] * nums[i + 1];
    }

    for (int i = 0; i < nums.length; i++) {
      result[i] = prefix[i] * suffix[i];
    }

    return result;
  }

  public int[] productExceptSelf(int[] nums) {
    // Example:
    // Input: [1,2,3,4]
    // Output: [24, 12, 8, 6]
    //
    // Solution:
    // result = [1, 0, ..., 0]
    //
    // Pre-Sum like Approach
    // Input: [1, 2, 3, 4]
    //
    // Iterate through nums from the second index to the last
    // Result: [1, 1, 2, 6]
    //
    // Iterate through nums from the second last index to the first
    // Curr = 12
    // Result: [1, 12, 8, 6]
    //
    int[] result = new int[nums.length];
    result[0] = 1;

    // From left to right
    for (int i = 1; i < nums.length; i++) {
      result[i] = nums[i - 1] * result[i - 1];
    }

    int cursor = 1;
    for (int i = nums.length - 2; i >= 0; i--) {
      result[i] *= nums[i + 1] * cursor;
      cursor *= nums[i + 1];
    }

    return result;
  }

  @Test
  void test_productExceptSelf() {
    ProductExceptSelfTest cl = new ProductExceptSelfTest();

    List<TestCase> testcases = new ArrayList<>();
    testcases.add(new ProductExceptSelfTest.TestCase(new int[] { 1, 2, 3, 4 }, new int[] { 24, 12, 8, 6 }));
    testcases.add(new ProductExceptSelfTest.TestCase(new int[] { -1, 1, 0, -3, 3 }, new int[] { 0, 0, 9, 0, 0 }));

    for (ProductExceptSelfTest.TestCase tc : testcases) {
      assertArrayEquals(tc.result, cl.productExceptSelf_TwoArrays(tc.input));
      assertArrayEquals(tc.result, cl.productExceptSelf(tc.input));
    }
  }
}
