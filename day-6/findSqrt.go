package main

func MySqrt(x int) int {
    // given: int
    // return: sqrt of given int

    // approach
    // say x = 50
    // sqrt of 50 is 7.something, we return 7
    // first thing that comes to my mind, is simply loop around until the sqr of curr i > x
    // return i
    // but this may mess around with the time complexity 

    // for i := 0; i <= x; i++ {
    //     if i * i > x {
    //         return i - 1
    //     } else if i * i == x {
    //         return i
    //     }
    // }
    // return 0
    // works but can be optimized further

    // lets explore binary search
    left := 0
    right := x
    for left <= right {
        mid := left + ((right - left) / 2)
        // say left = 3, right = 3
        // mid = 2 
        // if mid^2 > x, search in lhs, 4 < 8
        if mid * mid > x {
            right = mid - 1
        }
        // if mid^2 < x, search in rhs
        if mid * mid < x {
            left = mid + 1
        }
        // if mid^2 == x, return mid 
        if mid * mid == x {
            return mid
        }
    }
    return right
}

// 0ms
// 4.08mb