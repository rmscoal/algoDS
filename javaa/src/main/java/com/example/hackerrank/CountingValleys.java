package com.example.hackerrank;

public class CountingValleys {
  public static int countingValleys(int steps, String path) {
    /*
     * Given [U, D, D, D, U, D, U, U]
     * 
     * Note: It always start from sea level and end at sea level.
     * 
     * We are going to count how many valleys are there.
     * 
     * Variable "height" -> this stores the current height...
     * 
     * Then we are going to loop through the array and count each height...
     * Every time the height reaches zero from a negative value we are going
     * to count that as a valley
     */

    int valleyCount = 0;
    int height = 0;

    for (int i = 0; i < steps; i++) {
      if (path.charAt(i) == 'U') {
        if (height < 0 && height + 1 == 0) {
          valleyCount++;
        }
        height++;
      } else if (path.charAt(i) == 'D') {
        height--;
      }
    }

    return valleyCount;
  }
}
