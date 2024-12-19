package main

func LengthOfLastWord(s string) int {
    // given string s
    // s has words and spaces
    // toReturn: len(last word in s)
    // definition for a word: max substring ! containing spaces

    // approach
    // lets start our pointer from the last index
    // we will move our pointer towards the left until we see a char
    // and move until another space is read
    // update some external pointer
    // return that pointer

    length := 0
    for i := len(s) - 1; i >= 0; i-- {
        if s[i] != ' ' {
            length++
        } else if length > 0 {
            break
        }
    }
    return length
}

// 0ms
// 4.19mb