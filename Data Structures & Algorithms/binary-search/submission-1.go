func search(nums []int, target int) int {
	index:= len(nums) /2
	
	if len(nums) == 0 {
		return -1
	}

	if nums[index] == target {
		return index
	} else if nums[index] > target {
		return search(nums[:index],target)
	} else {
	res := search(nums[index+1:], target)
        if res == -1 {
            return -1
        }
        return index + 1 + res
	}
	return -1
}
