package main

// func F(n int) int {
//     if n == 1 {
//         return 1
//     } else if n == 0 {
//         return 0
//     } else {
//         return F(n-1) + F(n-2)
//     }
// }

func ClimbStairs(n int) int {
    // given: n steps in a staircase
    // each time, we can either climb 1 or 2 steps
    // to return: number of distinct ways to climb to top

    // approach
    // say n = 4
    // the combinations can be: 1,1,1,1; 1,2,1; 1,1,2; 2,1,1; 2,2;
    // the result will depend upon our first step
    // if N1 = 1 then n left = n-1 
    // and n-2 otherwise
    // say for n = 4, N1 = 1, n left = 3, and ways = 3
    // and if N1 = 2, n left = 2, and ways = 2
    // total ways = 3 + 2 = 5
    // this tree can be defined as a fibonacci series tree
    // 0,1,1,2,3,5,8,13,21....... so on
    // say n = 1, F(n) = 1
    // say given n = 1, update n++
    // n = 2, F(n) = F(2) = F(2-1) + F(2-2) = F(1) + F(0) = 1 + 0 = 1

    // return F(n+1)
    // Although this is a conceptually correct approach, however, as the n increases, the time complexity of the algorithm increases exponentially
    
	if n == 1 {
        return 1
    } else if n == 0 {
        return 0
    }

    a, b := 1, 1
    for i := 2; i <= n; i++ {
        a, b = b, a+b
    }
    return b
}

// 0ms
// 3.85mb