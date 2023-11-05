package com.example.hackerrank;

import java.util.*;

public class SherlockValidString {
  public static String isValid(String s) {
    if (s.length() <= 1) {
      return "YES";
    }

    // This hash map holds the occurences of a character in the string
    HashMap<Character, Integer> characterOccurance = new HashMap<Character, Integer>();

    for (int i = 0; i < s.length(); i++) {
      Character c = s.charAt(i);

      if (characterOccurance.containsKey(c)) {
        characterOccurance.put(c, characterOccurance.get(c) + 1);
      } else {
        characterOccurance.put(c, 1);
      }
    }

    // System.out.println("Characted Occurence");
    // for (Map.Entry<Character, Integer> entry : characterOccurance.entrySet()) {
    // System.out.println("Key " + entry.getKey() + " - Value " + entry.getValue());
    // }

    // This hash map holds the frequency of the occurences that is
    // how many times there are 3 character occurance and so on.
    HashMap<Integer, Integer> frequencyMap = new HashMap<Integer, Integer>();

    for (int value : characterOccurance.values()) {
      if (frequencyMap.containsKey(value)) {
        frequencyMap.put(value, frequencyMap.get(value) + 1);
      } else {
        frequencyMap.put(value, 1);
      }
    }

    // Now check how many frequency are there. If there are 2 frequency
    // we need to check whether the key 1 exists only once.
    if (frequencyMap.keySet().size() == 2) {
      for (Map.Entry<Integer, Integer> entry : frequencyMap.entrySet()) {
        if (entry.getKey() == 1 && entry.getValue() == 1) {
          return "YES";
        }
      }
    }

    // for (Map.Entry<Integer, Integer> entry : frequencyMap.entrySet()) {
    // System.out.println("Key " + entry.getKey() + " - Value " + entry.getValue());
    // }

    // Now we are going to sort by the most frequent character
    // occurance.
    List<Integer> freqList = new LinkedList<Integer>(frequencyMap.values());
    freqList.sort(Collections.reverseOrder());

    // for (Integer el : freqList) {
    // System.out.println(el);
    // }

    // Next, we are going to get the most frequent occurence
    // as the base value such that we can compare how many
    // characters should be removed.
    int mostFrequent = freqList.get(0);
    int base = 0;
    int diffTotal = 0;

    for (Map.Entry<Integer, Integer> entry : frequencyMap.entrySet()) {
      if (mostFrequent == entry.getValue()) {
        base = entry.getKey();
        break;
      }
    }

    // System.out.println("Base " + base);

    for (Map.Entry<Character, Integer> entry : characterOccurance.entrySet()) {
      diffTotal += Math.abs(base - entry.getValue());
    }

    if (diffTotal > 1) {
      return "NO";
    }

    return "YES";
  }
}
