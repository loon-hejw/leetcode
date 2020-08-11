package main

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := new(ListNode)
	tmp    := result
	sum := 0
	for l1 != nil || l2 != nil || sum > 0 {
		tmp.Next = new(ListNode)
		tmp = tmp.Next
		if l1 != nil {
			sum = l1.Val + sum
			l1 = l1.Next
		}
		if l2 != nil {
			sum = l2.Val + sum
			l2 = l2.Next
		}
		tmp.Val = sum % 10
		sum = sum / 10
	}
	return result.Next
}

func main() {
	
}
