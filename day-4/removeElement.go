package main

func RemoveElement(nums []int, val int) int {
    // output := 0
    // [3,2,2,3] , 3
    for i := 0; i < len(nums); i++ { // i = 3
        if nums[i] == val { // nums[2] = 2, val = 3, 
            // output++
            nums = append(nums[:i], nums[i+1:]...)
            i--
        } 
    }
    // output = 2
    return len(nums)
}

// 0ms
// 4mb