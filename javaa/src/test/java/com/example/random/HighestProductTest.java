package com.example.random;

import static org.junit.jupiter.api.Assertions.assertEquals;

import java.util.ArrayList;
import java.util.List;

import org.junit.jupiter.api.Test;

class HighestProductTestCase {
  public int[] nums;
  public int answer;

  HighestProductTestCase(int[] nums, int answer) {
    this.nums = nums;
    this.answer = answer;
  }
}

public class HighestProductTest {
  @Test
  void highestProductTest() {
    List<HighestProductTestCase> tests = new ArrayList<HighestProductTestCase>();
    tests.add(new HighestProductTestCase(new int[] { 1, 2, 3, 4 }, 24));
    tests.add(new HighestProductTestCase(new int[] { 0, -1, 10, 7, 5 }, 350));
    tests.add(new HighestProductTestCase(new int[] { -5, -2, -1, 0, 0, 3, 4, 5 }, 60));
    tests.add(new HighestProductTestCase(new int[] { -5, -5, 1, 0, 3, 4, 5 }, 125));
    tests.add(new HighestProductTestCase(new int[] { -5, 1, 0, 3, 4, -5, 5 }, 125));
    tests.add(new HighestProductTestCase(new int[] { -5, 1, 0, 3, 4, 5, 5 }, 100));

    for (HighestProductTestCase test : tests) {
      assertEquals(test.answer, HighestProduct.highestProductSlow(test.nums));
    }
    for (HighestProductTestCase test : tests) {
      assertEquals(test.answer, HighestProduct.highestProduct(test.nums));
    }
  }
}
