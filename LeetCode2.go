package main


//Definition for singly-linked list.
type ListNode struct {
     Val int
     Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var result ListNode = ListNode{Val:0}
	var curr *ListNode = &result
	var overflow int = 0
	for l1 != nil && l2 != nil {
		var temp = l1.Val + l2.Val + overflow
		curr.Val =  temp % 10
		overflow = temp / 10
		l1 = l1.Next
		l2 = l2.Next
		if l1 != nil || l2 != nil {
			curr.Next = &ListNode{Val:0}
			curr = curr.Next
		}
	}

	for l1 != nil {
		var temp = l1.Val + overflow
		curr.Val =  temp % 10
		overflow = temp / 10
		l1 = l1.Next
		if l1 != nil {
			curr.Next = &ListNode{Val:0}
			curr = curr.Next
		}
	}

	for l2 != nil {
		var temp = l2.Val + overflow
		curr.Val =  temp % 10
		overflow = temp / 10
		l2 = l2.Next
		if l2 != nil {
			curr.Next = &ListNode{Val:0}
			curr = curr.Next
		}
	}

	if overflow > 0 {
		curr.Next = &ListNode{Val:overflow}
	}

	return &result
}