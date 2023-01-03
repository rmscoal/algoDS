import { assert } from 'chai';
import Decibinary from '../src/leetcode/decibinary.js';

describe('Decibinary Test', () => {
  describe('Decibinary#Solver', () => {
    describe('Testcase One', () => {
      it('Returns 3', () => {
        const dec: Decibinary = new Decibinary('123');
        assert.equal(dec.Solver(), 3, 'Should returns 3');
      });
    });

    describe('Testcase Two', () => {
      it('Returns 9', () => {
        const dec: Decibinary = new Decibinary('27346209830709182346');
        assert.equal(dec.Solver(), 9, 'Should returns 9');
      });
    });

    describe('Testcase Three', () => {
      it('Returns 1', () => {
        const dec: Decibinary = new Decibinary('111111111111111111111');
        assert.equal(dec.Solver(), 1, 'Should returns 1');
      });
    });
  });

  describe('Decibinary#FastSolver', () => {
    describe('Testcase One', () => {
      it('Returns 3', () => {
        const dec: Decibinary = new Decibinary('123');
        assert.equal(dec.FastSolver(), 3, 'Should returns 3');
      });
    });

    describe('Testcase Two', () => {
      it('Returns 9', () => {
        const dec: Decibinary = new Decibinary('27346209830709182346');
        assert.equal(dec.FastSolver(), 9, 'Should returns 9');
      });
    });

    describe('Testcase Three', () => {
      it('Returns 1', () => {
        const dec: Decibinary = new Decibinary('111111111111111111111');
        assert.equal(dec.FastSolver(), 1, 'Should returns 1');
      });
    });
  });
});
