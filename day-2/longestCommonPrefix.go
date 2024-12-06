package main

func LongestCommonPrefix(strs []string) string {
	if len(strs[0]) == 0 {
		return ""
	}
	prefix := strs[0]
	for i := 0; i < len(strs); i++ {
		currString := strs[i]
        j := 0
		for j < len(prefix) && j < len(currString) && prefix[j] == currString[j] {
			j++
		}
        prefix = prefix[:j]
        if prefix == "" {
            break
        }
	}
	return prefix
}

// 0ms
// 4.37mb