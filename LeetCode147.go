package main

func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	Temp := &ListNode{Next:head}
	last , curr := head , head.Next
	for curr != nil {
		if last.Val <= curr.Val {
			last = last.Next
		} else {
			prev := Temp
			for prev.Next.Val <= curr.Val {
				prev = prev.Next
			}
			last.Next = curr.Next
			curr.Next = prev.Next
			prev.Next = curr
		}
		curr = last.Next
	}
	return Temp.Next
}

func main() {
	
}
