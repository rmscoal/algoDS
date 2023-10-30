package com.example.hackerrank;

import java.util.List;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.HashMap;
import java.util.concurrent.TimeUnit;

import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.Tag;
import org.junit.jupiter.api.Timeout;
import static org.junit.jupiter.api.Assertions.assertEquals;

class ArrayManipulationTestCase {
  public int n;
  public List<List<Integer>> queries;
  public long answer;

  public ArrayManipulationTestCase(int n, List<List<Integer>> queries, Long answer) {
    this.n = n;
    this.queries = queries;
    this.answer = answer;
  }
}

public class ArrayManipulationTest {
  @Test
  @Tag("Optimized")
  @Timeout(value = 42, unit = TimeUnit.MILLISECONDS)
  void arrayManipulationTest() {
    HashMap<String, ArrayManipulationTestCase> tests = new HashMap<String, ArrayManipulationTestCase>();

    // Test Case 1
    List<List<Integer>> arr1 = new ArrayList<>();
    arr1.add(Arrays.asList(1, 5, 3));
    arr1.add(Arrays.asList(4, 8, 7));
    arr1.add(Arrays.asList(6, 9, 1));
    tests.put("test_1", new ArrayManipulationTestCase(10, arr1, 10l));

    tests.forEach((key, test) -> {
      long result = ArrayManipulation.solution(test.n, test.queries);
      assertEquals(test.answer, result);
    });
  }
}
