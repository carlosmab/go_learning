func containsDuplicate(nums []int) bool {
    hash_map := make(map[int]int)    

    for _, num := range nums {
        _, num_exists := hash_map[num]

        if num_exists {
            return true
        }

        hash_map[num] = 10 
    }

    return false
}
