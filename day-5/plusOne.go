package main

func AllDigitsAreNine(digits []int) bool {
    for _, digit := range digits {
        if digit != 9 {
            return false
        }
    }
    return true
}

func PlusOne(digits []int) []int {
    // if after adding one to last index of the slice leads to number == 10
    // we have to implement carry over
    if AllDigitsAreNine(digits) {
        digits = append([]int{1}, make([]int, len(digits))...)
        return digits
    }
    carry := 1
    for i := len(digits) - 1; i >= 0 && carry > 0; i-- {
        digits[i] += carry
        if digits[i] == 10 {
            digits[i] = 0
            carry = 1
        } else {
            carry = 0
        }

    }
    return digits 
}

// 0ms
// 3.92mb