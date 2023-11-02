package com.example.leetcode;

import static org.junit.jupiter.api.Assertions.assertEquals;

import java.util.ArrayList;

import org.junit.jupiter.api.Test;

class ReverseWordStringTestCase {
  public String s;
  public String answer;

  ReverseWordStringTestCase(String s, String answer) {
    this.s = s;
    this.answer = answer;
  }
}

public class ReverseWordStringTest {
  @Test
  void reverseWordsTest() {
    ArrayList<ReverseWordStringTestCase> tests = new ArrayList<ReverseWordStringTestCase>();
    tests.add(new ReverseWordStringTestCase("the sky is blue", "blue is sky the"));
    tests.add(new ReverseWordStringTestCase("  hello world  ", "world hello"));
    tests.add(new ReverseWordStringTestCase("a good   example", "example good a"));

    for (ReverseWordStringTestCase test : tests) {
      assertEquals(test.answer, ReverseWordString.reverseWords(test.s));
    }
  }
}
