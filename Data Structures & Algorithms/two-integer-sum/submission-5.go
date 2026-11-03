func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	for index,val := range nums{
		numMap[val] = index
	}
	for index, value := range nums{
		difference := target-value
		position, ok := numMap[difference]
		if ok && position != index {
			return []int {index, position}
		}
	}
	return []int {}
}
