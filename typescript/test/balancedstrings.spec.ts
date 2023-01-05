import { assert } from 'chai';
import BalancedStrings from '../src/leetcode/balancedstrings.js';

describe('Balanced Strings Test', () => {
  describe('BalancedStrings#Solver', () => {
    describe('Test Case 1', () => {
      it('Should return 1', () => {
        const res: number = BalancedStrings.Solver('RRRRRRRRLLLLLLLL');
        assert.equal(res, 1, 'Result should be 1');
      });
    });

    describe('Test Case 2', () => {
      it('Should return 8', () => {
        const res: number = BalancedStrings.Solver('RLRRRLLLLRRLLRRLLLRRRRRLLL');
        assert.equal(res, 8, 'Result should be 8');
      });
    });

    describe('Test Case 3', () => {
      it('Should return 4', () => {
        const res: number = BalancedStrings.Solver('RLRRLLRLRL');
        assert.equal(res, 4, 'Result should be 4');
      });
    });
  });
});
