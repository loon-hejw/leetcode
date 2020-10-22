package main

func reorderList(head *ListNode)  {

	HeadList := []*ListNode{}
	for head != nil {
		HeadList = append(HeadList , head)
		head = head.Next
	}

	i  , j := 0 , len(HeadList) - 1
	for i < j {
		HeadList[i].Next = HeadList[j]
		i ++
		if i == j {
			break
		}
		HeadList[j].Next = HeadList[i]
		j --
	}
	HeadList[i].Next = nil
}

func main() {
	
}
