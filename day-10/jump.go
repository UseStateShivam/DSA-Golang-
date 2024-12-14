package main

// func max(a int, b int) int {
//     if a > b {
//         return a
//     }
//     return b
// }

func Jump(nums []int) int {
    // start from index 0
    // same as prev, but now we need to select the path with least jumps
    // but since the return type is int, its sure that all slices we get, its possible to reach the end somehow
    // if some big number is there somewhere in the middle, and i + nums[i] == len(nums) - 1, then we may try to reach that i asap
    // we may explore two pointers
    // if the right pointer i + nums[i] == len(nums) - 1 then we just need to make out left reach right 
    // if right does not meet the reqd then we push right to farthest we can go from the left + 1

    if len(nums) < 2 {
        return 0
    }

    currEnd := 0
    farthest := 0
    jump := 0

    for i := 0; i < len(nums); i++ {
        farthest = max(farthest, i + nums[i])

        if i == currEnd {
            jump++
            currEnd = farthest
            if currEnd >= len(nums) - 1 {
                break
            } 
        }
    }

    return jump
}

// 0ms
// 7.84mb