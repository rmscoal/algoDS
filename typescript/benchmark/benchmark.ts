/* eslint-disable func-names */
import { log } from 'console';
import Benchmark from 'benchmark';
import Decibinary from '../src/leetcode/decibinary.js';
import TwoSum from '../src/leetcode/twosum.js';
import CollectGarbage from '../src/leetcode/collectgarbage.js';
import MaxProfit from '../src/leetcode/maxprofit.js';
import BalancedStrings from '../src/leetcode/balancedstrings.js';
import RuleMatching from '../src/leetcode/rulematching.js';
import SentenceSorting from '../src/leetcode/sentencesorting.js';
import PangramSentence from '../src/leetcode/pangramsentence.js';
import TwoStringArray from '../src/leetcode/twostringarray.js';
import StayingIngrid from '../src/leetcode/stayingingrid.js';
import UniqueMorseCode from '../src/leetcode/uniquemorsecode.js';
import MaxLaserBeam from '../src/leetcode/maxlaserbeam.js';

export default class CustomBenchmark {
  static suite: Benchmark.Suite;

  public static DecibinaryBenchmark(s: string) {
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

  public static TwoSumBenchmark(nums: Array<number>, target: number) {
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

  public static CollectGarbageBenchmark(garbage: Array<string>, travel: Array<number>) {
    this.suite = new Benchmark.Suite();

    this.suite.add('CollectGarbage#Solver', () => {
      CollectGarbage.Solver(garbage, travel);
    })
      .add('CollectGarbage#FastSolver', () => {
        CollectGarbage.FastSolver(garbage, travel);
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

  public static MaxProfitBenchmark(input: Array<number>) {
    this.suite = new Benchmark.Suite();

    this.suite.add('MaxProfit#Solver', () => {
      MaxProfit.Solver(input);
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

  public static BalancedStringsBenchmark(input: string) {
    this.suite = new Benchmark.Suite();

    this.suite.add('BalancedStrings#Solver', () => {
      BalancedStrings.Solver(input);
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

  public static RuleMatchingBenchmark(input: {
    items: Array<Array<string>>, ruleKey: string, ruleValue: string
  }) {
    this.suite = new Benchmark.Suite();

    this.suite.add('RuleMatching#Solver', () => {
      RuleMatching.Solver(input.items, input.ruleKey, input.ruleValue);
    })
      .add('RuleMatching#FastSolver', () => {
        RuleMatching.FastSolver(input.items, input.ruleKey, input.ruleValue);
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

  public static SentenceSortingBenchmark(input: string): void {
    this.suite = new Benchmark.Suite();

    this.suite.add('SentenceSorting#Solver', () => {
      SentenceSorting.Solver(input);
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

  public static PangramSentenceBenchmark(input: string): void {
    this.suite = new Benchmark.Suite();

    this.suite.add('PangramSentence#Solver', () => {
      PangramSentence.Solver(input);
    })
      .add('PangramSentence#OutOfTheBoxSolver', () => {
        PangramSentence.OutOfTheBoxSolver(input);
      })
      .add('PangramSentence#FastSolver', () => {
        PangramSentence.FastSolver(input);
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

  public static TwoStringArrayBenchmark(input: {
    word1: Array<string>, word2: Array<string>
  }): void {
    this.suite = new Benchmark.Suite();

    this.suite.add('TwoStringArray#Solver', () => {
      TwoStringArray.Solver(input.word1, input.word2);
    })
      .on('cycle', (event: { target: any; }) => {
        log(String(event.target));
      })
      .on('complete', function () {
        log(`The fastest method is ${this.filter('fastest').map('name')}`);
      })
      // run async or not
      .run({ async: false });
  }

  public static StayingInGridBencmark(input: {
    n: number, startPos: Array<number>, s: string
  }): void {
    this.suite = new Benchmark.Suite();

    this.suite.add('StayingInGrid#Solver', () => {
      StayingIngrid.Solver(input.n, input.startPos, input.s);
    })
      .add('StayingInGrid#FastSolver', () => {
        StayingIngrid.FastSolver(input.n, input.startPos, input.s);
      })
      .on('cycle', (event: { target: any; }) => {
        log(String(event.target));
      })
      .on('complete', function () {
        log(`The fastest method is ${this.filter('fastest').map('name')}`);
      })
      // run async or not
      .run({ async: false });
  }

  public static UniqueMorseCodeBencmark(input: string[]): void {
    this.suite = new Benchmark.Suite();

    this.suite.add('UniqueMorseCode#Solver', () => {
      UniqueMorseCode.Solver(input);
    })
      .add('UniqueMorseCode#MapSolver', () => {
        UniqueMorseCode.MapSolver(input);
      })
      .add('UniqueMorseCode#UsingNewMapSolver', () => {
        UniqueMorseCode.UsingNewMapSolver(input);
      })
      .on('cycle', (event: { target: any; }) => {
        log(String(event.target));
      })
      .on('complete', function () {
        log(`The fastest method is ${this.filter('fastest').map('name')}`);
      })
      // run async or not
      .run({ async: false });
  }

  public static MaxLaserBeamBencmark(input: string[]): void {
    this.suite = new Benchmark.Suite();

    this.suite.add('MaxLaserBeam#Solver', () => {
      MaxLaserBeam.Solver(input);
    })
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
