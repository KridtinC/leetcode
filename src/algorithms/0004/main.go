package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var it1, it2 int
	var results []int

	for it1 < len(nums1) && it2 < len(nums2) {
		if nums1[it1] < nums2[it2] {
			results = append(results, nums1[it1])
			it1++
		} else {
			results = append(results, nums2[it2])
			it2++
		}
	}

	for i := it1; i < len(nums1); i++ {
		results = append(results, nums1[i])
	}
	for i := it2; i < len(nums2); i++ {
		results = append(results, nums2[i])
	}

	halfIdx := len(results) / 2
	if len(results)%2 == 0 {
		return (float64(results[halfIdx-1]) + float64(results[halfIdx])) / 2
	}
	return float64(results[halfIdx])
}

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))
}
