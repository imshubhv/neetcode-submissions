func maxProfit(prices []int) int {
	l := 0
	var maxSum int

	for r := 1; r < len(prices); r++ {
		if prices[r] < prices[l] {
			l = r
		}

		profit := prices[r] - prices[l]
		if profit > maxSum {
			maxSum = profit
		}

	}

	return maxSum
}
