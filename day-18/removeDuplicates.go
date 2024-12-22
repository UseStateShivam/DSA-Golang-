package main

func RemoveDuplicates(nums []int) int {
    if len(nums) == 0 {
        return 0
    }

    l := 0
    for r := 1; r < len(nums); r++ {
        if nums[l] != nums[r] {
            l++
            nums[l] = nums[r]
        }
    }
    return l+1
}

// 0ms
// 6.30mb