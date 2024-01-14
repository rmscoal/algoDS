package com.example.leetcode;

import java.util.Stack;

import static org.junit.jupiter.api.Assertions.assertArrayEquals;

import java.util.ArrayList;
import java.util.List;

import org.junit.jupiter.api.Test;

class DailyTemperatureTest {

  class TestCase {
    public int[] temperatures;
    public int[] result;

    public TestCase(int[] temperatures, int[] result) {
      this.temperatures = temperatures;
      this.result = result;
    }
  }

  public int[] dailyTemperatures(int[] temperatures) {
    // 1. Append the first temp to stack
    // 2. Iterate through the stack,
    //
    // a. Then iterate each stack
    // if the current temp top of the stack
    // is more then current temp, we add to the stack
    // if not, we pop and get the index diff.
    Stack<Integer> stack = new Stack<>();
    int[] result = new int[temperatures.length];

    stack.add(0);
    for (int i = 0; i < temperatures.length; i++) {
      while (!stack.empty()) {
        int idx = stack.getLast();
        if (temperatures[idx] >= temperatures[i])
          break;
        result[idx] = i - idx;
        stack.pop();
      }
      stack.add(i);
    }

    return result;
  }

  @Test
  void testDailyTemperature() {
    DailyTemperatureTest dtTest = new DailyTemperatureTest();
    List<DailyTemperatureTest.TestCase> testcases = new ArrayList<>();
    testcases.add(new DailyTemperatureTest.TestCase(new int[] { 73, 74, 75, 71, 69, 72, 76, 73 },
        new int[] { 1, 1, 4, 2, 1, 1, 0, 0 }));

    for (DailyTemperatureTest.TestCase tc : testcases) {
      assertArrayEquals(tc.result, dtTest.dailyTemperatures(tc.temperatures));
      assertArrayEquals(tc.result, dtTest.dailyTemperatures_FAST(tc.temperatures));
    }
  }

  public int[] dailyTemperatures_FAST(int[] T) {
    int n = T.length;
    int[] res = new int[n];
    int hottest = 0;
    for (int i = n - 1; i >= 0; i--) {
      int temp = T[i];
      if (temp >= hottest) {
        hottest = temp;
        continue;
      }
      int days = 1;
      while (T[i + days] <= temp) {
        days += res[i + days];
      }
      res[i] = days;
    }
    return res;
  }
}
