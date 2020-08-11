package main

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	result := new(ListNode)
	tmp    := result
	for l1 != nil && l2 != nil {
		next1 := l1.Next
		next2 := l2.Next
		if l1.Val > l2.Val {
			tmp.Next = l2
			l2 = next2
		} else {
			tmp.Next = l1
			l1 = next1
		}
		tmp = tmp.Next
	}
	if l1 != nil {
		tmp.Next = l1
	}
	if l2 != nil {
		tmp.Next = l2
	}
	return result.Next
}
func main() {
	
}
