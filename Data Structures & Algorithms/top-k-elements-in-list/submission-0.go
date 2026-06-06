func topKFrequent(nums []int, k int) []int {
    frequencyMap:= make(map[int]int)

    for _, value:= range nums {
        frequencyMap[value]++
    }

    buckets:= make([][]int, len(nums)+1)
    for number, freq:= range frequencyMap {
        buckets[freq] = append(buckets[freq],number)
    }

    
    result := []int{}
    for i := len(buckets) - 1; i >= 0 && len(result) < k; i-- {
        result = append(result, buckets[i]...)
    }


    fmt.Println(result)
    return result[:k]
}
