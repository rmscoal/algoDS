package com.example.leetcode;

import static org.junit.jupiter.api.Assertions.assertEquals;

import java.util.ArrayList;

import org.junit.jupiter.api.Test;

class MostCommonWordTestCase {
  public String paragraph;
  public String[] banned;
  public String answer;

  MostCommonWordTestCase(String paragraph, String[] banned, String answer) {
    this.paragraph = paragraph;
    this.banned = banned;
    this.answer = answer;
  }
}

public class MostCommonWordTest {
  @Test
  void mostCommonWordTest() {
    ArrayList<MostCommonWordTestCase> tests = new ArrayList<MostCommonWordTestCase>();
    tests.add(new MostCommonWordTestCase("Bob hit a ball, the hit BALL flew far after it was hit.",
        new String[] { "hit" }, "ball"));
    tests.add(new MostCommonWordTestCase("a.",
        new String[] {}, "a"));

    for (MostCommonWordTestCase test : tests) {
      assertEquals(test.answer, MostCommonWord.mostCommonWord(test.paragraph, test.banned));
    }
  }
}
