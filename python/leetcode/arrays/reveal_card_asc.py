"""
You are given an integer array deck. There is a deck of cards where every card 
has a unique integer. The integer on the ith card is deck[i].

You can order the deck in any order you want. Initially, all the cards start 
face down (unrevealed) in one deck.

You will do the following steps repeatedly until all cards are revealed:

Take the top card of the deck, reveal it, and take it out of the deck.
If there are still cards in the deck then put the next top card of the deck at 
the bottom of the deck.
If there are still unrevealed cards, go back to step 1. Otherwise, stop.
Return an ordering of the deck that would reveal the cards in increasing order.

Note that the first entry in the answer is considered to be the top of the deck.

https://leetcode.com/problems/reveal-cards-in-increasing-order/description/
"""

"""
Input: deck = [17,13,11,2,3,5,7]

Sorted Input = [2,3,5,7,11,13,17]

Iteration 1:
    - Fill in 2 largest elements in arr
    -> Sorted =  [2,3,5,7,11]
    -> Result = [13, 17]
Iteration 2:
    We want to make 11 at the start
    and the last to the be second. 
    Shift all to +2 index
    -> Sorted = [2, 3, 5, 7]
    -> Result = [11, 17, 13]
Iteration 3:
    We want to make 7 at the start
    and the last to the be second. 
    Shift all to +2 index
    -> Sorted = [2, 3, 5]
    -> Result = [7, 13, 11, 17]
Iteration 4:
    We want to make 5 at the start
    and the last to the be second. 
    Shift all to +2 index
    -> Sorted = [2, 3]
    -> Result = [5, 17, 7, 13, 11]
Iteration 5:
    We want to make 3 at the start
    and the last to the be second. 
    Shift all to +2 index
    -> Sorted = [2]
    -> Result = [3, 11, 5, 17, 7, 13]
Iteration 6:
    We want to make 2 at the start
    and the last to the be second. 
    Shift all to +2 index
    -> Sorted = []
    -> Result = [2, 13, 3, 11, 5, 17, 7]
"""




from typing import List
class Solution:
    def deck_revealed_increasing(self, deck: List[int]) -> List[int]:
        # Sort the deck in ascending order
        deck.sort()

        if len(deck) <= 2:
            return deck

        result = []

        result.insert(0, deck.pop(len(deck) - 1))
        result.insert(0, deck.pop(len(deck) - 1))

        while len(deck) != 0:
            last = result.pop()
            result = [deck.pop()] + [last] + result

        return result


if __name__ == "__main__":
    print("Result", Solution().deck_revealed_increasing(
        [17, 13, 11, 2, 3, 5, 7]))

    print("Result", Solution().deck_revealed_increasing(
        [17, 42, 8, 91, 56, 33, 19, 72, 24, 10, 5, 68, 51]))

    print("Result", Solution().deck_revealed_increasing(
        [1, 1000]))

    print("Result", Solution().deck_revealed_increasing(
        [100000000000, 1000]))
