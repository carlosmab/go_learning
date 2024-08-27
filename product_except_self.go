func productExceptSelf(nums []int) []int {
    result := []int{1}
    post := 1
    N := len(nums)

    for i := 0; i < len(nums) - 1; i++ {
        result = append(result, result[i] * nums[i])
    }

    for i := 0; i < len(nums); i++ {
        result[N - i - 1] *= post
        post *= nums[N - i - 1] 
    }

    return result
}
