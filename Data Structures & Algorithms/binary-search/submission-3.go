func search(nums []int, target int) int {
	left,right:= 0, len(nums)-1

	for left <= right {
		middle:= (right+left)/2
		val:= nums[middle]
		if val == target{
			return middle
		} else if val < target {
			left = middle+1
		} else {
			right = middle-1
		}
	}
	return -1
}
