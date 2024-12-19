package main

func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	} else {
		actualNum := x
		reversedNum := 0
		// 121
		// without converting to string
		// might make a for loop take out individual digits from the given number...
		for x != 0 {
			rem := x % 10
			reversedNum = (reversedNum * 10) + rem
			x /= 10
		}
		if reversedNum == actualNum {
			return true
		} else {
			return false
		}
	}
}

// 0ms
// 6.22mb