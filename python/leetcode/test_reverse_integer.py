import unittest

MAX_INT = 2147483647
MIN_INT = -2147483647

def reverse_integer(n):
    signed = False
    if n < 0:
        n *= -1
        signed = True

    # Split the integers into list of
    # integers while reversing at the
    # same time
    arr = []
    while n:
        arr.append(n % 10)
        n //= 10

    N, m = len(arr)-1, 0
    for i in range(N, -1, -1):
        m += arr[i] * (10 ** (N-i))

    m *= -1 if signed else 1
    if m > MAX_INT or  m < MIN_INT:
        return 0
    else:
        return m


class ReverseIntegerTest(unittest.TestCase):
    def test_reverse_integer(self):
        self.assertEqual(123, reverse_integer(321))
        self.assertEqual(21, reverse_integer(120))
        self.assertEqual(0, reverse_integer(2147483646))
        self.assertEqual(-321, reverse_integer(-123))

if __name__ == '__main__':
    unittest.main()