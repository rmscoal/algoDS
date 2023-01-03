/* eslint-disable func-names */
import Benchmark from 'benchmark';
import { log } from 'console';
import Decibinary from '../src/leetcode/decibinary.js';
import TwoSum from '../src/leetcode/twosum.js';

export default class CustomBenchmark {
  static suite: Benchmark.Suite;

  public static Decibinary(s: string) {
    this.suite = new Benchmark.Suite();

    const dec: Decibinary = new Decibinary(s);

    this.suite.add('Decibinary#Solver', () => {
      dec.Solver();
    })
      .add('Decibinary#FastSolver', () => {
        dec.FastSolver();
      })
      // add listeners
      .on('cycle', (event: { target: any; }) => {
        log(String(event.target));
      })
      .on('complete', function () {
        log(`The fastest method is ${this.filter('fastest').map('name')}`);
      })
      // run async or not
      .run({ async: false });
  }

  public static TwoSum(nums: Array<number>, target: number) {
    this.suite = new Benchmark.Suite();

    const twoSum: TwoSum = new TwoSum(nums, target);

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
      .on('cycle', (event: { target: any; }) => {
        log(String(event.target));
      })
      .on('complete', function () {
        log(`The fastest method is ${this.filter('fastest').map('name')}`);
      })
      // run async or not
      .run({ async: false });
  }
}
