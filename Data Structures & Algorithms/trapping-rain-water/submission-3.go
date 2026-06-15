// func trap(height []int) int {
//     maxLeft := []int{}
// 	leftMax:= 0
// 	maxRight := make([]int, len(height))
// 	rightMax:= 0
// 	water:=0
// 	for i,val:= range height{
// 		j:= len(height)-1- i
// 		maxLeft = append(maxLeft,leftMax)
// 		maxRight[j] = rightMax
// 		if val >= leftMax {
// 			leftMax = val	
// 		}
// 		if height[j] >= rightMax {
// 			rightMax = height[j]	
// 		}
// 	}
// 	for i,val:= range height{
// 		//maxRight,maxLeft
// 		tmp:= returnMin(maxRight[i],maxLeft[i])-val
// 		fmt.Println(tmp)
// 		if(tmp>0){
// 			water += tmp
// 		}
// 	}
// 	return water
// }

func returnMin(a int, b int) int {
	if a>=b {return b }
	return a
}

func trap(height []int) int {
	water:=0
	leftMax, rightMax :=0,0
	left, right := 0, len(height)-1
	for left<=right {
		// left to right
		if leftMax <= rightMax {
			tmp:= returnMin(rightMax,leftMax)-height[left]
			if(tmp>0){
				water += tmp
			}
			fmt.Println(tmp)
			//update max element -left
			if (leftMax < height[left]){
				leftMax = height[left]
			}
			left++
		}else { //right to left
			tmp:= returnMin(rightMax,leftMax)-height[right]
			if(tmp>0){
				water += tmp
			}
			fmt.Println(tmp)
			//update max element -right
			if (rightMax < height[right]){
				rightMax = height[right]
			}
			right--
		}
	}
	return water
}
