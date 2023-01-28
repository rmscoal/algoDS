/* eslint-disable func-names */
import Benchmark from 'benchmark';
import { log } from 'console';
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
    static MaxProfitBenchmark(input) {
        this.suite = new Benchmark.Suite();
        this.suite.add('MaxProfit#Solver', () => {
            MaxProfit.Solver(input);
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
    static BalancedStringsBenchmark(input) {
        this.suite = new Benchmark.Suite();
        this.suite.add('BalancedStrings#Solver', () => {
            BalancedStrings.Solver(input);
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
    static RuleMatchingBenchmark(input) {
        this.suite = new Benchmark.Suite();
        this.suite.add('RuleMatching#Solver', () => {
            RuleMatching.Solver(input.items, input.ruleKey, input.ruleValue);
        })
            .add('RuleMatching#FastSolver', () => {
            RuleMatching.FastSolver(input.items, input.ruleKey, input.ruleValue);
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
    static SentenceSortingBenchmark(input) {
        this.suite = new Benchmark.Suite();
        this.suite.add('SentenceSorting#Solver', () => {
            SentenceSorting.Solver(input);
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
    static PangramSentenceBenchmark(input) {
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
            .on('cycle', (event) => {
            log(String(event.target));
        })
            .on('complete', function () {
            log(`The fastest method is ${this.filter('fastest').map('name')}`);
        })
            // run async or not
            .run({ async: false });
    }
    static TwoStringArrayBenchmark(input) {
        this.suite = new Benchmark.Suite();
        this.suite.add('TwoStringArray#Solver', () => {
            TwoStringArray.Solver(input.word1, input.word2);
        })
            .on('cycle', (event) => {
            log(String(event.target));
        })
            .on('complete', function () {
            log(`The fastest method is ${this.filter('fastest').map('name')}`);
        })
            // run async or not
            .run({ async: false });
    }
    static StayingInGridBencmark(input) {
        this.suite = new Benchmark.Suite();
        this.suite.add('StayingInGrid#Solver', () => {
            StayingIngrid.Solver(input.n, input.startPos, input.s);
        })
            .add('StayingInGrid#FastSolver', () => {
            StayingIngrid.FastSolver(input.n, input.startPos, input.s);
        })
            .on('cycle', (event) => {
            log(String(event.target));
        })
            .on('complete', function () {
            log(`The fastest method is ${this.filter('fastest').map('name')}`);
        })
            // run async or not
            .run({ async: false });
    }
    static UniqueMorseCodeBencmark(input) {
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