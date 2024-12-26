package main

func CombinationSum(candidates []int, target int) [][]int {
    res := [][]int{}
    
    var backtrack func(path []int, index int, rem int)
    backtrack = func(path []int, index int, rem int) {
        if rem == 0 {
            temp := make([]int, len(path))
            copy(temp, path)
            res = append(res, temp)
            return
        }
        if rem < 0 {
            return
        }
        for i := index; i < len(candidates); i++ {
            path = append(path, candidates[i])
            backtrack(path, i, rem-candidates[i])
            path = path[:len(path)-1]
        }
    }
    
    backtrack([]int{}, 0, target)
    return res
}

// 0ms
// 5.01mb