func carFleet(target int, position []int, speed []int) int {
	length := len(position)
	timeRem:= make([]float64,length)

	for index,value:= range position {
		timeRem[index] = float64(target - value)/float64(speed[index])
	}
	indexArray:= make([]int,length)
	for i:= range indexArray{
		indexArray[i] = i
	}
	sort.Slice(indexArray, func(i,j int)bool {
		return position[indexArray[i]] > position[indexArray[j]]
	})
	fleet, prev:= 0, 0.0
	for _,v := range indexArray {
		if timeRem[v] > prev {
			prev = timeRem[v]
			fleet++
		}
	}
	return fleet

}
