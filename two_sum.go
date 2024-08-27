import "fmt"

func twoSum(nums []int, target int) []int {
    complements := make(map[int]int)

    for index, num := range nums {
        
        complement_index, exists := complements[num]

        if exists {
            return []int{index, complement_index}
        }

        complements[target - num] = index
    }

    return []int{}
}
