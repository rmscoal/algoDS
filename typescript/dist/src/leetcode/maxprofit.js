/**
 * A class to solve max profit.
 * Problem:
 * You are given an array prices where prices[i]
 * is the price of a given stock on the ith day.
 * You want to maximize your profit by choosing
 * a single day to buy one stock and choosing a
 * different day in the future to sell that stock.
 * Return the maximum profit you can achieve from
 * this transaction. If you cannot achieve any
 * profit, return 0.
 */
export default class MaxProfit {
    /**
     * @public @static
     * @param input defines the input of the stock prices
     * each day
     * @returns {number} maxProfit: number
     *
     * Time-complexity: O(n)
     * Space-complexity: O(1)
     */
    static Solver(input) {
        // Returned variable:
        let maxProfit = 0;
        // A variable to hold the value
        // of a bought price.
        let buyPrice = input[0];
        // Loop through the input and finds
        // the highest profit it could get.
        for (let i = 1; i < input.length; i += 1) {
            const sellPrice = input[i];
            if (maxProfit < sellPrice - buyPrice) {
                maxProfit = sellPrice - buyPrice;
            }
            if (sellPrice < buyPrice) {
                buyPrice = sellPrice;
            }
        }
        return maxProfit;
    }
}
//# sourceMappingURL=maxprofit.js.map