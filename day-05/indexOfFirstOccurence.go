package main

func StrStr(haystack string, needle string) int {
    for i := 0; i <= len(haystack) - len(needle); i++ {
        if haystack[i] == needle[0] {
            if haystack[i:i+len(needle)] == needle {
                return i
            }
        }
    }
    return -1
}

// 0ms
// 3.99 mb