package main

func ContainsNearbyDuplicate(nums []int, k int) bool {
    m := map[int]int{}
    for i := 0; i < len(nums); i++ {
        index, exists := m[nums[i]] 
        if exists {
            if mod(i, index) <= k {
                return true
            } else {
                m[nums[i]] = i
            }
        } else {
            m[nums[i]] = i
        }
    }
    return false
}

func mod(a, b int) int {
    if a - b < 0 {
        return b - a
    }
    return a - b
}

// 24ms
// 14.51mb