package main

func IsSubsequence(s string, t string) bool {
    if len(s) == 0 {
        return true
    }
    if len(t) == 0 {
        return false
    }
    sIndex, tIndex := 0, 0
    for sIndex < len(s) {
        if tIndex < len(t) {
            if s[sIndex] == t[tIndex] {
                sIndex++
            }
            tIndex++
        } else {
            return false
        }
    }
    return true
}

// 0ms
// 4.04mb