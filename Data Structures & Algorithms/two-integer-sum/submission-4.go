func twoSum(nums []int, target int) []int {
    for index,value := range nums {
        difference := target-value
        for i:= index+1;i< len(nums); i++ {
            if nums[i] == difference {
                return []int{index, i}
            }
        }
    }
    return []int{-1,-1}
}
