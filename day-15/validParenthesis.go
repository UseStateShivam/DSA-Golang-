package main

func IsValid(s string) bool {
    stack := []byte{}
    m := map[byte]byte{
        ')' : '(',
        '}' : '{',
        ']' : '[',
    }
    if len(s) % 2 != 0 {
        return false
    }
    for i := 0; i < len(s); i++ { 
        if s[i] == '(' || s[i] == '{' || s[i] == '[' {
            stack = append(stack, s[i])
        } else {
            if len(stack) == 0 || m[s[i]] != stack[len(stack) - 1] {
                return false
            } else {  
                stack = stack[: len(stack) - 1]
            }
        }
    } 
    if len(stack) == 0 {
        return true
    } else {
        return false
    }
} 

// 0ms
// 4.02mb