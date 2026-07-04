func productExceptSelf(nums []int) []int {
	zeros:=0
	var product int = 1
	output := make( []int, len(nums))
	for _, value := range nums{
		if value == 0 {
			zeros++
		} else {
			product = product * value
		}
	}
	if zeros >= 2{
		return output
	}else if zeros == 0{
		for index, value:= range nums {
			output[index] = product / value
		}
	} else {
		for index, value:= range nums {
			if (value == 0) {
				output[index] = product
				break
			}
		}
	}	
	return output
}
