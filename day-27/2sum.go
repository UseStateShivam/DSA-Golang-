package main

func TwoSum(nums []int, target int) []int {
    m := map[int]int{}
    res := []int{}
    for i := 0; i < len(nums); i++ {
        r := target - nums[i]
        ri, e := m[r] 
        if e {
            res = append(res, ri, i) 
        } else {
            m[nums[i]] = i
        }
    }
    return res
}

// 4ms
// 7.2mb