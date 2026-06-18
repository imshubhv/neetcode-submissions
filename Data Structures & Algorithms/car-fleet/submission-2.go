func carFleet(target int, position []int, speed []int) int {
	size:= len(position)
	timeRemaining:= make([]float64, size)
	for i,v:= range position{ 
		timeRemaining[i] = float64(target-v) / float64(speed[i])
	}
	indexArray:= make([]int, size)
	for i:= range indexArray{
		indexArray[i] = i
	}
	sort.Slice(indexArray, func(i,j int) bool {
		return position[indexArray[i]] > position[indexArray[j]]
	})
	prev,fleet:= -1.0,0
	for _,value := range indexArray {
		if timeRemaining[value]> prev{
			prev = timeRemaining[value]
			fleet++
		}
	}
	fmt.Println(indexArray)
	return fleet
}
func divison(a int, b int) int {
	return (a+b-1)/b;
}
