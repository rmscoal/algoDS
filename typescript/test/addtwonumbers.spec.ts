/* eslint-disable max-len */
import { assert } from 'chai';
import AddTwoNumbers from '../src/leetcode/addtwonumbers.js';
import { ListNode } from '../src/leetcode/types.js';

describe('Add Two Numbers Test', () => {
  describe('CleanSolver', () => {
    describe('Test case one', () => {
      it('Should be equal', () => {
        const correctAnswer = ListNode.newFromArray([7, 0, 8]);
        const res = AddTwoNumbers.CleanSolver(ListNode.newFromArray([2, 4, 3]), ListNode.newFromArray([5, 6, 4]));
        assert.equal(ListNode.equal(res, correctAnswer), true);
      });
    });
  });
});
