func twoSum(nums []int, target int) []int {
    
    mappOfArray := make(map[int]int)
    for index, value := range nums {
        mappOfArray[value] = index
    }

    for index, value := range nums {
        difference:= target - value

        if mappOfArray[difference] != 0  && mappOfArray[difference] != index {
            return []int {index, mappOfArray[difference]}
        }
    }
    return []int {0,0}
}
