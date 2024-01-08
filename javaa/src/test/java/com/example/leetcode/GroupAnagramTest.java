package com.example.leetcode;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

import org.junit.jupiter.api.Test;

class GroupAnagramTest {
  public static List<List<String>> groupAnagrams(String[] strs) {
    Map<String, List<String>> hash = new HashMap<String, List<String>>();

    for (String str : strs) {
      // We want to sort the string
      char[] chars = str.toCharArray();
      Arrays.sort(chars);
      String sorted = new String(chars);

      // If the sorted string is not in the hash
      // we make a new empty list.
      hash.putIfAbsent(sorted, new ArrayList<>());
      hash.get(sorted).add(str);
    }

    return new ArrayList<>(hash.values());
  }

  @Test
  void testGroupAnagrams() {
     GroupAnagramTest.groupAnagrams(new String[] {"eat","tea","tan","ate","nat","bat"});
  }
}
