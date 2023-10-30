package com.example.hackerrank;

import java.util.List;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.HashMap;
import java.util.concurrent.TimeUnit;

import org.junit.jupiter.api.Tag;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.Timeout;

import static org.junit.jupiter.api.Assertions.assertEquals;

public class HourglassSumTest {
  @Test
  @Tag("Working Solution")
  @Timeout(value = 42, unit = TimeUnit.MILLISECONDS)
  void hourglassSumTest() {
    HashMap<String, List<List<Integer>>> tests = new HashMap<String, List<List<Integer>>>();
    HashMap<String, Integer> answers = new HashMap<String, Integer>();

    // Test 1
    List<List<Integer>> arr1 = new ArrayList<>();
    arr1.add(Arrays.asList(-9, -9, -9, 1, 1, 1));
    arr1.add(Arrays.asList(0, -9, 0, 4, 3, 2));
    arr1.add(Arrays.asList(-9, -9, -9, 1, 2, 3));
    arr1.add(Arrays.asList(0, 0, 8, 6, 6, 0));
    arr1.add(Arrays.asList(0, 0, 0, -2, 0, 0));
    arr1.add(Arrays.asList(0, 0, 1, 2, 4, 0));
    tests.put("test_1", arr1);
    answers.put("test_1", 28);

    // Test 2
    List<List<Integer>> arr2 = new ArrayList<>();
    arr2.add(Arrays.asList(1, 1, 1, 0, 0, 0));
    arr2.add(Arrays.asList(0, 1, 0, 0, 0, 0));
    arr2.add(Arrays.asList(1, 1, 1, 0, 0, 0));
    arr2.add(Arrays.asList(0, 0, 2, 4, 4, 0));
    arr2.add(Arrays.asList(0, 0, 0, 2, 0, 0));
    arr2.add(Arrays.asList(0, 0, 1, 2, 4, 0));
    tests.put("test_2", arr2);
    answers.put("test_2", 19);

    tests.forEach((key, matrix) -> {
      System.out.println("Testing Hourglass Sum " + key);
      int result = HourglassSum.hourglassSum(matrix);
      assertEquals(answers.get(key), result);
    });
  }
}
