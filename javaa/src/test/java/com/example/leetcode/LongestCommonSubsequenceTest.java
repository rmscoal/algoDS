package com.example.leetcode;

import static org.junit.jupiter.api.Assertions.assertEquals;
import org.junit.jupiter.api.Test;

public class LongestCommonSubsequenceTest {
    public int solution(String a, String b) {
        // Cases
        // A: abcde
        // B: ace
        // -> 3

        // A: abccdame
        // B: acdfg
        // -> 3

        // A: accde
        // B: accc
        // -> 3

        // Memoization
        // x | a | c | c | c |
        // -----------------------
        // a | 1 | 1 | 1 | 1 |
        // c | 1 | 2 | 2 | 2 |
        // c | 1 | 2 | 3 | 3 |
        // d | 1 | 2 | 3 | 3 |
        // e | 1 | 2 | 3 | 3 |

        int[][] matrix = new int[a.length() + 1][b.length() + 1];

        for (int i = 1; i <= a.length(); i++) {
            for (int j = 1; j <= b.length(); j++) {
                if (a.charAt(i - 1) == b.charAt(j - 1)) {
                    matrix[i][j] = matrix[i - 1][j - 1] + 1;
                } else {
                    matrix[i][j] = Math.max(matrix[i - 1][j], matrix[i][j - 1]);
                }
            }
        }

        return matrix[a.length()][b.length()];
    }

    public void print(int[][] matrix) {
        for (int i = 0; i < matrix.length; i++) {
            for (int j = 0; j < matrix[i].length; j++) {
                System.out.print(matrix[i][j] + " ");
            }
            System.out.println();
        }
    }

    @Test
    public void test() {
        assertEquals(3, solution("abccdame", "acdfg"));
        assertEquals(3, solution("accc", "accde"));
    }
}
