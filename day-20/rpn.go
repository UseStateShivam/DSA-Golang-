package main

import "strconv"

func EvalRPN(tokens []string) int {
	st := []int{}
	for i := 0; i < len(tokens); i++ {
		if tokens[i] != "+" && tokens[i] != "-" && tokens[i] != "*" && tokens[i] != "/" {
			num, _ := strconv.Atoi(tokens[i])
			st = append(st, num)
		}
		if tokens[i] == "+" {
			a := st[len(st)-1]
			b := st[len(st)-2]
			st = st[:len(st)-2]
			st = append(st, b+a)
		} else if tokens[i] == "-" {
			a := st[len(st)-1]
			b := st[len(st)-2]
			st = st[:len(st)-2]
			st = append(st, b-a)
		} else if tokens[i] == "*" {
			a := st[len(st)-1]
			b := st[len(st)-2]
			st = st[:len(st)-2]
			st = append(st, b*a)
		} else if tokens[i] == "/" {
			a := st[len(st)-1]
			b := st[len(st)-2]
			st = st[:len(st)-2]
			st = append(st, b/a)
		}
	}
	return st[0]
}

// 0ms
// 6.01mb