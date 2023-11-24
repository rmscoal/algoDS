package com.example.leetcode;

import static org.junit.jupiter.api.Assertions.assertEquals;

import org.junit.jupiter.api.Test;

public class ContainsDuplicateTest {
  @Test
  void containsDuplicateTest() {
    assertEquals(true, ContainsDuplicate.containsDuplicateHash(new int[] { 1, 2, 3, 1 }));
    assertEquals(true, ContainsDuplicate.containsDuplicate(new int[] { 1, 2, 3, 1 }));
  }

  @Test
  void containsDuplicate_II_Test() {
    assertEquals(false, ContainsDuplicate.containsNearbyDuplicate(new int[] { 1, 2, 3, 1, 2, 3 }, 2));
    assertEquals(true, ContainsDuplicate.containsNearbyDuplicate(new int[] { 1, 2, 3, 1 }, 3));
    assertEquals(true, ContainsDuplicate.containsNearbyDuplicate(new int[] { 1, 0, 1, 1 }, 1));
    assertEquals(false, ContainsDuplicate.containsNearbyDuplicate(new int[] { 1 }, 1));
    assertEquals(true, ContainsDuplicate.containsNearbyDuplicate(new int[] { 99, 99 }, 2));
    assertEquals(false, ContainsDuplicate.containsNearbyDuplicate(new int[] { 99, 98 }, 2));

    assertEquals(false, ContainsDuplicate.containsNearbyDuplicateV2(new int[] { 1, 2, 3, 1, 2, 3 }, 2));
    assertEquals(true, ContainsDuplicate.containsNearbyDuplicateV2(new int[] { 1, 2, 3, 1 }, 3));
    assertEquals(true, ContainsDuplicate.containsNearbyDuplicateV2(new int[] { 1, 0, 1, 1 }, 1));
    assertEquals(false, ContainsDuplicate.containsNearbyDuplicateV2(new int[] { 1 }, 1));
    assertEquals(true, ContainsDuplicate.containsNearbyDuplicateV2(new int[] { 99, 99 }, 2));
    assertEquals(false, ContainsDuplicate.containsNearbyDuplicateV2(new int[] { 99, 98 }, 2));

    assertEquals(false, ContainsDuplicate.containsNearbyDuplicateV3(new int[] { 1, 2, 3, 1, 2, 3 }, 2));
    assertEquals(true, ContainsDuplicate.containsNearbyDuplicateV3(new int[] { 1, 2, 3, 1 }, 3));
    assertEquals(true, ContainsDuplicate.containsNearbyDuplicateV3(new int[] { 1, 0, 1, 1 }, 1));
    assertEquals(false, ContainsDuplicate.containsNearbyDuplicateV3(new int[] { 1 }, 1));
    assertEquals(true, ContainsDuplicate.containsNearbyDuplicateV3(new int[] { 99, 99 }, 2));
    assertEquals(false, ContainsDuplicate.containsNearbyDuplicateV3(new int[] { 99, 98 }, 2));
  }

  @Test
  void containsDuplicate_III_Test() {
    assertEquals(true, ContainsDuplicate.containsNearbyAlmostDuplicate(new int[] { 1, 2, 3, 1 }, 3, 0));
    assertEquals(false, ContainsDuplicate.containsNearbyAlmostDuplicate(new int[] { 1, 5, 9, 1, 5, 9 }, 2, 3));
    assertEquals(true, ContainsDuplicate.containsNearbyAlmostDuplicate(new int[] { -3, 3, -6 }, 2, 3));
  }
}
