import { assert } from 'chai';
import SentenceSorting from '../src/leetcode/sentencesorting.js';

describe('Sentence Sorting', () => {
  describe('Test Case 1', () => {
    it('should output a reconstructed string', () => {
      const res: string = SentenceSorting.Solver('is2 sentence4 This1 a3');
      assert.equal(res, 'This is a sentence', 'Should match');
    });
  });

  describe('Test Case 2', () => {
    it('should output a reconstructed string', () => {
      const res: string = SentenceSorting.Solver('Myself2 Me1 I4 and3');
      assert.equal(res, 'Me Myself and I', 'Should match');
    });
  });
});
