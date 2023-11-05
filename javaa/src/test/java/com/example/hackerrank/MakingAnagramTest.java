package com.example.hackerrank;

import static org.junit.jupiter.api.Assertions.assertEquals;

import java.util.Vector;

import org.junit.jupiter.api.Test;

class MakingAnagramTestCase {
  public String a;
  public String b;
  public int answer;

  MakingAnagramTestCase(String a, String b, int answer) {
    this.a = a;
    this.b = b;
    this.answer = answer;
  }
}

public class MakingAnagramTest {
  @Test
  void makingAnagramTest() {
    Vector<MakingAnagramTestCase> tests = new Vector<MakingAnagramTestCase>();
    tests.add(new MakingAnagramTestCase("abc", "cde", 4));

    for (MakingAnagramTestCase test : tests) {
      assertEquals(test.answer, MakingAnagram.makeAnagrams(test.a, test.b));
    }
  }
}
