package main

func IsIsomorphic(s string, t string) bool {
    ms := map[byte]byte{}
    mt := map[byte]byte{}
    for i := 0; i < len(s); i++ {
        if c, e := ms[s[i]]; e {
            if c != t[i] {
                return false
            }
        } else {
            ms[s[i]] = t[i]
        }
        if c, e := mt[t[i]]; e {
            if c != s[i] {
                return false
            }
        } else {
            mt[t[i]] = s[i]
        }
    }
    return true
}

// 0ms
// 4.58mb