package main

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode = nil

	curr := head

	for curr != nil {
		temp := curr.Next
		curr.Next = pre
		pre = curr
		curr = temp
	}
	return pre
}
