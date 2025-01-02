package main

func IsAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    m := map[byte]int{}
    for i := 0; i < len(s); i++ {
        if _, e := m[s[i]]; e {
            m[s[i]]++
        } else {
            m[s[i]] = 1
        }
    }
    for i := 0; i < len(t); i++ {
        if _, e := m[t[i]]; e {
            if m[t[i]] == 1 {
                delete(m, t[i])
            } else {
                m[t[i]]--
            }
        } else {
            return false
        }
    }
    if len(m) == 0 {
        return true
    }
    return false
}

// 11ms
// 4.9mb