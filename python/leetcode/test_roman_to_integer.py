import unittest

def roman_to_integer(s):
    N = len(s)
    total = 0
    for i in range(N):
        curr = s[i]
        prev = s[max(i - 1, 0)]
        match curr:
            case "I":
                total += 1
            case "V":
                if prev == "I":
                    total -= 2
                total += 5
            case "X":
                if prev == "I":
                    total -= 2
                total += 10
            case "L":
                if prev == "X":
                    total -= 20
                total += 50
            case "C":
                if prev == "X":
                    total -= 20
                total += 100
            case "D":
                if prev == "C":
                    total -= 200
                total += 500
            case "M":
                if prev == "C":
                    total -= 200
                total += 1000
    return total


class RomanToInteger(unittest.TestCase):
    def test_roman_to_integer(self):
        self.assertEqual(3, roman_to_integer("III"))
        self.assertEqual(58, roman_to_integer("LVIII"))
        self.assertEqual(1994, roman_to_integer("MCMXCIV"))


if __name__ == '__main__':
    unittest.main()
