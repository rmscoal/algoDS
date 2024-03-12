import unittest


class Solution:
    yes = "YES"
    no = "NO"

    def two_strings(self, s1: str, s2: str) -> str:
        hashmap = {}
        for c in s1:
            hashmap[c] = True
        for c in s2:
            if c in hashmap:
                return self.yes

        return self.no


class TestSolution(unittest.TestCase):
    def test_two_strings(self):
        self.assertEqual("YES", Solution().two_strings("hello", "world"))
        self.assertEqual("NO", Solution().two_strings("hi", "world"))


if __name__ == "__main__":
    unittest.main()
