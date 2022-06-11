package main

import (
	"fmt"
	"math"
	"strconv"
)

func reverse(x int) int {

	var result string
	var isReverse bool
	if x < 0 {
		x = x * -1
		isReverse = true
	}
	xstr := strconv.Itoa(x)
	for idx := range xstr {
		result += string(xstr[len(xstr)-1-idx])
	}

	x, _ = strconv.Atoi(result)

	if float64(x) > math.Pow(2, 31)-1 || float64(x) < -math.Pow(2, 31) {
		return 0
	}

	if isReverse {
		return x * -1
	}
	return x
}

func main() {
	fmt.Println(reverse(1534236469))
}
