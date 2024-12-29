package main
func GenerateParenthesis(n int) []string {
	r := []string{}
	var backtrack func(c string, o, cl int)
	backtrack = func(c string, o, cl int) {
		if len(c) == 2*n {
			r = append(r, c)
			return
		}
		if o < n {
			backtrack(c+"(", o+1, cl)
		}
		if cl < o {
			backtrack(c+")", o, cl+1)
		}
	}
	backtrack("", 0, 0)
	return r   
}

// 0ms
// 4.76mb