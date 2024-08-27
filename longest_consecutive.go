func getMax(a, b int) int {
    if a > b {
        return a
    }
    return b
}


func longestConsecutive(nums []int) int {
    numsHash := make(map[int]bool)
    longest := 0
    count := 0
    var exists bool

    for _, num := range nums {
        numsHash[num] = true
    }

    for _, num := range nums {
        count = 1

        if _, exists := numsHash[num - 1]; exists {
            continue
        }

        _, exists = numsHash[num + 1]
        for exists {
            count++
            _, exists = numsHash[num + count]

        }

        longest = getMax(count, longest)
    }

    return longest
}
