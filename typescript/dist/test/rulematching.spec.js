import { assert } from 'chai';
import RuleMatching from '../src/leetcode/rulematching.js';
describe('Rule Matching Test', () => {
    describe('RuleMatching#Solver', () => {
        describe('Test Case 1', () => {
            it('Should return 1', () => {
                const res = RuleMatching.Solver([['phone', 'blue', 'pixel'], ['computer', 'silver', 'lenovo'], ['phone', 'gold', 'iphone']], 'color', 'silver');
                assert.equal(res, 1, 'Result should be 1');
            });
        });
        describe('Test Case 2', () => {
            it('Should return 2', () => {
                const res = RuleMatching.Solver([['phone', 'blue', 'pixel'], ['computer', 'silver', 'phone'], ['phone', 'gold', 'iphone']], 'type', 'phone');
                assert.equal(res, 2, 'Result should be 2');
            });
        });
    });
    describe('RuleMatching#FastSolver', () => {
        describe('Test Case 1', () => {
            it('Should return 1', () => {
                const res = RuleMatching.FastSolver([['phone', 'blue', 'pixel'], ['computer', 'silver', 'lenovo'], ['phone', 'gold', 'iphone']], 'color', 'silver');
                assert.equal(res, 1, 'Result should be 1');
            });
        });
        describe('Test Case 2', () => {
            it('Should return 2', () => {
                const res = RuleMatching.FastSolver([['phone', 'blue', 'pixel'], ['computer', 'silver', 'phone'], ['phone', 'gold', 'iphone']], 'type', 'phone');
                assert.equal(res, 2, 'Result should be 2');
            });
        });
    });
});
//# sourceMappingURL=rulematching.spec.js.map