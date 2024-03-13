from typing import List
from math import trunc
import unittest


class TestReversePolishNotation(unittest.TestCase):
    @st_time
    def eval_rpn_slow(self, tokens: List[str]) -> int:
        stack: List[int] = []

        # For example we have tokens as ["4","13","5","/","+"]
        # stack = [4, 13, 5]
        # eval(13/5)
        # stack = [4, 2]
        # eval(4 + 2)
        # stack = [6]
        #
        # ["2","1","+","3","*"]
        # stack = [2, 1]
        # eval(2 + 1)
        # stack = [3]
        # eval(3 * 3)
        # stack = [9]

        for string in tokens:
            try:
                stack.append(int(string))
            except ValueError:
                second_num = stack.pop()
                first_num = stack.pop()
                result = trunc(
                    eval(str(first_num) + string + str(second_num)))
                stack.append(result)

        return stack[0]

    def test_eval_rpn_slow(self):
        self.assertEqual(self.eval_rpn_slow(['1', '1', '+']), 2, "Should be 2")
        self.assertEqual(self.eval_rpn_slow(
            ["2", "1", "+", "3", "*"]), 9, "Should be 9")
        self.assertEqual(self.eval_rpn_slow(
            ["4", "13", "5", "/", "+"]), 6, "Should be 6")
        self.assertEqual(self.eval_rpn_slow(
            ["10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"]), 22, "Should be 22")

    @st_time
    def eval_rpn(self, tokens: List[str]) -> int:
        stack: List[int] = []

        for string in tokens:
            try:
                stack.append(int(string))
            except ValueError:
                second_num = stack.pop()
                first_num = stack.pop()
                if string == "/":
                    stack.append(trunc(first_num/second_num))
                elif string == "*":
                    stack.append(first_num*second_num)
                elif string == "-":
                    stack.append(first_num-second_num)
                elif string == "+":
                    stack.append(first_num+second_num)
        return stack[0]

    def test_eval_rpn(self):
        self.assertEqual(self.eval_rpn(['1', '1', '+']), 2, "Should be 2")
        self.assertEqual(self.eval_rpn(
            ["2", "1", "+", "3", "*"]), 9, "Should be 9")
        self.assertEqual(self.eval_rpn(
            ["4", "13", "5", "/", "+"]), 6, "Should be 6")
        self.assertEqual(self.eval_rpn(
            ["10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"]), 22, "Should be 22")


if __name__ == "__main__":
    unittest.main()
