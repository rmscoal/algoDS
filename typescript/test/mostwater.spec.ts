import { assert } from 'chai';
import MostWater from '../src/leetcode/mostwater.js';

describe('Container with Most Water', () => {
  it('Should return 49', () => {
    const res = MostWater.PerfSolver([1, 8, 6, 2, 5, 4, 8, 3, 7]);
    assert.equal(res, 49);
  });

  it('Should return 1', () => {
    const res = MostWater.PerfSolver([1, 1]);
    assert.equal(res, 1);
  });
});
