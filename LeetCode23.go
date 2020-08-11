package main


func mergeKLists(lists []*ListNode) *ListNode {
	return merge(lists , 0 , len(lists))
}

func merge( lists []*ListNode , left , right int ) *ListNode {

	if left == right {
		return lists[left]
	}
	if left > right {
		return nil
	}
	mid := (left + right) >> 1
	return mergeTwoLists(merge(lists , left , mid -1 ) , merge( lists , mid + 1 ,right))
}
func main() {
	
}
