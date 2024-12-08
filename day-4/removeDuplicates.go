package main

// brute force
func RemoveDuplicates(nums []int) int {
    mapOfSeenNumbers := make(map[int]struct{})
    // [0,0,1,1,1,2,2,3,3,4]
    for i := 0; i < len(nums); i++ { // 2
        if _, duplicateExists := mapOfSeenNumbers[nums[i]]; !duplicateExists {
            // nums[2] = 2, exist 
            mapOfSeenNumbers[nums[i]] = struct{}{}
            // {0, 1}
        } else {
            nums = append(nums[:i], nums[i+1:]...)
            // nums = nums[:2] + nums[3:]
            // nums = [0,1,2,2,3,3,4]
            // i = 2, 
            i--
            // i = 1
        }
    }
    return len(nums)
}

// To optimize:
// func removeDuplicates(nums []int) int {
//     k := 1
//     for i := 1; i < len(nums); i++ {
//         if nums[i-1] != nums[i] {
//             nums[k] = nums[i]
//             k++
//         }
//     }
//     return k
// }