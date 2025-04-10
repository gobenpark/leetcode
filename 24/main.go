package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	first := head
	second := head.Next

	first.Next = swapPairs(second.Next)
	second.Next = first

	return second
}

func main() {
	result := swapPairs(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
				},
			},
		},
	})

	for result != nil {
		fmt.Println(result.Val)
		result = result.Next
	}
}
