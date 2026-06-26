func maxProfit(prices []int) int {
	min,max :=prices[0], prices[0]
	answer:=0
	for _,value := range prices {

		if value> max {
			max = value
		} else if value < min {
			min = value
			max = value
		}
		fmt.Println(max, min, value)
		if (max-min > answer){
			answer = max-min
		}
	}
	return answer
}
