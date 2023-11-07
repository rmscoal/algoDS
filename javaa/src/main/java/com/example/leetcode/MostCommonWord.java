package com.example.leetcode;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.Map;

public class MostCommonWord {
  public static String mostCommonWord(String paragraph, String[] banned) {
    ArrayList<String> collections = new ArrayList<String>();
    boolean wasStr = false;
    int prevStr = 0;

    // Lower case our paragraph
    paragraph = paragraph.toLowerCase();

    for (int i = 0; i < paragraph.length(); i++) {
      char currentChar = paragraph.charAt(i);

      if (Character.isLetter(currentChar)) {
        if (!wasStr) {
          prevStr = i;
          wasStr = true;
        }
      } else if (wasStr) {
        collections.add(paragraph.substring(prevStr, i));
        wasStr = false;
      }
    }

    // Add the last word if the paragraph ends with a word
    if (wasStr) {
      collections.add(paragraph.substring(prevStr));
    }

    Map<String, Character> bannedWords = new HashMap<String, Character>();
    Map<String, Integer> wordCount = new HashMap<String, Integer>();

    if (banned != null) {
      for (String str : banned) {
        bannedWords.put(str, ' ');
      }
    }

    for (String str : collections) {
      if (!bannedWords.isEmpty()) {
        if (bannedWords.get(str) != null) {
          continue;
        }
      }
      wordCount.put(str, wordCount.get(str) == null ? 1 : wordCount.get(str) + 1);
    }

    int maxCount = 0;
    String mostWord = "";

    for (Map.Entry<String, Integer> entry : wordCount.entrySet()) {
      if (maxCount < entry.getValue()) {
        maxCount = entry.getValue();
        mostWord = entry.getKey();
      }
    }

    return mostWord;
  }
}
