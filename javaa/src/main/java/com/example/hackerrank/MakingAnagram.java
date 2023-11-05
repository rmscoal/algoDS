package com.example.hackerrank;

import java.util.HashMap;

public class MakingAnagram {

  public static int makeAnagrams(String a, String b) {
    // Result variable, the count of deletion to be made.
    int deletionCount = 0;

    // Make a hasmap for 'a' and count the occurence of the letters
    HashMap<Character, Integer> mapA = new HashMap<Character, Integer>();

    for (int i = 0; i < a.length(); i++) {
      Character key = a.charAt(i);
      if (mapA.containsKey(key)) {
        mapA.put(a.charAt(i), mapA.get(key) + 1);
      } else {
        mapA.put(key, 1);
      }
    }

    // Make a hashmap for 'b' and count the occurence of the letters.
    // Here we are adding extra check that is checking if our letter
    // is also present in mapA. If not then it needs to be count as
    // deleted.
    HashMap<Character, Integer> mapB = new HashMap<Character, Integer>();

    for (int i = 0; i < b.length(); i++) {
      Character key = b.charAt(i);
      if (!mapA.containsKey(key)) {
        deletionCount++;
      } else if (mapB.containsKey(key)) {
        mapB.put(key, mapB.get(key) + 1);
      } else {
        mapB.put(key, 1);
      }
    }

    // Now we iterate from map A to find the minimum of both count
    // for the given characters.
    for (Character key : mapA.keySet()) {
      int countFromA = mapA.get(key);
      int countFromB = mapB.get(key) == null ? 0 : mapB.get(key);
      int minimum = countFromA > countFromB ? countFromB : countFromA;
      deletionCount += (countFromA + countFromB - (2 * minimum));
    }

    return deletionCount;
  }
}
