package main

import (
	"fmt"
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	// var substrmap = map[string]int{}
	var substr string
	var result string

	for i := 0; i < len(s); i++ {

		if strings.Contains(substr, string(s[i])) {
			substr = string(s[i])
		} else {
			substr += string(s[i])

			if len(substr) > len(result) {
				result = substr
			}
		}

		// substrmap[substr]++

		// substrmap[s[i:]]++
	}

	// var max = 0
	// var maxstr string
	// for str, v := range substrmap {
	// 	if v > max {
	// 		max = v
	// 		maxstr = str
	// 	}
	// }

	// fmt.Println(result)

	return len(result)
}

// abcabcbb

// a
// ab
// abc
// abca
// abcab
// abcabc
// abcabcb
// abcabcbb
//  bcabcbb
//   cabcbb

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}
