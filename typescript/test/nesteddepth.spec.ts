import { assert } from 'chai';
import NestingDepth from '../src/leetcode/nestingdepth.js';

describe('Nesting Depth Test', () => {
  describe('Nesting Depth#Solver', () => {
    describe('Testcase One', () => {
      it('Returns 3', () => {
        assert.equal(NestingDepth.Solver('(1+(2*3)+((8)/4))+1'), 3, 'Should return 3');
      });
    });

    describe('Testcase Two', () => {
      it('Returns 3', () => {
        assert.equal(NestingDepth.Solver('(1)+((2))+(((3)))'), 3, 'Should return 3');
      });
    });
  });
});
