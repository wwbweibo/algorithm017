func removeDuplicates(nums []int) int {
    i := 1
    j := 0
    for ; i < len(nums); {
        if nums[i] == nums[j]{
            i ++
        } else {
            nums[j + 1] = nums[i]
            j ++
            i ++
        }
    }
    return j  + 1
}