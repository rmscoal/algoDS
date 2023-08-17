/* eslint-disable no-param-reassign */
import { ListNode } from './types.js';

export default class AddTwoNumbers {
  public static Solver(l1?: ListNode, l2?: ListNode): ListNode | null {
    const l3 = new ListNode();
    let tmp = l3;

    while (l1 != null || l2 != null) {
      if (l1 != null) {
        tmp.val += l1.val;
        l1 = l1.next;
      }

      if (l2 != null) {
        tmp.val += l2.val;
        l2 = l2.next;
      }

      if (tmp.val > 9) {
        tmp.val -= 10;
        tmp.next = new ListNode(1);
      } else if (l1 != null || l2 != null) {
        tmp.next = new ListNode();
      }

      tmp = tmp.next;
    }

    return l3;
  }

  public static CleanSolver(l1?: ListNode, l2?: ListNode, carry: number = 0): ListNode | null {
    if (!l1 && !l2 && !carry) return null;

    const total: number = (l1 ? l1.val : 0) + (l2 ? l2.val : 0) + (carry || 0);
    carry = total > 9 ? 1 : 0;

    return new ListNode(total % 10, this.CleanSolver(l1?.next, l2?.next, carry));
  }
}
