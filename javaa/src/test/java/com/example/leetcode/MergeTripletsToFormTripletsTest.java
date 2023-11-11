package com.example.leetcode;

import static org.junit.jupiter.api.Assertions.assertEquals;

import java.util.ArrayList;

import org.junit.jupiter.api.Test;

class MergeTripletsToFormTripletsTestCase {
  public int[][] triplets;
  public int[] targets;
  public boolean answer;

  MergeTripletsToFormTripletsTestCase(int[][] triplets, int[] targets, boolean answer) {
    this.triplets = triplets;
    this.targets = targets;
    this.answer = answer;
  }
}

public class MergeTripletsToFormTripletsTest {
  @Test
  void mergeTripletsTest() {
    ArrayList<MergeTripletsToFormTripletsTestCase> tests = new ArrayList<MergeTripletsToFormTripletsTestCase>();
    tests.add(new MergeTripletsToFormTripletsTestCase(
        new int[][] { { 9, 10, 11 }, { 9, 7, 5 }, { 1, 2, 5 }, { 5, 2, 3 } }, new int[] { 5, 2, 5 }, true));
    tests.add(new MergeTripletsToFormTripletsTestCase(
        new int[][] { { 2, 5, 3 }, { 2, 3, 4 }, { 1, 2, 5 }, { 5, 2, 3 } }, new int[] { 5, 5, 5 }, true));
    tests.add(new MergeTripletsToFormTripletsTestCase(
        new int[][] { { 2, 5, 3 }, { 1, 8, 4 }, { 1, 7, 5 } }, new int[] { 2, 7, 5 },
        true));
    tests.add(new MergeTripletsToFormTripletsTestCase(
        new int[][] { { 3, 4, 5 }, { 4, 5, 6 } }, new int[] { 3, 2, 5 }, false));
    tests.add(new MergeTripletsToFormTripletsTestCase(
        new int[][] { { 1, 1, 1 }, { 2, 2, 2 }, { 3, 2, 3 } }, new int[] { 2, 2, 2 }, true));
    tests.add(new MergeTripletsToFormTripletsTestCase(
        new int[][] { { 1, 1, 1 }, { 2, 2, 2 }, { 3, 2, 3 } }, new int[] { 3, 2, 2 }, false));

    for (MergeTripletsToFormTripletsTestCase test : tests) {
      assertEquals(test.answer, MergeTripletsToFormTriplets.mergeTriplets(test.triplets, test.targets));
    }
  }

  @Test
  void mergeTripletsTestFaster() {
    ArrayList<MergeTripletsToFormTripletsTestCase> tests = new ArrayList<MergeTripletsToFormTripletsTestCase>();
    tests.add(new MergeTripletsToFormTripletsTestCase(
        new int[][] { { 9, 10, 11 }, { 9, 7, 5 }, { 1, 2, 5 }, { 5, 2, 3 } }, new int[] { 5, 2, 5 }, true));
    tests.add(new MergeTripletsToFormTripletsTestCase(
        new int[][] { { 2, 5, 3 }, { 2, 3, 4 }, { 1, 2, 5 }, { 5, 2, 3 } }, new int[] { 5, 5, 5 }, true));
    tests.add(new MergeTripletsToFormTripletsTestCase(
        new int[][] { { 2, 5, 3 }, { 1, 8, 4 }, { 1, 7, 5 } }, new int[] { 2, 7, 5 },
        true));
    tests.add(new MergeTripletsToFormTripletsTestCase(
        new int[][] { { 3, 4, 5 }, { 4, 5, 6 } }, new int[] { 3, 2, 5 }, false));
    tests.add(new MergeTripletsToFormTripletsTestCase(
        new int[][] { { 1, 1, 1 }, { 2, 2, 2 }, { 3, 2, 3 } }, new int[] { 2, 2, 2 }, true));
    tests.add(new MergeTripletsToFormTripletsTestCase(
        new int[][] { { 1, 1, 1 }, { 2, 2, 2 }, { 3, 2, 3 } }, new int[] { 3, 2, 2 }, false));

    for (MergeTripletsToFormTripletsTestCase test : tests) {
      assertEquals(test.answer, MergeTripletsToFormTriplets.mergeTripletsFaster(test.triplets, test.targets));
    }
  }
}
