package com.example.leetcode;

import static org.junit.jupiter.api.Assertions.assertEquals;

import java.util.ArrayList;

import org.junit.jupiter.api.Test;

class SemiRepititiveSubstrTestCase {
  public String s;
  public int asnwer;

  SemiRepititiveSubstrTestCase(String s, int answer) {
    this.s = s;
    this.asnwer = answer;
  }
}

public class SemiRepititiveSubstrTest {
  @Test
  void semiRepititiveSubstrTest() {
    ArrayList<SemiRepititiveSubstrTestCase> tests = new ArrayList<SemiRepititiveSubstrTestCase>();
    tests.add(new SemiRepititiveSubstrTestCase("2233", 3));
    tests.add(new SemiRepititiveSubstrTestCase("1111111", 2));
    tests.add(new SemiRepititiveSubstrTestCase("5494", 4));
    tests.add(new SemiRepititiveSubstrTestCase("52233", 4));
    tests.add(new SemiRepititiveSubstrTestCase("12", 2));
    tests.add(new SemiRepititiveSubstrTestCase("14", 2));
    tests.add(new SemiRepititiveSubstrTestCase("0123", 4));
    tests.add(new SemiRepititiveSubstrTestCase("54944", 5));
    tests.add(new SemiRepititiveSubstrTestCase("00101022", 7));
    tests.add(new SemiRepititiveSubstrTestCase("0001", 3));
    tests.add(new SemiRepititiveSubstrTestCase("1", 1));

    for (SemiRepititiveSubstrTestCase test : tests) {
      assertEquals(test.asnwer, SemiRepititiveSubstr.longestSemiRepititiveSubstring(test.s));
    }
  }

  @Test
  void semiRepititiveSubstrV2Test() {
    ArrayList<SemiRepititiveSubstrTestCase> tests = new ArrayList<SemiRepititiveSubstrTestCase>();
    tests.add(new SemiRepititiveSubstrTestCase("2233", 3));
    tests.add(new SemiRepititiveSubstrTestCase("1111111", 2));
    tests.add(new SemiRepititiveSubstrTestCase("5494", 4));
    tests.add(new SemiRepititiveSubstrTestCase("52233", 4));
    tests.add(new SemiRepititiveSubstrTestCase("12", 2));
    tests.add(new SemiRepititiveSubstrTestCase("14", 2));
    tests.add(new SemiRepititiveSubstrTestCase("0123", 4));
    tests.add(new SemiRepititiveSubstrTestCase("54944", 5));
    tests.add(new SemiRepititiveSubstrTestCase("00101022", 7));
    tests.add(new SemiRepititiveSubstrTestCase("0001", 3));
    tests.add(new SemiRepititiveSubstrTestCase("1", 1));

    for (SemiRepititiveSubstrTestCase test : tests) {
      assertEquals(test.asnwer, SemiRepititiveSubstr.longestSemiRepititiveSubstringV2(test.s));
    }
  }
}
