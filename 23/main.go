package main

import (
	"container/heap"
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type Node []*ListNode

func (n Node) Len() int {
	return len(n)
}

func (n Node) Less(i, j int) bool {
	return n[i].Val < n[j].Val
}

func (n Node) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func (n *Node) Push(x interface{}) {
	*n = append(*n, x.(*ListNode))
}

func (n *Node) Pop() interface{} {
	old := *n
	nLen := len(old)
	elem := old[nLen-1]
	*n = old[0 : nLen-1]
	return elem
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	node := &Node{}
	heap.Init(node)

	// 각 리스트의 첫 번째 노드만 힙에 추가
	for _, list := range lists {
		if list != nil {
			heap.Push(node, list)
		}
	}

	if node.Len() == 0 {
		return nil
	}

	dummy := &ListNode{}
	curr := dummy

	for node.Len() > 0 {
		// 최소값 노드 추출
		minNode := heap.Pop(node).(*ListNode)

		curr.Next = minNode
		curr = curr.Next

		if minNode.Next != nil {
			heap.Push(node, minNode.Next)
		}
	}

	curr.Next = nil

	return dummy.Next
}

func main() {
	result := mergeKLists([]*ListNode{
		{
			Val: -2,
			Next: &ListNode{
				Val: -1,
				Next: &ListNode{
					Val: -1,
					Next: &ListNode{
						Val: -1,
					},
				},
			},
		},
		//{
		//	Val: 1,
		//	Next: &ListNode{
		//		Val: 3,
		//		Next: &ListNode{
		//			Val:  4,
		//			Next: nil,
		//		},
		//	},
		//},
		{},
	})

	for result != nil { // 노드가 존재하는지 확인
		fmt.Println(result.Val) // 먼저 값을 출력
		result = result.Next    // 그 후에 다음 노드로 이동
	}
}
