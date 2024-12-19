package main

import (
	"strconv"
	"strings"
)

func AddBinary(a string, b string) string {
	// given: 2 binary strings
	// toReturn: their sum as another binary string

	// Binary addition is similar to decimal addition but in base 2.
	// Rules for Binary addition are:
	// 0 + 0 = 0
	// 0 + 1 = 1
	// 1 + 0 = 1
	// 1 + 1 = 10 (0 with carry 1)
	// 1 + 1 + 1 = 11 (1 with carry 1)

	// approach
	// start addition on the least significant side
	// or simply the rightmost side
	// say a = 11 + b = 1
	// first lets check whose length is more
	// a[len(a) - 1] + b[len(b) - 1], apply addition rules
	// initialize a carry container and update that as per need
	// declare and update an output string as per each step
	// return output :)

	carry := 0
	var output string

	// a is always the longer string
	if len(a) < len(b) {
		a, b = b, a
	}

	// Align shorter string with longer string
	diff := len(a) - len(b)
	b = strings.Repeat("0", diff) + b

	for i := len(a) - 1; i >= 0; i-- {
		bitA := int(a[i] - '0')
		bitB := int(b[i] - '0')

		sum := bitA + bitB + carry
		output = strconv.Itoa(sum%2+'0') + output
		carry = sum / 2
	}

	if carry > 0 {
		output = "1" + output
	}

	return output
}

// 0ms
// 5.56mb