func topKFrequent(nums []int, k int) []int {
    counts := make(map[int]int)
    bucket := make([][]int, len(nums) + 1)

    for _, num := range nums {
        counts[num]++
    }

    for key, val := range counts {
        bucket[val] = append(bucket[val], key)
    }

    var result []int
    for i := len(bucket) - 1; i > 0; i-- {
        result = append(result, bucket[i]...)
        if len(result) == k {
            return result
        }
    }

    return result

}
