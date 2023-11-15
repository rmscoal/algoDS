package com.example.random;

import static org.junit.jupiter.api.Assertions.assertEquals;

import org.junit.jupiter.api.Test;

public class DisjointIntervalTest {
  @Test
  void disjointIntervalTest() {
    assertEquals(3, DisjoinInterval.disjointInterval(new int[][] { { 1, 4 }, { 2, 3 }, { 4, 6 }, { 8, 9 } }));
  }
}
