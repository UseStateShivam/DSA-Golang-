package main

func FindSubstring(s string, words []string) []int {
    if len(words) == 0 || len(words[0]) == 0 {
        return []int{}
    }

    wordLen := len(words[0])
    wordCount := len(words)
    totalLen := wordLen * wordCount
    wordsMap := make(map[string]int)

    // Initialize wordsMap
    for _, word := range words {
        wordsMap[word]++
    }

    output := []int{}

    // Iterate through each possible starting point within the word length
    for i := 0; i < wordLen; i++ {
        l, r := i, i
        currMap := make(map[string]int)

        for r+wordLen <= len(s) {
            // Extract the word at the current right pointer
            word := s[r : r+wordLen]
            r += wordLen

            if freq, exists := wordsMap[word]; exists {
                currMap[word]++

                // If word frequency exceeds the allowed frequency, shrink the window
                for currMap[word] > freq {
                    leftWord := s[l : l+wordLen]
                    currMap[leftWord]--
                    l += wordLen
                }

                // Check if the window is valid
                if r-l == totalLen {
                    output = append(output, l)
                }
            } else {
                // If the word doesn't exist in words[], reset the window
                currMap = make(map[string]int)
                l = r
            }
        }
    }

    return output
}

// 8ms
// 8.74mb