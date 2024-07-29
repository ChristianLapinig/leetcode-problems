package besttimetobuyandsellstock

/*
* 121. Best Time to Buy and Sell Stock
* Description: https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
* Editorial: https://leetcode.com/problems/best-time-to-buy-and-sell-stock/editorial
 */
func maxProfit(prices []int) int {
	// the profit will be the difference between prices[i] - the lowest price seen thus far
	maxProfit := 0
	lowestPrice := prices[0] // the price at day 0 is the lowest seen at the start

	// start at day 1 since we have already processed day 0
	for i := 1; i < len(prices); i++ {
		profit := prices[i] - lowestPrice
		maxProfit = max(profit, maxProfit)
		lowestPrice = min(prices[i], lowestPrice)
	}

	// time complexity: O(n) - space complexity: O(1)
	// memory doesn't grow, it's just updated
	return maxProfit
}
