import { assert } from 'chai';
import PangramSentence from '../src/leetcode/pangramsentence.js';
describe('Pangram Sentence Test', () => {
    describe('PangramSentence#Solver', () => {
        describe('Test Case 1', () => {
            it('Should return true', () => {
                const res = PangramSentence.Solver('thequickbrownfoxjumpsoverthelazydog');
                assert.equal(res, true);
            });
        });
        describe('Test Case 2', () => {
            it('Should return false', () => {
                const res = PangramSentence.Solver('leetcode');
                assert.equal(res, false);
            });
        });
        describe('Test Case 3', () => {
            it('Should return true', () => {
                const res = PangramSentence.Solver('abcdefghijklmnopqrstuvwxyz');
                assert.equal(res, true);
            });
        });
    });
});
//# sourceMappingURL=pangramsentence.spec.js.map