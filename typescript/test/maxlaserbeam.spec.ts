import { assert } from 'chai';
import MaxLaserBeam from '../src/leetcode/maxlaserbeam.js';

describe('MaxLaserBeams Test', () => {
  describe('MaxLaserBeams#Solver', () => {
    describe('Testcase One', () => {
      it('Returns 8', () => {
        assert.equal(MaxLaserBeam.Solver(['011001', '000000', '010100', '001000']), 8, 'Should returns 8');
      });
    });
    describe('Testcase Two', () => {
      it('Returns 0', () => {
        assert.equal(MaxLaserBeam.Solver(['0000', '0001', '0000']), 0, 'Should returns 0');
      });
    });
  });
});
