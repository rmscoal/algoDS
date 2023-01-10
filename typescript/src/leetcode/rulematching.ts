/**
 * A class to solve Count Items Matching a Rule.
 *
 * Problem:
 * You are given an array items, where each items[i] = [type_i, color_i, name_i]
 * describes the type, color, and name of the ith item. You are also given a
 * rule represented by two strings, ruleKey and ruleValue.
 *
 * The ith item is said to match the rule if one of the following is true:
 * ruleKey == "type" and ruleValue == type_i.
 * ruleKey == "color" and ruleValue == color_i.
 * ruleKey == "name" and ruleValue == name_i.
 *
 * Return the number of items that match the given rule.
 */
export default class RuleMatching {
  /**
   * Solver
   * @param items 2d array where items[i] = [type_i, color_i, name_i]
   * @param ruleKey the rule to apply
   * @param ruleValue the value of the rule it should abide
   * @returns the number of items that match the given rule.
   */
  public static Solver(items: string[][], ruleKey: string, ruleValue: string): number {
    // Returned variable
    let count: number = 0;
    // Variable to hold what value to check given the rule key.
    let ruleKeyIndex: number = 0;

    // Route the ruleKey to index
    switch (ruleKey) {
      case 'type': {
        ruleKeyIndex = 0;
        break;
      }

      case 'color': {
        ruleKeyIndex = 1;
        break;
      }

      default: {
        ruleKeyIndex = 2;
      }
    }

    for (let i = 0; i < items.length; i += 1) {
      if (ruleValue === items[i][ruleKeyIndex]) {
        count += 1;
      }
    }

    return count;
  }

  /**
   * FastSolver
   * This solver has slight the same performance as Solver although coming
   * a tiny bit better. The differences is too small it is insignificant.
   * @param items 2d array where items[i] = [type_i, color_i, name_i]
   * @param ruleKey the rule to apply
   * @param ruleValue the value of the rule it should abide
   * @returns the number of items that match the given rule.
   */
  public static FastSolver(items: string[][], ruleKey: string, ruleValue: string): number {
    // Returned variable
    let count: number = 0;
    // Variable to hold what value to check given the rule key.
    let ruleKeyIndex: number = 0;

    // Route the ruleKey to index
    if (ruleKey === 'type') ruleKeyIndex = 0;
    if (ruleKey === 'color') ruleKeyIndex = 1;
    if (ruleKey === 'name') ruleKeyIndex = 2;

    for (let i = 0; i < items.length; i += 1) {
      if (ruleValue === items[i][ruleKeyIndex]) {
        count += 1;
      }
    }

    return count;
  }
}
