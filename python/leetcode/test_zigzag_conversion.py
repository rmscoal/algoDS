import unittest

def zigzag_conversion(s: str, row: int) -> str:
    N, current_idx = len(s), 0
    matrix, current_column = [], []
    going_diagonal = False

    while current_idx < N:
        if going_diagonal:
            for i in range(row - 2, 0, -1):
                for j in range(row):
                    if j == i and current_idx < N:
                        current_column.append(s[current_idx])
                        current_idx += 1
                    else:
                        current_column.append("")
                matrix.append(current_column)
                current_column = []
            going_diagonal = False
        else:
            for i in range(row):
                if current_idx < N:
                    current_column.append(s[current_idx])
                    current_idx += 1
                else:
                    current_column.append("")
            matrix.append(current_column)
            current_column = []
            going_diagonal = True

    # Build the string result
    res = ""
    for r in range(row):
        for l in range(len(matrix)):
            res += matrix[l][r]
    return res

def better_zigzag_conversion(s: str, row: int) -> str:
    if row == 1 or row > len(s):
        return s

    rows = ['' for _ in range(min(row, len(s)))]
    current_row = 0
    going_down = False

    for char in s:
        rows[current_row] += char
        # Change direction when we are
        # at the start or end of the row
        if current_row == 0 or current_row == row - 1:
            going_down = not going_down
        current_row += 1 if going_down else -1

    return ''.join(rows)

class ZigzagConversion(unittest.TestCase):
    def test(self):
        self.assertEqual("PAHNAPLSIIGYIR", zigzag_conversion("PAYPALISHIRING", 3))
        self.assertEqual("PINALSIGYAHRPI", zigzag_conversion("PAYPALISHIRING", 4))
    def test_better(self):
        self.assertEqual("PAHNAPLSIIGYIR", better_zigzag_conversion("PAYPALISHIRING", 3))
        self.assertEqual("PINALSIGYAHRPI", better_zigzag_conversion("PAYPALISHIRING", 4))
        self.assertEqual("AB", better_zigzag_conversion("AB", 1))

if __name__ == '__main__':
    unittest.main()
