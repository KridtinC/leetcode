package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	var substr = map[string]int{}

	for i := 0; i < len(s); i++ {
		substr[s[:i+1]]++
		substr[s[i:]]++
	}

	var max = 0
	var maxstr string
	for str, v := range substr {
		if v > max {
			max = v
			maxstr = str
		}
	}

	fmt.Println(substr)

	return len(maxstr)
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
