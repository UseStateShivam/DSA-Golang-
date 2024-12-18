package main

func TwoSum(numbers []int, target int) []int {
    // given a sorted array
    // hehe, its basically the famous 2 sum problem
    
    // make a map for seen chars as the key and their index as value
    // mapOfSeenDigits := make(map[int]int)

    // // output to be returned
    // output := []int{}

    // for i := 0; i < len(numbers); i++ {
    //     if index, exists := mapOfSeenDigits[(target - numbers[i])]; exists {
    //         output = append(output, index+1)
    //         output = append(output, i+1)
    //     } else {
    //         mapOfSeenDigits[numbers[i]] = i
    //     }
    // }

    // return output

    // although the problem can be solved with this
    // to further improve uppon the space complexity
    // we can take advantage of the fact that the input is sorted
    l, r := 0, len(numbers) - 1

    for l < r {
        sum := numbers[l] + numbers[r]
        if sum == target {
            return []int{l+1, r+1}
        } else if sum < target {
            l++
        } else {
            r--
        }
    }

    return []int{}
}

// 0ms
// 7.96mb
// earlier we used map that has a growing space which expands for each index
// but we optimized this to a 2 pointer apporach where we attained a constant space complexity