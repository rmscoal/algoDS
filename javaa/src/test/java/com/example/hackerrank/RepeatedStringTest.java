package com.example.hackerrank;

import static org.junit.jupiter.api.Assertions.assertEquals;

import java.util.ArrayList;
import java.util.concurrent.TimeUnit;

import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.Timeout;

class RepeatedStringTestCase {
  public String s;
  public long n;
  public long answer;

  public RepeatedStringTestCase(String s, long n, long answer) {
    this.s = s;
    this.n = n;
    this.answer = answer;
  }
}

public class RepeatedStringTest {
  @Test
  @Timeout(value = 1, unit = TimeUnit.MILLISECONDS)
  void repeatedStringTest() {
    ArrayList<RepeatedStringTestCase> tests = new ArrayList<RepeatedStringTestCase>();
    tests.add(new RepeatedStringTestCase("aba", 10, 7));
    tests.add(new RepeatedStringTestCase("a", 1000000000000l, 1000000000000l));
    tests.add(new RepeatedStringTestCase("abbbcsd", 19, 3));
    tests.add(new RepeatedStringTestCase("abbbcaa", 19, 7));

    tests.forEach((test) -> {
      long result = RepeatedString.repeatedString(test.s, test.n);
      assertEquals(test.answer, result);
    });
  }
}
