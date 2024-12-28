package main

func CanConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) == 0 {
		return true
	}

	if len(magazine) == 0 {
		return false
	}

	byteMap := map[byte]int{}

	for i := 0; i < len(ransomNote); i++ {
		_, exists := byteMap[ransomNote[i]]
		if !exists {
			byteMap[ransomNote[i]] = 1
		} else {
			byteMap[ransomNote[i]]++
		}
	}

	for j := 0; j < len(magazine); j++ {
		_, exists := byteMap[magazine[j]]
		if exists {
			if byteMap[magazine[j]] > 1 {
				byteMap[magazine[j]]--
			} else if byteMap[magazine[j]] == 1 {
				delete(byteMap, magazine[j])
			}
		}
	}

	if len(byteMap) == 0 {
		return true
	} else {
		return false
	}
}

// 8ms
// 5.70mb
