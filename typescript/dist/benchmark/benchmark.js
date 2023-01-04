/* eslint-disable func-names */
import Benchmark from 'benchmark';
import { log } from 'console';
import Decibinary from '../src/leetcode/decibinary.js';
import TwoSum from '../src/leetcode/twosum.js';
import CollectGarbage from '../src/leetcode/collectgarbage.js';
export default class CustomBenchmark {
    static DecibinaryBenchmark(s) {
        this.suite = new Benchmark.Suite();
        const dec = new Decibinary(s);
        this.suite.add('Decibinary#Solver', () => {
            dec.Solver();
        })
            .add('Decibinary#FastSolver', () => {
            dec.FastSolver();
        })
            // add listeners
            .on('cycle', (event) => {
            log(String(event.target));
        })
            .on('complete', function () {
            log(`The fastest method is ${this.filter('fastest').map('name')}`);
        })
            // run async or not
            .run({ async: false });
    }
    static TwoSumBenchmark(nums, target) {
        this.suite = new Benchmark.Suite();
        const twoSum = new TwoSum(nums, target);
        this.suite.add('TwoSum#Solver', () => {
            twoSum.Solver();
        })
            .add('TwoSum#SlowSolver', () => {
            twoSum.SlowSolver();
        })
            .add('TwoSum#FastSolver', () => {
            twoSum.FastSolver();
        })
            // add listeners
            .on('cycle', (event) => {
            log(String(event.target));
        })
            .on('complete', function () {
            log(`The fastest method is ${this.filter('fastest').map('name')}`);
        })
            // run async or not
            .run({ async: false });
    }
    static CollectGarbageBenchmark(garbage, travel) {
        this.suite = new Benchmark.Suite();
        this.suite.add('CollectGarbage#Solver', () => {
            CollectGarbage.Solver(garbage, travel);
        })
            .add('CollectGarbage#FastSolver', () => {
            CollectGarbage.FastSolver(garbage, travel);
        })
            // add listeners
            .on('cycle', (event) => {
            log(String(event.target));
        })
            .on('complete', function () {
            log(`The fastest method is ${this.filter('fastest').map('name')}`);
        })
            // run async or not
            .run({ async: false });
    }
}
//# sourceMappingURL=benchmark.js.map