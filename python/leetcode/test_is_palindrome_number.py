import unittest

def is_palindrome_number(n):
    return str(n) == str(n)[::-1]


class IsPalindromeNumber(unittest.TestCase):
    def test_is_palindrome_number(self):
        self.assertEqual(True, is_palindrome_number(121))
        self.assertEqual(False, is_palindrome_number(-121))
        self.assertEqual(False, is_palindrome_number(10))


if __name__ == '__main__':
    unittest.main()
