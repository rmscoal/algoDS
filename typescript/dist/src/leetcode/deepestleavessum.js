/* eslint-disable max-classes-per-file */
export class TreeNode {
    constructor(val, left, right) {
        this.val = (val === undefined ? 0 : val);
        this.left = (left === undefined ? null : left);
        this.right = (right === undefined ? null : right);
    }
}
export class DeepestLeavesSum {
    static Solver(root) {
        if (root === null) {
            return 0;
        }
        const maxDepth = this.findMaxDepth(root, 1);
        return this.helper(root, 1, maxDepth);
    }
    static helper(node, curr, depth) {
        if (node === null) {
            return 0;
        }
        if (node.left === null && node.right === null) {
            if (curr === depth) {
                return node.val;
            }
            return 0;
        }
        if (node.left === null) {
            return this.helper(node.right, curr + 1, depth);
        }
        if (node.right === null) {
            return this.helper(node.left, curr + 1, depth);
        }
        return this.helper(node.left, curr + 1, depth) + this.helper(node.right, curr + 1, depth);
    }
    static findMaxDepth(node, curr) {
        if (node === null) {
            return curr;
        }
        if (node.left === null && node.right === null) {
            return curr;
        }
        if (node.left === null) {
            return this.findMaxDepth(node.right, curr + 1);
        }
        if (node.right === null) {
            return this.findMaxDepth(node.left, curr + 1);
        }
        const left = this.findMaxDepth(node.left, curr + 1);
        const right = this.findMaxDepth(node.right, curr + 1);
        if (left > right) {
            return left;
        }
        return right;
    }
}
//# sourceMappingURL=deepestleavessum.js.map