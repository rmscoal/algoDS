import { assert } from 'chai';
import TwoStringArray from '../src/leetcode/twostringarray.js';

describe('Equivalence Two String Array', () => {
  describe('Test Case 1', () => {
    it('should output false', () => {
      const res: boolean = TwoStringArray.Solver(['a', 'cb'], ['ab', 'c']);
      assert.equal(res, false, 'Should match');
    });
  });

  describe('Test Case 2', () => {
    it('should output true', () => {
      const res: boolean = TwoStringArray.Solver(['abc', 'd', 'defg'], ['abcddefg']);
      assert.equal(res, true, 'Should match');
    });
  });
});
