func topKFrequent(nums []int, k int) []int {
    freqMap := make(map[int]int)

    for _,value := range nums {
        freqMap[value]++
    }

    bucketList := make([][]int, len(nums)+1)

    for value, freq := range freqMap {
        bucketList[freq] = append(bucketList[freq], value)
    }
    var result []int
    for i:= len(bucketList)-1; i>=0 && len(result)< k; i-- {
        result = append(result,bucketList[i]...)
    } 
    return result
}
