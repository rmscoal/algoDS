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

  /**
   * This is my best solution so far for:
   * https://leetcode.com/problems/find-the-longest-semi-repetitive-substring/description/
   * 
   * @param s - a string of length 1 <= x <= 50 with character ranging from '0' -
   *          '9'
   * @return The longest semi-repititive substring
   */
  public static int longestSemiRepititiveSubstringFastest(String s) {
    if (s.length() == 1) {
      return 1;
    }

    // Result variable
    int longest = 0;
    // Keeps track of how many pairs are there in the window
    int pair = 0;
    // Pointers of our window
    int left = 0;
    int right = 1;

    while (right < s.length()) {
      if (s.charAt(right) == s.charAt(right - 1)) {
        pair++;
      }

      while (pair >= 2) {
        if (s.charAt(left) == s.charAt(left + 1)) {
          pair--;
        }
        left++;
      }

      longest = Math.max(longest, right - left + 1);
      right++;
    }

    return longest;
  }
}
