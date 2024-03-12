from typing import List


class Solution:
    def check_magazine(self, magazine: List[str], note: List[str]):
        hashmap = {}

        for word in magazine:
            hashmap[word] = hashmap.get(word, 0) + 1

        for word in note:
            if word in hashmap:
                hashmap[word] -= 1
                if hashmap[word] < 0:
                    print("No")
                    return
            else:
                print("No")
                return

        print("Yes")


if __name__ == "__main__":
    Solution().check_magazine(["two", "times", "three", "is", "not", "four"], [
        "two", "times", "two", "is", "four"])
