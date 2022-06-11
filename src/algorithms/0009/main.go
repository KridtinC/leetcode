package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	var val = x
	var results = []int{}
	for val > 0 {
		digit := val % 10
		val = val / 10
		results = append(results, digit)

	}

	var mid = len(results) / 2
	var left = 0
	var right = len(results) - 1
	for left < mid || right > mid {

		if results[left] != results[right] {
			return false
		}

		left++
		right--
	}
	return true
}

func main() {
	fmt.Println(isPalindrome(12321))
}
