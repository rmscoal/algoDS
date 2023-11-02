package com.example.leetcode;

public class ReverseWordString {
  public static String reverseWords(String s) {
    String[] strArr = s.split("\\s+");
    String result = new String();

    for (int i = strArr.length - 1; i > -1; i--) {
      if (!strArr[i].equals("\\s+")) {
        result += strArr[i].trim() + " ";
      }
    }

    return result.trim();
  }
}
