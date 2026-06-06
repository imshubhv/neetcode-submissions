func hasDuplicate(nums []int) bool {
    frequencyMap:= make(map[int]int)

    for _,value:= range nums {
        if frequencyMap[value] == 1 {
            return true
        } else {
            frequencyMap[value] = frequencyMap[value]+1
        }
    }
    return false
}
