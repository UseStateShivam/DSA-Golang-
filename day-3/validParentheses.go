package main

func IsValid(s string) bool {
    stack := []byte{}
    matching := map[byte]byte{
        ')' : '(',
        '}' : '{',
        ']' : '[',
    }
    for i := 0; i < len(s); i++ {
        if open, exists := matching[s[i]]; exists {
            if len(stack) == 0 || stack[len(stack)-1] != open {
                return false
            } 
            stack = stack[:len(stack)-1]
        } else {
            stack = append(stack, s[i])
        }
    }
    return len(stack) == 0
}

// 0ms
// 4.05mb