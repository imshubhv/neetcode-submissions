func hasDuplicate(nums []int) bool {
    maps := make(map[int]int);
    for i:=0; i<len(nums); i++ {
        if maps[nums[i]] == 0 {
            maps[nums[i]] = 1;
        }else {
            return true;
        }
    }
    return false;
}
