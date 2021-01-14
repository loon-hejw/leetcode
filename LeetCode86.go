package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {

	small := &ListNode{}
	big   := &ListNode{}
	smallHead := small
	bigHead := big
	for head != nil {
		if head.Val < x {
			small.Next = head
			small = small.Next
		}else {
			big.Next = head
			big = big.Next
		}
		head = head.Next
	}
	big.Next = nil
	small.Next = bigHead.Next
	return smallHead.Next
}

func main() {
	
}
