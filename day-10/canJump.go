package main

func max(a int, b int) int {
    if a > b {
        return a
    } 
    return b
}

func CanJump(nums []int) bool {
    // start at index 0
    // number of jumps can be <= nums[i]
    // if somehow pointer reaches len(nums) - 1 return true
    // else return false

    // to check if nums[i] has the capacity to jump directly to the last
    // we can check if i + nums[i] == len(nums) - 1 
    
    // for i := 0; i < len(nums); i++ { 
    //     i = i + nums[i]
    //     if i == len(nums) - 1 {
    //         return true
    //     }
    // }

    // return false

    // works but i am concerned for the edge case where [2,3,1,1,4] 
    // say rather than 1 at index 2, was 0 or anything else
    // but still we can jump 1 at index 0 then jump 3 
    // or simply alternate paths

    // so lets implement this using 
    // if the farthest we can jump from a given emelent >= last index
    // this means we can simply reach the last index

    farthest := 0

    for i := 0; i < len(nums); i++ {
        if i > farthest {
            return false // if the farthest we can go is less than the curr index, we simply cant reach the end
        }
        farthest = max(farthest, i + nums[i]) // farthest is the max among prev farthest reachable and the sum of the cuu index and the element itself that states how long can we jump
        if farthest >= len(nums) - 1 {
            return true // if the farthest we can move is >= last index, we can infact touch the last index
        }
    }
    return false
}   

// 0ms
// 8.92mb