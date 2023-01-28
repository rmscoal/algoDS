import { assert } from 'chai';
import UniqueMorseCode from '../src/leetcode/uniquemorsecode.js';

describe('Unique Morse Code', () => {
  describe('Test Case 1', () => {
    it('should output false', () => {
      const res: number = UniqueMorseCode.Solver(['gin', 'zen', 'gig', 'msg']);
      assert.equal(res, 2, 'Should match');
    });
  });

  describe('Test Case 2', () => {
    it('should output true', () => {
      const res: number = UniqueMorseCode.Solver(['a']);
      assert.equal(res, 1, 'Should match');
    });
  });
});
