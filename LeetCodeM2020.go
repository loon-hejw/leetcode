package main

import (
	"fmt"
)

type ListNode struct {
	  Val int
	  Next *ListNode
}
func removeDuplicateNodes(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	check := make(map[int]bool)
	check[head.Val] = true
	pos := head
	for pos.Next != nil {
		cur := pos.Next
		if check[cur.Val] != false {
			pos.Next = pos.Next.Next
		} else {
			check[cur.Val] = true
			pos = pos.Next
		}
	}

	return head
}
func main() {
	
	check := make(map[int]bool)
	fmt.Println(check[4])
}
