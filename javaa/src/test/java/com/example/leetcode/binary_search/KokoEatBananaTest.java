package com.example.leetcode.binary_search;

import java.util.List;
import java.util.ArrayList;
import org.junit.jupiter.api.Test;
import static org.junit.jupiter.api.Assertions.assertEquals;

class KokoEatBananaTest {
  public int minEatingSpeed(int[] piles, int h) {
    // We can count the amount of time koko eats the banana
    // by having to iterate through the piles.

    // Naive way is to find all the speed from 1 until the
    // max number of pile available in the given piles.
    // Hence, to shorten that amount, instead of linearly
    // try the speeds, we can do a binary search approach
    // instead

    // We can find the max pile
    int maxPile = 0;
    for (int pile : piles) {
      maxPile = Math.max(maxPile, pile);
    }

    int left = 1;
    int right = maxPile;
    int result = maxPile;
    while (left <= right) {
      int speed = (left + right) / 2;
      long time = 0;
      for (int pile : piles) {
        time += pile / speed;
        if (pile % speed != 0) {
          time += 1;
        }
      }

      if (time <= h) {
        result = Math.min(result, speed);
        right = speed - 1;
      } else {
        left = speed + 1;
      }
    }

    return result;
  }

  class TestCase {
    int[] piles;
    int h;
    int result;

    public TestCase(int[] piles, int h, int result) {
      this.piles = piles;
      this.h = h;
      this.result = result;
    }
  }

  @Test
  public void testKokoEatBanana() {
    List<TestCase> testcases = new ArrayList<>();

    testcases.add(new TestCase(new int[] { 3, 6, 7, 11 }, 8, 4));
    testcases.add(new TestCase(new int[] { 30, 11, 23, 4, 20 }, 5, 30));
    testcases.add(new TestCase(new int[] { 30, 11, 23, 4, 20 }, 6, 23));
    testcases.add(new TestCase(new int[] { 312884470 }, 312884469, 2));
    testcases.add(new TestCase(
        new int[] { 332484035, 524908576, 855865114, 632922376, 222257295, 690155293, 112677673, 679580077, 337406589,
            290818316, 877337160, 901728858, 679284947, 688210097, 692137887, 718203285, 629455728, 941802184 },
        823855818, 14));
    testcases.add(new TestCase(new int[] { 1, 1, 1, 999999999 }, 10, 142857143));
    testcases.add(new TestCase(new int[] { 805306368, 805306368, 805306368 }, 1000000000, 3));

    for (TestCase tc : testcases) {
      assertEquals(tc.result, new KokoEatBananaTest().minEatingSpeed(tc.piles, tc.h));
    }
  }
}
