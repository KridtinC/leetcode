package main

import (
	"fmt"
	"strings"
)

func convert(s string, numRows int) string {
	if numRows <= 1 {
		return s
	}

	var results = make([]string, numRows)
	var it int
	var f = next
	for i := 0; i < len(s); i++ {
		if it == 0 {
			f = next
		} else if it == numRows-1 {
			f = prev
		}

		results[it] += string(s[i])
		it = f(it)
	}

	return strings.Join(results, "")
}

func next(it int) int {
	return it + 1
}

func prev(it int) int {
	return it - 1
}

func main() {
	fmt.Println(convert("PAYPALISHIRING", 1))
}

// P Y
// A P

// P     H
// A   S I
// Y  I  R
// P L   I G
// A     N

// 2 --> 2
// 3 --> 4
// 4 --> 6
// 5 --> 8
