func twoSum(numbers []int, target int) []int {
	for i:= range numbers {
		diff:= target-numbers[i]
		for j:=i+1; j<len(numbers) ; j++{
			if diff == numbers[j]{
				return []int{i+1,j+1}
			}
		}
	}
	return []int{-1,-1}
}
