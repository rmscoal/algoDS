package com.example.hackerrank;

import static org.junit.jupiter.api.Assertions.assertEquals;

import java.util.ArrayList;

import org.junit.jupiter.api.Test;

class CountingValleyTestCase {
  public String path;
  public int answer;

  // Constructor
  CountingValleyTestCase(String path, int answer) {
    this.path = path;
    this.answer = answer;
  }
}

public class CountingValleyTest {
  @Test
  void countingValleyTest() {
    ArrayList<CountingValleyTestCase> tests = new ArrayList<CountingValleyTestCase>();
    tests.add(new CountingValleyTestCase("UDDDUDUU", 1));
    tests.add(new CountingValleyTestCase("DDUUUUDD", 1));

    for (CountingValleyTestCase test : tests) {
      int result = CountingValleys.countingValleys(test.path.length(), test.path);
      assertEquals(test.answer, result);
    }
  }
}
