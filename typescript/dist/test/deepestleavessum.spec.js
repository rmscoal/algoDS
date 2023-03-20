import { assert } from 'chai';
import { DeepestLeavesSum, TreeNode } from '../src/leetcode/deepestleavessum.js';
describe('Deepest Leaves Sum', () => {
    describe('DeepestLeavesSum#Solver', () => {
        const tree1leave1 = new TreeNode(1, null, null);
        const tree1leave2 = new TreeNode(3, null, null);
        const tree1 = new TreeNode(2, tree1leave1, tree1leave2);
        const tree2leave1 = new TreeNode(7, null, null);
        const tree2leave2 = new TreeNode(8, null, null);
        const tree2node5 = new TreeNode(6, null, tree2leave2);
        const tree2node4 = new TreeNode(5, null, null);
        const tree2node3 = new TreeNode(4, tree2leave1, null);
        const tree2node2 = new TreeNode(3, null, tree2node5);
        const tree2node1 = new TreeNode(2, tree2node3, tree2node4);
        const tree2 = new TreeNode(1, tree2node2, tree2node1);
        describe('Testcase one', () => {
            it('should return one', () => {
                assert.equal(DeepestLeavesSum.Solver(tree1), 4);
            });
        });
        describe('Testcase two', () => {
            it('should return 15', () => {
                assert.equal(DeepestLeavesSum.Solver(tree2), 15);
            });
        });
    });
});
//# sourceMappingURL=deepestleavessum.spec.js.map