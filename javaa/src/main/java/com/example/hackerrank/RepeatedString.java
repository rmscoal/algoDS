package com.example.hackerrank;

public class RepeatedString {
  public static long repeatedString(String s, long n) {
    long additional = 0;
    long count = 0;
    long multiplier = n / s.length();

    for (int i = 0; i < s.length(); i++) {
      if (s.charAt(i) == 'a') {
        count++;
      }
    }

    count *= multiplier;

    for (int i = 0; i < n % s.length(); i++) {
      if (s.charAt(i) == 'a') {
        additional++;
      }
    }

    return count + additional;
  }
}
