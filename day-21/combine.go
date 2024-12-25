package main

import "fmt"

func Combine(n int, k int) [][]int {
	res := [][]int{}
	var backtrack func(start int, comb []int)
	backtrack = func(start int, comb []int) {
		if len(comb) == k {
			combCopy := make([]int, len(comb))
			copy(combCopy, comb)
			res = append(res, combCopy)
			return
		}
		for i := start; i <= n; i++ {
			comb = append(comb, i)
			backtrack(i+1, comb)
			comb = comb[:len(comb)-1]
		}
	}
	backtrack(1, []int{})
	return res
}

func main() {
	fmt.Println(Combine(4, 2))
}

// 48ms
// 53.93mb