import unittest


class Solution:
    # Neetcode Solution
    def check_inclusion(self, s1: str, s2: str) -> bool:
        if len(s1) > len(s2):
            return False

        s1_len = len(s1)
        s2_len = len(s2)
        s1_freq = [0] * 26
        s2_freq = [0] * 26
        a = ord("a")

        for i in range(s1_len):
            s1_freq[ord(s1[i]) - a] += 1
            s2_freq[ord(s2[i]) - a] += 1

        matches = 0
        for i in range(26):
            matches += 1 if s1_freq[i] == s2_freq[i] else 0

        for i in range(s1_len, s2_len):
            if matches == 26:
                return True

            # Remove left
            freq_index = ord(s2[i - s1_len]) - a
            s2_freq[freq_index] -= 1
            if s1_freq[freq_index] == s2_freq[freq_index]:
                matches += 1
            elif s2_freq[freq_index] + 1 == s1_freq[freq_index]:
                matches -= 1

            # Add right
            freq_index = ord(s2[i]) - a
            s2_freq[freq_index] += 1
            if s1_freq[freq_index] == s2_freq[freq_index]:
                matches += 1
            elif s1_freq[freq_index] + 1 == s2_freq[freq_index]:
                matches -= 1

        return matches == 26


class TestSolution(unittest.TestCase):
    def test_check_inclusion(self):
        s1 = "ab"
        s2 = "eidbaooo"
        self.assertTrue(Solution().check_inclusion(s1, s2))

        s1 = "ab"
        s2 = "eidboaoo"
        self.assertFalse(Solution().check_inclusion(s1, s2))

        s1 = "abc"
        s2 = "bbbca"
        self.assertTrue(Solution().check_inclusion(s1, s2))


if __name__ == "__main__":
    unittest.main()
