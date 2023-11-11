package com.example.leetcode;

import java.util.ArrayList;
import java.util.List;

public class MergeTripletsToFormTriplets {
  /**
   * Firstly, collect all triplet where it
   * is lesser than or equal to the target
   * and store the triplet index in a new
   * array.Next, we want to find the max
   * for each of the triplet element. If
   * the max of each is equal to our target
   * then we return true else we return
   * false.
   * 
   * @param triplets
   * @param target
   * @return
   */
  public static boolean mergeTriplets(int[][] triplets, int[] target) {
    List<Integer> indexes = new ArrayList<Integer>();

    for (int i = 0; i < triplets.length; i++) {
      if (triplets[i][0] <= target[0] && triplets[i][1] <= target[1] && triplets[i][2] <= target[2]) {
        indexes.add(i);
      }
    }

    int a = 0, b = 0, c = 0;

    for (int idx : indexes) {
      a = Math.max(a, triplets[idx][0]);
      b = Math.max(b, triplets[idx][1]);
      c = Math.max(c, triplets[idx][2]);

      if (a == target[0] && b == target[1] && c == target[2]) {
        return true;
      }
    }

    return false;
  }

  /**
   * The steps here is to find for each triplet element. Such that
   * x is found and y is found and z is found.
   * 
   * @param triplets
   * @param target
   * @return
   */
  public static boolean mergeTripletsFaster(int[][] triplets, int[] target) {
    boolean xFound = false;
    for (int i = 0; i < triplets.length; i++) {
      if (triplets[i][0] == target[0] && triplets[i][1] <= target[1] && triplets[i][2] <= target[2]) {
        xFound = true;
        break;
      }
    }

    boolean yFound = false;
    for (int i = 0; i < triplets.length; i++) {
      if (triplets[i][1] == target[1] && triplets[i][0] <= target[0] && triplets[i][2] <= target[2]) {
        yFound = true;
        break;
      }
    }

    boolean zFound = false;
    for (int i = 0; i < triplets.length; i++) {
      if (triplets[i][2] == target[2] && triplets[i][0] <= target[0] && triplets[i][1] <= target[1]) {
        zFound = true;
        break;
      }
    }

    return xFound && yFound && zFound;
  }
}
