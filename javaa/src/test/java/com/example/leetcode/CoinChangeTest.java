package com.example.leetcode;

import static org.junit.jupiter.api.Assertions.assertEquals;

import java.util.Arrays;

import org.junit.jupiter.api.Test;

public class CoinChangeTest {
  public int coinChange(int[] coins, int amount) {
    // coins = [1,2,5], amount = 11
    // [0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0]
    int[] amounts = new int[amount + 1];
    // We fill the amounts with minimum values
    // except for the first element.
    Arrays.fill(amounts, Integer.MAX_VALUE - 1);
    amounts[0] = 0;

    // Iterate through all the amounts sequentially
    for (int i = 1; i < amount + 1; i++) {
      for (int coin : coins) {
        // If the coin's value is the same as the
        // current amount, set it to 1 immediately
        // the break and go to the next amount.
        if (coin == i) {
          amounts[i] = 1;
          break;
        }

        // If the coin's value is greater than the
        // current amount, just go to the next amounts.
        if (coin > i) {
          continue;
        }

        amounts[i] = Math.min(amounts[i], amounts[i - coin] + 1);
      }
    }

    return amounts[amount] == Integer.MAX_VALUE - 1 ? -1 : amounts[amount];
  }

  @Test
  public void coinChangeTest() {
    CoinChangeTest cc = new CoinChangeTest();

    assertEquals(3, cc.coinChange(new int[]{5,2,1}, 11));
    assertEquals(-1, cc.coinChange(new int[]{2}, 3));
    assertEquals(0, cc.coinChange(new int[]{1}, 0));
    assertEquals(20, cc.coinChange(new int[]{186,419,83,408}, 6249));
  }
}
