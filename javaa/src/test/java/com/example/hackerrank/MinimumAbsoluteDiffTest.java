package com.example.hackerrank;

import org.junit.jupiter.api.Test;
import java.util.List;
import java.util.ArrayList;

class MinimumAbsoluteDiffTestCase {
  public int[] list;
  public int answer;

  MinimumAbsoluteDiffTestCase(int[] list, int answer) {
    this.list = list;
    this.answer = answer;
  }
}

public class MinimumAbsoluteDiffTest {
  @Test
  void minimumAbsoluteDiffTest() {
    List<MinimumAbsoluteDiffTestCase> tests = new ArrayList<MinimumAbsoluteDiffTestCase>();
    tests.add(new MinimumAbsoluteDiffTestCase(new int[] { 1, 2, 4, 5, 6, 7, 8 }, 1));
  }
}
