package main

func reverseBetween(head *ListNode, m int, n int) *ListNode {

	var next *ListNode
	dmy  := &ListNode{0 ,head}
	cut  := dmy
	tail := head
	for m > 1 {
		cut = tail
		tail = tail.Next
		m --
	}
	delta := n - m
	for delta > 0 {
		next = tail.Next
		tail.Next = next.Next
		next.Next = cut.Next
		cut.Next  = next
		delta --
	}
	return dmy.Next
}
func main() {
	
}
