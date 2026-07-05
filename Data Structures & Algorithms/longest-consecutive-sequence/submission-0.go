func longestConsecutive(nums []int) int {
    occur := map[int]bool{}
    for _,value := range nums {
        occur[value] = true
    }
    count:= 0
    for _,v := range nums {
        if occur[v-1] {
            continue
        }
        inc, tmpCount:=1,1
        for occur[v+inc] == true {
            tmpCount++
            inc++
        }
        if tmpCount > count {
            count = tmpCount
        }
        inc = 0
    }
    return count
}
