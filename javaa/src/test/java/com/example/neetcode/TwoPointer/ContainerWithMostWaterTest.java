package com.example.neetcode.TwoPointer;

import static org.junit.jupiter.api.Assertions.assertEquals;
import org.junit.jupiter.api.Test;

import java.util.ArrayList;

public class ContainerWithMostWaterTest {
    static class TestCase {
        int[] heights;
        int expected;

        public TestCase(int[] heights, int expected) {
            this.heights = heights;
            this.expected = expected;
        }
    }

    public int maxArea(int[] height) {
        int maxArea = 0;

        // The two pointers
        int left = 0;
        int right = height.length - 1;

        while (left < right) {
            maxArea = Math.max(maxArea, calculateArea(height, left, right));
            if (height[left] < height[right]) {
                left++;
            } else {
                right--;
            }
        }

        return maxArea;
    }

    private int calculateArea(int[] height, int left, int right) {
        return (right - left) * Math.min(height[left], height[right]);
    }

    @Test
    void testMaxArea() {
        ArrayList<TestCase> testCases = new ArrayList<>();
        testCases.add(new TestCase(new int[]{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49));
        testCases.add(new TestCase(new int[]{1,1}, 1));

        for (TestCase testCase : testCases) {
            assertEquals(testCase.expected, maxArea(testCase.heights));
        }
    }
}
