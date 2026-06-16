func maxArea(heights []int) int {
	area:=0
	left,right := 0,len(heights)-1
	for left<right {
		tmp:= findMin(heights[left],heights[right]) * (right-left)
		if tmp>=area{
			area = tmp
		}
		if heights[left] >= heights[right] {
			right--
		} else{
			left++
		}
	}
	return area
}

func findMin(a int, b int) int {

	if a>b{
		return b
	}
	return a
}