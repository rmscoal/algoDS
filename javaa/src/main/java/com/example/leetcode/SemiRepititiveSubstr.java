package com.example.leetcode;

import java.util.ArrayList;

public class SemiRepititiveSubstr {
  /**
   * https://leetcode.com/problems/find-the-longest-semi-repetitive-substring/description/
   * Here I think we can use sliding window
   * 
   * @param s
   * @return The longest semi-repititive substring
   */
  public static int longestSemiRepititiveSubstring(String s) {
    // Result variable
    int longest = 1;
    // Sliding window pointer
    int left = 0;
    int right = 1;

    while (right < s.length()) {
      boolean alrOnePair = false;

      for (int i = left; i < right; i++) {
        if (s.charAt(i) == s.charAt(i + 1)) {
          if (alrOnePair) {
            longest = Math.max(longest, i - left + 1);
            left++;
            break;
          } else {
            alrOnePair = true;
          }
        }
        longest = Math.max(longest, i - left + 2);
      }

      right++;
    }

    return longest;
  }

  public static int longestSemiRepititiveSubstringV2(String s) {
    int longest = 0;
    ArrayList<Integer> lengths = new ArrayList<Integer>();
    lengths.add(1);

    for (int i = 1; i < s.length(); i++) {
      if (s.charAt(i) == s.charAt(i - 1)) {
        lengths.add(1);
      } else {
        lengths.set(lengths.size() - 1, lengths.get(lengths.size() - 1) + 1);
      }
    }

    if (lengths.size() == 1) {
      return lengths.getLast();
    }

    for (int i = 1; i < lengths.size(); i++) {
      longest = Math.max(longest, lengths.get(i) + lengths.get(i - 1));
    }

    return longest;
  }
}
