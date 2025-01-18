import unittest

def number_of_operations(binary: str) -> int:
    binary = binary.lstrip("0")

    times = 0
    while binary != "0":
        N = len(binary)
        if binary[N - 1] == "1":
            binary = binary[:N - 1] + "0"
        else:
            binary = binary[:N - 1]
        times += 1

    return times

class BCA(unittest.TestCase):
    def test_number_of_operations(self):
        self.assertEqual(5, number_of_operations("111"))
        self.assertEqual(5, number_of_operations("000000111"))
        self.assertEqual(14, number_of_operations("00111001101"))


if __name__ == '__main__':
    unittest.main()
