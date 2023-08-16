/* eslint-disable max-classes-per-file */
import { assert } from 'chai';
import IsInstance from '../src/leetcode/isinstance.js';

describe('IsInstance Test', () => {
  describe('IsInstance#Solver', () => {
    it('Returns true', () => {
      class Animal {}
      const animal = new Animal();
      assert.equal(IsInstance.Solver(animal, Animal), true, 'Should be true');
    });

    it('Returns true', () => {
      class Animal {}
      class Mammal extends Animal {}
      const mammal = new Mammal();
      assert.equal(IsInstance.Solver(mammal, Animal), true, 'Should be true');
    });

    it('Returns true', () => {
      assert.equal(IsInstance.Solver(Object, Object), true, 'Should be true');
    });

    it('Returns true', () => {
      assert.equal(IsInstance.Solver(5n, BigInt), true, 'Should be true');
    });

    it('Returns true', () => {
      assert.equal(IsInstance.Solver(5n, Object), true, 'Should be true');
    });

    it('Returns true', () => {
      assert.equal(IsInstance.Solver(5, Number), true, 'Should be true');
    });

    it('Returns false', () => {
      assert.equal(IsInstance.Solver(Date, Date), false, 'Should be true');
    });

    it('Returns false', () => {
      assert.equal(IsInstance.Solver(5, String), false, 'Should be true');
    });
  });
});
