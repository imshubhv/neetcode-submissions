func trap(height []int) int {
	water:=0
	leftMax, rightMax :=0,0
	left, right := 0, len(height)-1
	for left<=right {
		// left to right
		if leftMax <= rightMax {
			hl:= height[left]
			tmp:= leftMax-hl
			if(tmp>0){
				water += tmp
			}
			//update max element -left
			if (leftMax < hl){
				leftMax = hl
			}
			left++
		}else { //right to left
			hr:= height[right]
			tmp:= rightMax-hr
			if(tmp>0){
				water += tmp
			}
			//update max element -right
			if (rightMax < hr){
				rightMax = hr
			}
			right--
		}
	}
	return water
}
