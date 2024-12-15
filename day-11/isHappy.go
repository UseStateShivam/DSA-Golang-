package main

func IsHappy(n int) bool {
    rem := 0
    sumOfSqrOfDigits := 0
    mapOfSeenTotals := make(map[int]struct{})
    mapOfSeenTotals[n] = struct{}{}
    for sumOfSqrOfDigits != 1 {
        for n != 0 { 
            rem = n % 10 
            sumOfSqrOfDigits += rem * rem 
            n /= 10 
        }
        if sumOfSqrOfDigits == 1 {
            return true
        } else if _, exists := mapOfSeenTotals[sumOfSqrOfDigits]; exists {
            return false 
        } else {
            mapOfSeenTotals[sumOfSqrOfDigits] = struct{}{}
        }
        n = sumOfSqrOfDigits
        sumOfSqrOfDigits = 0
    }
    return false
}

// 0ms
// 4.07mb