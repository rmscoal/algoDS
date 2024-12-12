package com.example.leetcode;

import static org.junit.jupiter.api.Assertions.assertEquals;
import org.junit.jupiter.api.Test;

import java.util.ArrayList;

public class AtoiTest {
    public int solution(String s) {
        s = s.trim();

        long result = 0;
        int multiplier = 1;

        for (int i = 0; i < s.length(); i++) {
            if ((s.charAt(i) == '-' || s.charAt(i) == '+') && i == 0) {
                multiplier = s.charAt(i) == '-' ? -1 : 1;
                continue;
            }

            // If the character is not a digit, we assume it is broken.
            if (s.charAt(i) < '0' || s.charAt(i) > '9') {
                break;
            }

            result = (result * 10) + s.charAt(i) - '0';
            if (multiplier * result > Integer.MAX_VALUE) {
                return Integer.MAX_VALUE;
            } else if (result * multiplier < Integer.MIN_VALUE) {
                return Integer.MIN_VALUE;
            }
        }

        return (int) result * multiplier;
    }

    static class TestCase {
        String input;
        int want;

        public TestCase(String input, int want) {
            this.input = input;
            this.want = want;
        }
    }

    @Test
    public void test() {
        ArrayList<TestCase> testCases = new ArrayList<>();
        testCases.add(new TestCase("1", 1));
        testCases.add(new TestCase("2", 2));
        testCases.add(new TestCase("3", 3));
        testCases.add(new TestCase("4", 4));
        testCases.add(new TestCase("5", 5));
        testCases.add(new TestCase("42", 42));
        testCases.add(new TestCase("-42", -42));
        testCases.add(new TestCase("-042", -42));
        testCases.add(new TestCase("1337c0d3", 1337));
        testCases.add(new TestCase("0-1", 0));
        testCases.add(new TestCase("words and 987", 0));
        testCases.add(new TestCase("-91283472332", -2147483648));

        for (TestCase testCase : testCases) {
            assertEquals(testCase.want, solution(testCase.input));
        }
    }
}
