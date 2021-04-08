package main

func deleteDuplicates(head *ListNode) *ListNode {

	node := &ListNode{Val: 0 , Next: nil}
	cur  := node
	temp := []*ListNode{}
	for head != nil {
		if  len(temp) > 0 && temp[len(temp) - 1].Val == head.Val {
			for head != nil && temp[len(temp) - 1].Val == head.Val {
				head = head.Next
			}
			temp = temp[:len(temp) - 1]
		}
		if head != nil {
			temp = append(temp , head)
			head = head.Next
		}
	}
	for _ , v := range temp {
		cur.Next = v
		cur = cur.Next
		cur.Next = nil
	}
	return node.Next
}

func main() {

}
