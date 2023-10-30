package com.example.hackerrank;

import java.util.List;

public class JumpOnClouds {
  public static int jumpingOnClouds(List<Integer> c) {
    // The player must skip to a 0-value by 1 or 2 indices.
    // The player must avoid the 1-value.
    // Return the minimum steps it takes to jump to the last cloud
    //
    // Note:
    // First and last cloud will always be 0-value.
    //
    // Example -> [0, 0, 1, 0, 0, 1, 0]
    // Example -> [0, 1, 0, 0, 0, 1, 0]
    // Working solution:
    // We start at index i = 0.
    // i + 2, if c[i + 2] == 1, i++, if i + 2 < c.length - 1;

    int i = 0;
    int jumps = 0;

    while (i != c.size() - 1) {
      try {
        if (c.get(i + 2) != 1) {
          i += 2;
        } else {
          i += 1;
        }
        jumps++;
      } catch (Exception e) {
        jumps++;
        break;
      }
    }

    return jumps;
  }
}
