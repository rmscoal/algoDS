import unittest

def calculation(integer: list[int], addition: int) -> list[int]:
    '''
    Given an integer in a representation of array of integer
    and another integer, calculate the sum of the two values.
    
    Please note that you are not allowed to do type conversion.
    '''
    N = len(integer)
    number = 0
    signed = False

    for i in range(N - 1, -1, -1):
        number += (10 ** (N - 1 - i)) * integer[i]

    total = number + addition
    if total < 0:
        total *= -1
        signed = True

    result = []
    while total:
        unit = total % 10
        result.insert(0, unit)
        total //= 10

    if signed:
        result[0] *= -1

    return result


class Calculation(unittest.TestCase):
    def test(self):
        self.assertEqual([2], calculation([1], 1))
        self.assertEqual([2,0,8], calculation([2,0,5], 3))
        self.assertEqual([2,1,0], calculation([2,0,5], 5))
        self.assertEqual([9], calculation([1,0], -1))
        self.assertEqual([-1,9], calculation([3,5], -54))


if __name__ == '__main__':
    unittest.main()
