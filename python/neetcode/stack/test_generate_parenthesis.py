from typing import List, Dict, Set, Tuple
from collections import Counter
from utils import st_time
import unittest


class TestGenerateParenthesisMyElegantSolution(unittest.TestCase):
    '''
         (
        / \
       ((  ()
      / \   |
    ((( (() ()(
    .   .   .
    .   .   .
    '''
    class ParenTree:
        def __init__(self, left: int, right: int, val: str):
            self.left = left
            self.right = right
            self.val = val

        def duplicate(self) -> Tuple['TestGenerateParenthesisMyElegantSolution.ParenTree', 'TestGenerateParenthesisMyElegantSolution.ParenTree']:
            left = TestGenerateParenthesisMyElegantSolution.ParenTree(
                left=self.left, right=self.right, val=self.val)
            right = TestGenerateParenthesisMyElegantSolution.ParenTree(
                left=self.left, right=self.right, val=self.val)

            return left, right

        def __repr__(self) -> str:
            return 'Left: {}, right: {} and val: {}'.format(self.left, self.right, self.val)

    @st_time
    def generate_parenthesis_tree(self, n: int) -> List[str]:
        stack: List[TestGenerateParenthesisMyElegantSolution.ParenTree] = [
            TestGenerateParenthesisMyElegantSolution.ParenTree(left=1, right=0, val='(')]

        for _ in range(0, 2 * n - 1):
            tmp = stack
            stack = []
            for i in range(0, len(tmp)):
                left_clone, right_clone = tmp[i].duplicate()
                if left_clone.left < n:
                    left_clone.left += 1
                    left_clone.val += '('
                    stack.append(left_clone)

                if right_clone.left > right_clone.right:
                    right_clone.right += 1
                    right_clone.val += ')'
                    stack.append(right_clone)

        result = [node.val for node in stack if node.left ==
                  n and node.right == n]
        return result

    def test_generate_parenthesis(self):
        self.assertListEqual(["()"], self.generate_parenthesis_tree(1))
        self.assertListEqual(
            ["(())", "()()"], self.generate_parenthesis_tree(2))
        self.assertListEqual(["((()))", "(()())", "(())()",
                             "()(())", "()()()"], self.generate_parenthesis_tree(3))
        self.assertListEqual(["(((())))", "((()()))", "((())())", "((()))()", "(()(()))", "(()()())", "(()())()", "(())(())",
                             "(())()()", "()((()))", "()(()())", "()(())()", "()()(())", "()()()()"], self.generate_parenthesis_tree(4))
        self.assertListEqual(["((((()))))", "(((()())))", "(((())()))", "(((()))())", "(((())))()", "((()(())))", "((()()()))", "((()())())", "((()()))()", "((())(()))", "((())()())", "((())())()", "((()))(())", "((()))()()", "(()((())))", "(()(()()))", "(()(())())", "(()(()))()", "(()()(()))", "(()()()())", "(()()())()", "(()())(())",
                             "(()())()()", "(())((()))", "(())(()())", "(())(())()", "(())()(())", "(())()()()", "()(((())))", "()((()()))", "()((())())", "()((()))()", "()(()(()))", "()(()()())", "()(()())()", "()(())(())", "()(())()()", "()()((()))", "()()(()())", "()()(())()", "()()()(())", "()()()()()"], self.generate_parenthesis_tree(5))


class TestGenerateParenthesisBackTracking(unittest.TestCase):
    @st_time
    def generate_parenthesis_backtracking(self, n: int) -> List[str]:
        def backtracking(left, right, val):
            if len(val) == 2 * n:
                result.append(val)
            if left < n:
                backtracking(left + 1, right, val + '(')
            if right < left:
                backtracking(left, right+1, val + ')')

        result = []
        backtracking(1, 0, '(')

        return result

    def test_generate_parenthesis(self):
        self.assertListEqual(["()"], self.generate_parenthesis_backtracking(1))
        self.assertListEqual(
            ["(())", "()()"], self.generate_parenthesis_backtracking(2))
        self.assertListEqual(["((()))", "(()())", "(())()",
                             "()(())", "()()()"], self.generate_parenthesis_backtracking(3))
        self.assertListEqual(["(((())))", "((()()))", "((())())", "((()))()", "(()(()))", "(()()())", "(()())()", "(())(())",
                             "(())()()", "()((()))", "()(()())", "()(())()", "()()(())", "()()()()"], self.generate_parenthesis_backtracking(4))
        self.assertListEqual(["((((()))))", "(((()())))", "(((())()))", "(((()))())", "(((())))()", "((()(())))", "((()()()))", "((()())())", "((()()))()", "((())(()))", "((())()())", "((())())()", "((()))(())", "((()))()()", "(()((())))", "(()(()()))", "(()(())())", "(()(()))()", "(()()(()))", "(()()()())", "(()()())()", "(()())(())",
                             "(()())()()", "(())((()))", "(())(()())", "(())(())()", "(())()(())", "(())()()()", "()(((())))", "()((()()))", "()((())())", "()((()))()", "()(()(()))", "()(()()())", "()(()())()", "()(())(())", "()(())()()", "()()((()))", "()()(()())", "()()(())()", "()()()(())", "()()()()()"], self.generate_parenthesis_backtracking(5))


if __name__ == "__main__":
    unittest.main()
