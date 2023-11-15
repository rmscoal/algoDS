package com.example.random;

import static org.junit.jupiter.api.Assertions.assertArrayEquals;

import org.junit.jupiter.api.Test;

public class LargestPermutationTest {
  @Test
  void largestPermutationTest() {
    assertArrayEquals(new int[] { 3, 1, 2 }, LargestPermutation.largestPermutation(new int[] { 1, 3, 2 }, 1));
    assertArrayEquals(new int[] { 3, 2, 1 }, LargestPermutation.largestPermutation(new int[] { 1, 3, 2 }, 2));
    assertArrayEquals(new int[] { 3, 2, 1 }, LargestPermutation.largestPermutation(new int[] { 1, 2, 3 }, 2));
    assertArrayEquals(new int[] { 4, 2, 3, 1 }, LargestPermutation.largestPermutation(new int[] { 1, 2, 3, 4 }, 1));
    assertArrayEquals(new int[] { 4, 3, 2, 1 }, LargestPermutation.largestPermutation(new int[] { 1, 2, 3, 4 }, 2));
    assertArrayEquals(new int[] { 4, 3, 2, 1 }, LargestPermutation.largestPermutation(new int[] { 1, 2, 3, 4 }, 2));
  }
}
