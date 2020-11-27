package main

func mergaNode ( head1 , head2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	temp , temp1 , temp2 := dummyHead , head1 , head2
	for temp1 != nil && temp2 != nil {
		if temp1.Val > temp2.Val {
			temp.Next = temp2
			temp2 = temp2.Next
		} else {
			temp.Next = temp1
			temp1 = temp1.Next
		}
		temp = temp.Next
	}

	if temp1 == nil {
		temp.Next = temp2
	} else if temp2 == nil {
		temp.Next = temp1
	}
	return dummyHead.Next
}

func SortNode(head, tail *ListNode) *ListNode {

	if head == nil {
		return head
	}

	if head.Next == tail {
		head.Next = nil
		return head
	}

	slow , fast := head , head
	for fast != tail {
		slow = slow.Next
		fast = fast.Next
		if fast != tail {
			fast = fast.Next
		}
	}

	mid := slow
	return mergaNode(SortNode(head,mid) , SortNode(mid , tail))
}

func sortList(head *ListNode) *ListNode {
	return SortNode(head , nil)
}



func main() {
	
}
