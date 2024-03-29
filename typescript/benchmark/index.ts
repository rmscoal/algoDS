import CustomBenchmark from './benchmark.js';

CustomBenchmark.DecibinaryBenchmark('27346209830709182346');
CustomBenchmark.TwoSumBenchmark([1, 2, 3, 4, 5, 6, 7, 8, 2, 3, 1, 4, 5], 6);
CustomBenchmark.TwoSumBenchmark([1, 2, 3, 4, 5, 6, 7, 8, 2, 3, 1, 4, 5], 13);
CustomBenchmark.CollectGarbageBenchmark(['MMM', 'PGM', 'GP'], [3, 10]);
CustomBenchmark.MaxProfitBenchmark([1, 2, 3, 4, 1]);
CustomBenchmark.BalancedStringsBenchmark('RLRRRLLLLRRLLRRLLLRRRRRLLL');
CustomBenchmark.RuleMatchingBenchmark({
  items: [['phone', 'blue', 'pixel'], ['computer', 'silver', 'lenovo'], ['phone', 'gold', 'iphone']],
  ruleKey: 'color',
  ruleValue: 'silver',
});
CustomBenchmark.SentenceSortingBenchmark('is2 sentence4 This1 a3');
CustomBenchmark.PangramSentenceBenchmark('thequickbrownfoxjumpsoverthelazydog');
CustomBenchmark.TwoStringArrayBenchmark({
  word1: ['abc', 'd', 'defg'],
  word2: ['abcddefg'],
});
CustomBenchmark.StayingInGridBencmark({
  n: 100,
  startPos: [60, 52],
  s: 'LLLDURDLRUURRLLDLDLRUUDDDLRLRUUDDLRRLRRLDUDLLLRDLRUDL',
});
CustomBenchmark.UniqueMorseCodeBencmark(['gin', 'zen', 'gig', 'msg']);
CustomBenchmark.MaxLaserBeamBencmark(['011001', '000000', '010100', '001000', '011001', '000000', '010100', '001000', '011001', '000000', '010100', '001000']);
