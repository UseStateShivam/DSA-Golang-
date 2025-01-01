package main

func Rotate(nums []int, k int) {
    if len(nums) <= 1 {
        return
    }
    k = k % len(nums)
	n1 := nums[:len(nums)-k]
	n2 := nums[len(nums)-k:]
	reverse(n1)
	reverse(n2)
	nums = append(n1, n2...)
	reverse(nums)
}

func reverse(n []int) {
	a := make([]int, len(n))
	for i := len(n)-1; i >= 0; i-- {
		a[len(a)-i-1] = n[i]
	}
	copy(n, a)
}

// 4ms
// 11.60mb