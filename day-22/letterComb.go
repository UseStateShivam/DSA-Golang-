package main

func LetterCombinations(digits string) []string {
    if len(digits) == 0 {
        return []string{}
    }
    
    var digitsMap = map[byte]string{
        '2' : "abc",
        '3' : "def",
        '4' : "ghi",
        '5' : "jkl",
        '6' : "mno",
        '7' : "pqrs",
        '8' : "tuv",
        '9' : "wxyz",
    }
    
    res := []string{}
    
    var backtrack func(index int, comb string)
    backtrack = func(index int, comb string) {
        if len(comb) == len(digits) {
            res = append(res, comb)
            return
        }
        letters := digitsMap[digits[index]]
        for i := 0; i < len(letters); i++ {
            backtrack(index+1, comb+string(letters[i]))
        }
    }
 
    backtrack(0, "")
    return res

}

// 0ms
// 4.08mb