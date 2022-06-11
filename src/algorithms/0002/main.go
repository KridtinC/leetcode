package main

import "fmt"

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	var result = &ListNode{}
	var commute int

	var resultIt = result
	for resultIt != nil {
		sum := commute
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		commute = sum / 10
		sum = sum % 10

		resultIt.Val = sum

		if l1 != nil || l2 != nil || commute > 0 {
			resultIt.Next = &ListNode{}
		}

		resultIt = resultIt.Next
	}

	return result
}

func main() {
	result := addTwoNumbers(
		&ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: nil}}},
		&ListNode{Val: 8, Next: &ListNode{Val: 8, Next: &ListNode{Val: 9, Next: nil}}},
	)

	for result != nil {
		fmt.Printf("%d", result.Val)
		result = result.Next
	}

	fmt.Println()

}
