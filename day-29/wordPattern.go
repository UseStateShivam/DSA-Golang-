package main

import "strings"

func WordPattern(pattern string, s string) bool {
	w := strings.Split(s, " ")
	m1 := map[byte]string{}
	m2 := map[string]byte{}
	if len(pattern) != len(w) {
		return false
	}
	for i := 0; i < len(pattern); i++ {
		p := pattern[i]
		w1 := w[i]
		if v, e := m1[p]; e {
			if v != w1 {
				return false
			}
		} else {
			m1[p] = w1
		}
		if v2, e := m2[w1]; e {
			if v2 != p {
				return false
			}
		} else {
			m2[w1] = p
		}
	}
	return true
}

// 0ms
// 3.96mb