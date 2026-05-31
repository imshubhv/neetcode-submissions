func twoSum(nums []int, target int) []int {
    
    for i:= range nums {
        div := target-nums[i];
        for j:= i+1; j<len(nums); j++{
            if nums[j] == div{
                return []int{i, j}
            }
        }
    }
    return []int {}
}
