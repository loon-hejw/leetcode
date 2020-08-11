package main

func swapPairs(head *ListNode) *ListNode {

	if head == nil || head.Next == nil{
		return head
	}

	fristList := head
	secondList := head.Next

	fristList.Next = swapPairs( secondList.Next )
	secondList.Next = fristList

	return secondList
}

func main() {
	
}
