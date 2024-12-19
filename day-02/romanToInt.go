package main

func IntSelector(roman string) int {
    switch {
        case roman == "I":
            return 1
        case roman == "V":
            return 5
        case roman == "X":
            return 10
        case roman == "L":
            return 50
        case roman == "C":
            return 100
        case roman == "D":
            return 500
        case roman == "M":
            return 1000
    }
    return 0
}

func RomanToInt(s string) int {
    finalOutput := 0
    for i := 0; i < len(s); i++ {
        if string(s[i]) == "I" && i < len(s)-1 {
            if string(s[i+1]) == "V" {
                finalOutput += 4
                i++
                continue
            } 
            if string(s[i+1]) == "X" {
                finalOutput += 9
                i++
                continue
            } 
        }
        if string(s[i]) == "X" && i < len(s)-1 {
            if string(s[i+1]) == "L" {
                finalOutput += 40
                i++
                continue
            } 
            if string(s[i+1]) == "C" {
                finalOutput += 90
                i++
                continue
            } 
        }
        if string(s[i]) == "C" && i < len(s)-1 {
            if string(s[i+1]) == "D" {
                finalOutput += 400
                i++
                continue
            } 
            if string(s[i+1]) == "M" {
                finalOutput += 900
                i++
                continue
            } 
        }
        finalOutput += IntSelector(string(s[i]))
    }
    return finalOutput
}

// 3ms
// 4.69mb
// to optimize further to 0ms we can use map instead of switch case