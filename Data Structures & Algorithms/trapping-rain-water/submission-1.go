func trap(height []int) int {
    maxLeft := []int{}
	leftMax:= 0
	maxRight := make([]int, len(height))
	rightMax:= 0
	for _,val:= range height{
		maxLeft = append(maxLeft,leftMax)
		if val >= leftMax {
			leftMax = val	
		}
	}
	for j:= len(height)-1; j>=0 ; j--{
		maxRight[j] = rightMax
		if height[j] >= rightMax {
			rightMax = height[j]	
		}
	}
	fmt.Println(maxRight,maxLeft)
	water:=0
	for i,val:= range height{
		//maxRight,maxLeft
		tmp:= returnMin(maxRight[i],maxLeft[i])-val
		if(tmp>0){
			water += tmp
		}

	}
	return water
}

func returnMin(a int, b int) int {
	if a>=b {return b }
	return a
}
