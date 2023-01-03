import { log } from 'console';
import Decibinary from './leetcode/decibinary.js';
import TwoSum from './leetcode/twosum.js';
import CollectGarbage from './leetcode/collectgarbage.js';
log('------------ Welcome to Competitive Programming ------------\n');
log('------------ Leetcode ------------');
const dec = new Decibinary('1234');
log('Decibinary Solver: ', dec.Solver());
log('Decibinary FastSolver: ', dec.FastSolver());
const twoSum = new TwoSum([2, 1, -5, 11, 15], -3);
log('Two Sum Solver: ', twoSum.Solver());
log('Two Sum SlowSolver: ', twoSum.SlowSolver());
log('Two Sum FastSolver: ', twoSum.FastSolver());
const twoSum2 = new TwoSum([3, 3], 6);
log('Two Sum Solver: ', twoSum2.Solver());
log('Two Sum SlowSolver: ', twoSum2.SlowSolver());
log('Two Sum FastSolver: ', twoSum2.FastSolver());
// Renewed Class constructor
log('Collect Garbage: ', CollectGarbage.Solver(['G', 'P', 'GP', 'GG'], [2, 4, 3]));
log('Collect Garbage: ', CollectGarbage.Solver(['MMM', 'PGM', 'GP'], [3, 10]));
log('Collect Garbage: ', CollectGarbage.Solver(['G', 'M', 'P'], [1, 3])); // => 8.. G counts only first house
//# sourceMappingURL=index.js.map