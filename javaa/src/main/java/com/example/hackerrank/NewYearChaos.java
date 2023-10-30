package com.example.hackerrank;

import java.util.List;

public class NewYearChaos {
  public static int minimumBribes(List<Integer> q) {
    /*
     * Working Solution Method:
     * Input: [2, 1, 5, 3, 4]
     * 
     * [1, 2, 5, 3, 4] -> 1
     * [1, 2, 3, 5, 4] -> 2
     * [1, 2, 3, 4, 5] -> 3
     * 
     * We want that the result is sorted...
     * 
     * Bubble Sort -> count how many swaps have taken place... For each element
     * that is currently doing a swap, if it swaps more than twice, then we can
     * print Too Chaotic...
     * 
     * [1,2,6,3,4,5]
     */

    int swaps = 0;

    // Bubble Sort
    for (int i = 0; i < q.size(); i++) {
      boolean swapped = false;
      int currentSwap = 0;

      for (int j = 0; j < q.size() - i - 1; j++) {
        if (q.get(j) > q.get(j + 1)) {
          // Swap elements between j and j+1;
          int tmp = q.get(j);
          q.set(j, q.get(j + 1));
          q.set(j + 1, tmp);
          swaps++;
          currentSwap++;
          swapped = true;
        } else {
          currentSwap = 0;
        }

        if (currentSwap > 2) {
          System.out.println("Too chaotic");
          return -1;
        }
      }

      if (!swapped) {
        break;
      }
    }

    System.out.println(swaps);
    return swaps;
  }

  public static int minimumBribesFaster(List<Integer> q) {
    /*
     * Input: [1 2 5 3 7 8 6 4]
     * 
     * [1,2,3,4,5,6,7,8]
     * [1,2,3,5,4,6,7,8]
     * [1,2,3,5,6,4,7,8]
     * [1,2,3,5,7,6,4,8]
     * [1,2,3,5,7,8,6,4]
     */
    int swaps = 0;
    int n = q.size();
    int person, // Not yet initialized
        position, // Not yet initialized
        // This min variable will save the lowest original position found in the queue
        min = n; // 8

    for (int i = n - 1; i > -1; i--) {
      person = q.get(i);
      position = i + 1;

      if (person > position) { // this means that the person was moved forward or bribed forward
        swaps += person - position; // calculate how many positions did the person bribes...

        if (person - position > 2) {
          System.out.println("Too chaotic");
          return -1;
        }
      } else {
        if (person < min) {
          min = person;
        } else if (person != min) {
          swaps++; // i still dont know why this
        }

      }
    }

    System.out.println(swaps);
    return swaps;
  }
}
