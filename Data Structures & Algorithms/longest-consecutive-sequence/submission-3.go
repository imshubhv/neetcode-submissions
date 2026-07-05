func longestConsecutive(nums []int) int {
    occur := make(map[int]bool)

    for _, num := range nums {
        occur[num] = true
    }

    longest := 0

    for num := range occur {
        if occur[num-1] {
            continue
        }

        length := 1
        for occur[num+length] {
            length++
        }

        if length > longest {
            longest = length
        }
    }
    return longest
}
