package com.example.leetcode;

import static org.junit.jupiter.api.Assertions.assertEquals;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

import org.junit.jupiter.api.Test;

public class GenerateParenthesisTest {
  public List<String> generateParenthesisBacktrack(int n) {
    ArrayList<String> result = new ArrayList<String>();
    backtrack("(", 1, 0, result, n);
    return result;
  }

  void backtrack(String s, int left, int right, List<String> result, int n) {
    if (s.length() == 2 * n) {
      result.add(s);
    }

    if (left < n) {
      backtrack(s + '(', left + 1, right, result, n);
    }

    if (right < left) {
      backtrack(s + ')', left, right + 1, result, n);
    }
  }

  @Test
  void testGenerateParenthesisBacktracking() {
    assertEquals(new ArrayList<String>(Arrays.asList(new String[] { "()" })), generateParenthesisBacktrack(1));
    assertEquals(new ArrayList<String>(Arrays.asList(new String[] { "(())", "()()" })), generateParenthesisBacktrack(2));
    assertEquals(new ArrayList<String>(Arrays.asList(new String[] { "((()))","(()())","(())()","()(())","()()()" })), generateParenthesisBacktrack(3));
    assertEquals(new ArrayList<String>(Arrays.asList(new String[] { "(((())))","((()()))","((())())","((()))()","(()(()))","(()()())","(()())()","(())(())","(())()()","()((()))","()(()())","()(())()","()()(())","()()()()" })), generateParenthesisBacktrack(4));
  }


  class ParenTree {
    String val;
    int left;
    int right;

    public ParenTree(String val, int left, int right) {
      this.val = val;
      this.left = left;
      this.right = right;
    }
  }

  public List<String> generateParenthesisStack(int n) {
    List<ParenTree> stack = new ArrayList<ParenTree>(Arrays.asList(new ParenTree("(", 1, 0)));

    for (int i = 0; i < 2 * n - 1; i++) {
      List<ParenTree> tmp = new ArrayList<ParenTree>(stack);
      stack = new ArrayList<>();
      for (ParenTree node : tmp) {
        ParenTree leftClone = new ParenTree(node.val, node.left, node.right);
        ParenTree rightClone = new ParenTree(node.val, node.left, node.right);

        if (leftClone.left < n) {
          leftClone.left++;
          leftClone.val += '(';
          stack.add(leftClone);
        }

        if (rightClone.right < rightClone.left) {
          rightClone.right++;
          rightClone.val += ')';
          stack.add(rightClone);
        }
      }
    }

    List<String> result = new ArrayList<String>();
    for (ParenTree node : stack) {
      if (node.left == node.right && node.left == n) {
        result.add(node.val);
      }
    }
    return result;
  }

   @Test
  void testGenerateParenthesisStack() {
    assertEquals(new ArrayList<String>(Arrays.asList(new String[] { "()" })), generateParenthesisStack(1));
    assertEquals(new ArrayList<String>(Arrays.asList(new String[] { "(())", "()()" })), generateParenthesisStack(2));
    assertEquals(new ArrayList<String>(Arrays.asList(new String[] { "((()))","(()())","(())()","()(())","()()()" })), generateParenthesisStack(3));
    assertEquals(new ArrayList<String>(Arrays.asList(new String[] { "(((())))","((()()))","((())())","((()))()","(()(()))","(()()())","(()())()","(())(())","(())()()","()((()))","()(()())","()(())()","()()(())","()()()()" })), generateParenthesisStack(4));
  }
}
