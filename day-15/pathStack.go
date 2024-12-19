package main

import "strings"

func SimplifyPath(path string) string {
	// we use split func
	// this makes an arr of dir
	dirArr := strings.Split(path, "/")
	s := []string{}

	for _, dir := range dirArr {
		if dir == "" || dir == "." {
			continue
		} else if dir == ".." {
			// pop and ignore
			if len(s) > 0 {
				s = s[:len(s)-1]
			}
		} else {
			s = append(s, dir)
		}
	}

	return "/" + strings.Join(s, "/")
}

// 0ms
// 5.16mb