import { log } from 'console';
import Decibinary from './leetcode/decibinary.js';
import TwoSum from './leetcode/twosum.js';

log('------------ Welcome to Competitive Programming ------------\n');

log('------------ Leetcode ------------');

const dec: Decibinary = new Decibinary('1234');
log('Decibinary Solver: ', dec.Solver());
log('Decibinary FastSolver: ', dec.FastSolver());

const twoSum: TwoSum = new TwoSum([2, 1, -5, 11, 15], -3);
log('Two Sum Solver: ', twoSum.Solver());
log('Two Sum SlowSolver: ', twoSum.SlowSolver());
log('Two Sum FastSolver: ', twoSum.FastSolver());
const twoSum2: TwoSum = new TwoSum([3, 3], 6);
log('Two Sum Solver: ', twoSum2.Solver());
log('Two Sum SlowSolver: ', twoSum2.SlowSolver());
log('Two Sum FastSolver: ', twoSum2.FastSolver());
