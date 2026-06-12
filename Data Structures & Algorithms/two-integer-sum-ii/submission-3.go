func twoSum(numbers []int, target int) []int {
    seen := make(map[int]int) // value -> index

    for i, num := range numbers {
        diff := target - num

        if idx, found := seen[diff]; found {
            return []int{idx + 1, i + 1}
        }

        seen[num] = i
    }

    return []int{-1, -1}
}
