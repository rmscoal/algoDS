package com.example.hackerrank;

import java.util.HashMap;
import java.util.List;

import static org.junit.jupiter.api.Assertions.assertEquals;

import java.util.Arrays;

import org.junit.jupiter.api.Tag;
import org.junit.jupiter.api.Test;

class NewYearChaosTestCase {
  public List<Integer> q;
  public int answer;

  public NewYearChaosTestCase(List<Integer> q, int answer) {
    this.q = q;
    this.answer = answer;
  }
}

public class NewYearChaosTest {
  @Test
  void minimumBribesTest() {
    HashMap<String, NewYearChaosTestCase> tests = new HashMap<String, NewYearChaosTestCase>();

    tests.put("test_1", new NewYearChaosTestCase(Arrays.asList(2, 1, 5, 3, 4), 3));
    tests.put("test_2", new NewYearChaosTestCase(Arrays.asList(2, 5, 1, 3, 4), -1));
    tests.put("test_3", new NewYearChaosTestCase(Arrays.asList(1, 2, 5, 3, 7, 8, 6, 4), 7));

    tests.forEach((key, test) -> {
      System.out.println("New Year Chaos Test " + key);
      int result = NewYearChaos.minimumBribes(test.q);
      assertEquals(test.answer, result);
    });
  }

  @Test
  @Tag("slow")
  void minimumBribesTestFaster() {
    HashMap<String, NewYearChaosTestCase> tests = new HashMap<String, NewYearChaosTestCase>();

    tests.put("test_1", new NewYearChaosTestCase(Arrays.asList(2, 1, 5, 3, 4), 3));
    tests.put("test_2", new NewYearChaosTestCase(Arrays.asList(2, 5, 1, 3, 4), -1));
    tests.put("test_3", new NewYearChaosTestCase(Arrays.asList(1, 2, 5, 3, 7, 8, 6, 4), 7));

    tests.forEach((key, test) -> {
      System.out.println("New Year Chaos Test " + key);
      int result = NewYearChaos.minimumBribes(test.q);
      assertEquals(test.answer, result);
    });
  }
}
