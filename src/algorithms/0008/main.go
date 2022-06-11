package main

import (
	"fmt"
	"math"
)

const (
	INT32_MAX = 2147483647
	INT32_MIN = -2147483648
)

func asciiToInt(ascii int) int {
	return ascii - 48
}

func myAtoi(s string) int {
	var results []int
	var sign = 1
	var foundSign bool
	var foundDigit bool
	for _, char := range s {

		s := int(char)
		if s == 32 && (!foundSign && !foundDigit) {
			continue
		}
		if s == 43 && (!foundSign && !foundDigit) {
			foundSign = true
			sign = 1
			continue
		}
		if s == 45 && (!foundSign && !foundDigit) {
			foundSign = true
			sign = -1
			continue
		}
		if s >= 48 && s <= 57 {
			foundDigit = true
			results = append(results, asciiToInt(s))
		} else {
			break
		}
	}

	if len(results) > 10 { // case int overflow
		if sign == -1 {
			return INT32_MIN
		}
		return INT32_MAX
	}

	var result int
	for i := 0; i < len(results); i++ {
		pow := math.Pow10(len(results) - 1 - i)
		result += results[i] * int(pow)

	}

	result = sign * result
	if result > INT32_MAX {
		return INT32_MAX
	}
	if result < INT32_MIN {
		return INT32_MIN
	}

	return result
}

func main() {
	// WIP
	fmt.Println(myAtoi("-20000000000000000000"))
}
