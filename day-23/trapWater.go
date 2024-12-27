package main

func Trap(height []int) int {
	leftArr := []int{}
	rightArr := []int{}
	leftMax := 0
	rightMax := 0
	res := 0
	for i := 0; i < len(height); i++ {
		leftMax = max(leftMax, height[i])
		if i == 0 {
			leftArr = append(leftArr, 0)
		} else {
			leftArr = append(leftArr, leftMax)
		}
	}
	for i := len(height) - 1; i >= 0; i-- {
		rightMax = max(rightMax, height[i])
		if i == len(height)-1 {
			rightArr = append(rightArr, 0)
		} else {
			rightArr = append(rightArr, rightMax)
		}
	}
	for i := 0; i < len(height); i++ {
		if (min(leftArr[i], rightArr[len(height)-1-i]) - height[i]) >= 0 {
			res += (min(leftArr[i], rightArr[len(height)-1-i]) - height[i])
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 5ms
// 8.31mb