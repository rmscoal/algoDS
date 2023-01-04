import { assert } from 'chai';
import CollectGarbage from '../src/leetcode/collectgarbage.js';

describe('Collect Garbage Test', () => {
  describe('CollectGarbage#Solver', () => {
    describe('Test Case 1', () => {
      it('Should return 21', () => {
        const res: number = CollectGarbage.Solver(['G', 'P', 'GP', 'GG'], [2, 4, 3]);
        assert.equal(res, 21, 'Result should be 21');
      });
    });

    describe('Test Case 2', () => {
      it('Should return 37', () => {
        const res: number = CollectGarbage.Solver(['MMM', 'PGM', 'GP'], [3, 10]);
        assert.equal(res, 37, 'Result should be 37');
      });
    });

    describe('Test Case 3', () => {
      it('Should return 8', () => {
        const res: number = CollectGarbage.Solver(['G', 'M', 'P'], [1, 3]);
        assert.equal(res, 8, 'Result should be 8');
      });
    });
  });

  describe('CollectGarbage#FastSolver', () => {
    describe('Test Case 1', () => {
      it('Should return 21', () => {
        const res: number = CollectGarbage.FastSolver(['G', 'P', 'GP', 'GG'], [2, 4, 3]);
        assert.equal(res, 21, 'Result should be 21');
      });
    });

    describe('Test Case 2', () => {
      it('Should return 37', () => {
        const res: number = CollectGarbage.FastSolver(['MMM', 'PGM', 'GP'], [3, 10]);
        assert.equal(res, 37, 'Result should be 37');
      });
    });

    describe('Test Case 3', () => {
      it('Should return 8', () => {
        const res: number = CollectGarbage.FastSolver(['G', 'M', 'P'], [1, 3]);
        assert.equal(res, 8, 'Result should be 8');
      });
    });
  });
});
