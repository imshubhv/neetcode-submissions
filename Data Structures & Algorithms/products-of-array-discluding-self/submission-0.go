func productExceptSelf(nums []int) []int {
	hasZero, hasTwoZeros := false, false
	var product int = 1
	output := make( []int, len(nums))
	for _, value := range nums{
		if value == 0 {
			if hasZero == true {
				hasTwoZeros = true
				break
			}
			hasZero = true
		} else {
			product = product * value
		}
	}
	if hasTwoZeros == true {
		return output
	}else if hasZero == false{
		for index, value:= range nums {
			output[index] = product / value
		}
	} else {
		for index, value:= range nums {
			if (value == 0) {
				output[index] = product
			} else {
				output[index] = 0
			}
		}
	}	
	return output
}
