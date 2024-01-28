package com.example.leetcode.binary_search;

import static org.junit.jupiter.api.Assertions.assertEquals;

import java.util.ArrayList;
import java.util.List;

import org.junit.jupiter.api.Test;

class Search2DMatrixTest {
  public boolean searchMatrix_OkeBuddy(int[][] matrix, int target) {
    for (int i = 0; i < matrix.length; i++) {
      for (int j = 0; j < matrix[0].length; j++) {
        if (matrix[i][j] == target) {
          return true;
        }
      }
    }
    return false;
  }

  public boolean searchMatrix(int[][] matrix, int target) {
    // First, binary search the first column of each row...
    // We should be looking at matrix[i] where matrix[i][0] is lesser
    // and matrix[i+1][0] is greater than the target.
    //
    // Once we have that row, we can now search amongst all
    // elements in that row.
    //
    // To make it easier, we are going to check that the
    // target is indeed inside the matrix bound, that means
    // target <= matrix[last][last] and target >= matrix[first][first]

    int m = matrix.length;
    int n = matrix[0].length;

    if (target < matrix[0][0] || target > matrix[m - 1][n - 1]) {
      return false;
    }

    int left = 0;
    int right = m - 1;
    int pickedRow = 0;

    while (left < right) {
      int mid = (right + left) / 2;

      if (target >= matrix[mid][0]) {
        if (target < matrix[mid + 1][0]) {
          pickedRow = mid;
          break;
        } else {
          left = mid + 1;
          pickedRow = left;
        }
      } else if (target <= matrix[mid][0]) {
        if (target > matrix[mid - 1][0]) {
          pickedRow = mid - 1;
          break;
        } else {
          right = mid - 1;
          pickedRow = right;
        }
      }
    }

    // After we found our row, we do binary search on that row
    left = 0;
    right = n - 1;

    while (left <= right) {
      int mid = (right + left) / 2;

      if (target == matrix[pickedRow][mid]) {
        return true;
      } else if (target > matrix[pickedRow][mid]) {
        left = mid + 1;
      } else {
        right = mid - 1;
      }
    }

    return false;
  }

  class TestCase {
    int[][] matrix;
    int target;
    boolean result;

    public TestCase(int[][] matrix, int target, boolean result) {
      this.matrix = matrix;
      this.target = target;
      this.result = result;
    }
  }

  @Test
  public void testSearch2DMatrixTest() {
    List<TestCase> testcases = new ArrayList<>();

    testcases.add(new TestCase(new int[][] { { 1, 3, 5, 7 }, { 10, 11, 16, 20 },
        { 23, 30, 34, 60 } }, 5, true));
    testcases.add(new TestCase(new int[][] { { 1, 3, 5, 7 }, { 10, 11, 16, 20 },
        { 23, 30, 34, 60 } }, 13, false));
    testcases.add(new TestCase(new int[][] { { 1, 3, 5, 7 }, { 10, 11, 16, 20 },
        { 23, 30, 34, 60 } }, 40, false));
    testcases.add(new TestCase(new int[][] { { 1, 3, 5, 7 }, { 10, 11, 16, 20 },
        { 23, 30, 34, 60 } }, 20, true));
    testcases.add(new TestCase(new int[][] { { 1, 3, 5, 7 }, { 10, 11, 16, 20 },
        { 23, 30, 34, 60 } }, 21, false));
    testcases.add(new TestCase(new int[][] { { 1, 3, 5, 7 }, { 10, 11, 16, 20 },
        { 23, 30, 34, 60 } }, 30, true));

    for (TestCase tc : testcases) {
      assertEquals(tc.result, new Search2DMatrixTest().searchMatrix_OkeBuddy(tc.matrix,
          tc.target));
    }
  }
}
