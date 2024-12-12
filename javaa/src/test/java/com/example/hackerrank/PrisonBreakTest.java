package com.example.hackerrank;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertEquals;
import java.util.ArrayList;

public class PrisonBreakTest {
    public int solution(int n, int m, int[] h, int[] v) {
        int maxVerticalGap = 1;
        int maxHorizontalGap = 1;
        int iterationGap = 1;

        for (int i = 0; i < h.length; i++) {
            if (i == 0) {
                iterationGap++;
            } else if (h[i] == h[i-1] + 1) {
                iterationGap++;
            } else {
                iterationGap = 1;
            }
            maxVerticalGap = Math.max(maxVerticalGap, iterationGap);
        }

        iterationGap = 1;
        for (int i = 0; i < v.length; i++) {
            if (i == 0) {
                iterationGap++;
            } else if (v[i] == v[i-1] + 1) {
                iterationGap++;
            } else {
                iterationGap = 1;
            }
            maxHorizontalGap = Math.max(maxHorizontalGap, iterationGap);
        }

        return maxVerticalGap * maxHorizontalGap;
    }

    static class TestCase {
        int n;
        int m;
        int[] h;
        int[] v;
        int want;

        public TestCase(int n, int m, int[] h, int[] v, int want) {
           this.n = n;
           this.m = m;
           this.h = h;
           this.v = v;
           this.want = want;
        }
    }

    @Test
    public void test() {
        ArrayList<TestCase> testCases = new ArrayList<>();
        testCases.add(new PrisonBreakTest.TestCase(6,6,new int[] {4},
                new int[] {2}, 4));
        testCases.add(new PrisonBreakTest.TestCase(1,1,new int[] {},
                new int[] {}, 1));
        testCases.add(new PrisonBreakTest.TestCase(3,3,new int[] {},
                new int[] {}, 1));
        testCases.add(new PrisonBreakTest.TestCase(3,3,new int[] {1,2,3},
                new int[] {1,2,3}, 16));
        testCases.add(new PrisonBreakTest.TestCase(5,5,new int[] {2,4},
                new int[] {1,3,5}, 4));
        testCases.add(new PrisonBreakTest.TestCase(4,4,new int[] {1,4},
                new int[] {1,4}, 4));
        testCases.add(new PrisonBreakTest.TestCase(10,10,new int[] {2,4,6,8,10},
                new int[] {1,3,5,7,9}, 4));

        for (TestCase testCase : testCases) {
            assertEquals(testCase.want, solution(testCase.n, testCase.m, testCase.h, testCase.v));
        }
    }
}
