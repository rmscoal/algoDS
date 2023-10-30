package com.example.hackerrank;

import java.util.HashMap;
import java.util.List;

import org.junit.jupiter.api.Tag;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertEquals;

import java.util.Arrays;

class SalesByMatchTestCase {
  int n;
  List<Integer> arr;
  int answer;

  public SalesByMatchTestCase(int n, List<Integer> arr, int answer) {
    this.n = n;
    this.arr = arr;
    this.answer = answer;
  }
}

public class SalesByMatchTest {
  @Test
  @Tag("fast")
  void sockMerchantTest() {
    HashMap<String, SalesByMatchTestCase> tests = new HashMap<String, SalesByMatchTestCase>();
    tests.put("test_1", new SalesByMatchTestCase(7, Arrays.asList(1, 2, 1, 2, 1, 3, 2), 2));
    tests.put("test_2", new SalesByMatchTestCase(9, Arrays.asList(10, 20, 20, 10, 10, 30, 50, 10, 20), 3));

    tests.forEach((name, test) -> {
      System.out.println("Test Sales By Match " + name);
      int result = SalesByMatch.sockMerchant(test.n, test.arr);
      assertEquals(test.answer, result);
    });
  }

}
