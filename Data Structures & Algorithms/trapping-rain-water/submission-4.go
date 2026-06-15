func trap(height []int) int {
	water:=0
	leftMax, rightMax :=0,0
	left, right := 0, len(height)-1
	for left<=right {
		// left to right
		if leftMax <= rightMax {
			tmp:= leftMax-height[left]
			if(tmp>0){
				water += tmp
			}
			//update max element -left
			if (leftMax < height[left]){
				leftMax = height[left]
			}
			left++
		}else { //right to left
			tmp:= rightMax-height[right]
			if(tmp>0){
				water += tmp
			}
			//update max element -right
			if (rightMax < height[right]){
				rightMax = height[right]
			}
			right--
		}
	}
	return water
}
