package com.example.hackerrank;

import static org.junit.jupiter.api.Assertions.assertEquals;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

import org.junit.jupiter.api.Test;

class JumpOnCloudsTestCase {
  public List<Integer> c;
  public int answer;

  public JumpOnCloudsTestCase(List<Integer> c, int answer) {
    this.c = c;
    this.answer = answer;
  }
}

public class JumpOnCloudsTest {
  @Test
  void jumpingOnCloudsTest() {
    ArrayList<JumpOnCloudsTestCase> tests = new ArrayList<JumpOnCloudsTestCase>();
    tests.add(new JumpOnCloudsTestCase(Arrays.asList(0, 0, 1, 0, 0, 1, 0), 4));
    tests.add(new JumpOnCloudsTestCase(Arrays.asList(0, 0, 0, 0, 1, 0), 3));
    tests.add(new JumpOnCloudsTestCase(Arrays.asList(0, 0, 0, 1, 0, 0), 3));

    tests.forEach((test) -> {
      int result = JumpOnClouds.jumpingOnClouds(test.c);
      assertEquals(test.answer, result);
    });
  }
}
