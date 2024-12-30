package main

func Calculate(s string) int {
    st := []int{}
    c := 0
    sign := 1
    for i := 0; i < len(s); i++ {
        if s[i] >= '0' && s[i] <= '9' {
            n := 0
            for i < len(s) && s[i] >= '0' && s[i] <= '9' {
                n = n*10 + int(s[i]-'0')
                i++
            }
            i--
            c += sign*n
        } else if s[i] == '+' {
            sign = 1
        } else if s[i] == '-' {
            sign = -1
        } else if s[i] == '(' {
            st = append(st, c)
            st = append(st, sign)
            c = 0
            sign = 1
        } else if s[i] == ')' {
            snew := st[len(st)-1]
            st = st[:len(st)-1]
            cnew := st[len(st)-1]
            st = st[:len(st)-1]
            c = cnew + snew*c
        }
    }
    return c
}

// 0ms
// 4.94mb