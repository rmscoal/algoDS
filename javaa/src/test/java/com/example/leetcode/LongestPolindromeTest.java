package com.example.leetcode;

import org.junit.jupiter.api.Test;
import static org.junit.jupiter.api.Assertions.assertEquals;

import java.util.HashMap;

public class LongestPolindromeTest {
    public int solution(String s) {
        int result = 0;
        HashMap<Character, Integer> map = new HashMap<>();

        for (int i = 0; i < s.length(); i++) {
            map.put(s.charAt(i), map.getOrDefault(s.charAt(i), 0) + 1);
        }

        for (Character c : map.keySet()) {
            int count = map.get(c);
            if (count > 1) {
                int pairs = (count / 2) * 2; // get the highest even number
                result += pairs;
                map.put(c, count - pairs);
            }
        }

        // Now, we see if there are ones left in the map
        // and add to our result.
        for (Character c : map.keySet()) {
            int count = map.get(c);
            if (count == 1) {
                result++;
                break;
            }
        }

        return result;
    }

    @Test
    public void test() {
        assertEquals(7, solution("abccccdd"));
        assertEquals(1, solution("a"));
        assertEquals(1, solution("Aa"));
        assertEquals(3, solution("Aaaa"));
    }
}
