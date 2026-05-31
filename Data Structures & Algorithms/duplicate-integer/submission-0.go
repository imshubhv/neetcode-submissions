func hasDuplicate(nums []int) bool {
    occurance := make(map[int]int)
    for i := 0; i < len(nums); i++ {
        if occurance[nums[i]] == 1 {
            return true
        } else {
            occurance[nums[i]] = 1
        }
    }
    return false
}
