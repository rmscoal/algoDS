package com.example.leetcode;

import static org.junit.jupiter.api.Assertions.assertEquals;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

import org.junit.jupiter.api.Test;

class MinStack {
  private List<Integer> arr;
  private Map<Integer, Integer> hash;
  private int currentMin;

  public MinStack() {
    // The list as usual
    this.arr = new ArrayList<Integer>();
    // Current minimum index to previous minimum index
    this.hash = new HashMap<Integer, Integer>();
  }

  public void push(int val) {
    if (this.hash.size() == 0) {
      currentMin = 0;
      this.hash.put(currentMin, currentMin);
    } else if (val < this.arr.get(this.currentMin)) {
      int tmp = this.arr.size();
      this.hash.put(tmp, currentMin);
      currentMin = tmp;
    }

    this.arr.add(val);
  }

  public void pop() {
    if (currentMin == this.arr.size() - 1) {
      int tmp = this.hash.get(currentMin);
      this.hash.remove(currentMin);
      currentMin = tmp;
    }
    this.arr.removeLast();
  }

  public int top() {
    return this.arr.get(this.arr.size() - 1);
  }

  public int getMin() {
    return this.arr.get(currentMin);
  }
}

class MinStackTest {
  @Test
  void minstackTest() {
    MinStack minstack = new MinStack();

    int result;

    minstack.push(-2);
    minstack.push(0);
    minstack.push(-3);
    result = minstack.getMin();
    assertEquals(-3, result);
    minstack.pop();
    assertEquals(0, minstack.top());
    assertEquals(-2, minstack.getMin());

  }
}

class MinStack2 {
  private List<Integer> arr;
  private List<Integer> mins;

  public MinStack2() {
    this.arr = new ArrayList<>();
    this.mins = new ArrayList<>();
  }

  public void push(int val) {
    if (this.arr.size() == 0) {
      this.mins.add(0);
    } else if (val < this.mins.getLast()) {
      this.mins.add(this.arr.size());
    }
    this.arr.add(val);
  }

  public void pop() {
    if (this.arr.size() - 1 == this.mins.getLast()) {
      this.mins.removeLast();
    }
    this.arr.removeLast();
  }

  public int top() {
    return this.arr.getLast();
  }

  public int getMin() {
    return this.arr.get(this.mins.getLast());
  }
}

class MinStackTest2 {
  @Test
  void minstackTest() {
    MinStack2 minstack = new MinStack2();

    int result;

    minstack.push(-2);
    minstack.push(0);
    minstack.push(-3);
    result = minstack.getMin();
    assertEquals(-3, result);
    minstack.pop();
    assertEquals(0, minstack.top());
    assertEquals(-2, minstack.getMin());

  }
}
