package binary_search;


class Search2DMatrixTest {
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

    if (matrix[0][0] > target || target > matrix[m-1][n-1]) {
      return false;
    }

    int left = 0;
    int right = m-1;
    int pickedRow = 0;

    while (left < right) {
      int mid = (right + left) / 2;

      if (target <= matrix[mid][0]) {
        if (target > matrix[mid-1][0]) {
          pickedRow = mid + 1;
          break;
        } else {
          right = mid - 1;
          pickedRow = right;
        }
      } else if (target >= matrix[mid][0]) {
        if (target < matrix[mid+1][0]) {
          pickedRow = mid;
          break;
        } else {
          left = mid + 1;
          pickedRow = left;
        }
      }
    }

    // After we found our row, we do binary search on that row
    left = 0;
    right = n-1;

    while (left <= right) {
      int mid = (right + left) / 2;
      if (target == matrix[pickedRow][mid]) {
        return true;
      } else if (target < matrix[pickedRow][mid]) {
        right = mid - 1;
      } else {
        left = mid - 1;
      }
    }

    return false;
  }
}