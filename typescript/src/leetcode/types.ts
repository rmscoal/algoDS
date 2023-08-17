/* eslint-disable no-param-reassign */
/* eslint-disable import/prefer-default-export */
export class ListNode {
  val: number;

  next: ListNode | null;

  constructor(val?: number, next?: ListNode | null) {
    this.val = val ?? 0;
    this.next = next ?? null;
  }

  public static newFromArray(arr: number[]): ListNode {
    const node = new ListNode();
    let tmp = node;

    for (let i = 0; i < arr.length; i += 1) {
      if (i !== 0) {
        tmp.next = new ListNode();
        tmp = tmp.next;
      }
      tmp.val = arr[i];
    }

    return node;
  }

  public static equal(l1: ListNode | null, l2: ListNode | null): boolean {
    if (!l1 && !l2) return true;
    if (!l1 && l2) return false;

    while (l1 || l2) {
      if (!l2 && l1) return false;
      if (!l1 && l2) return false;

      if (l1.val !== l2.val) return false;

      l1 = l1.next;
      l2 = l2.next;
    }

    return true;
  }
}
