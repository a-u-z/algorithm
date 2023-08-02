package main

// https://www.bilibili.com/video/BV1YT411g7br/
func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{
		Val:  0,
		Next: head,
	}
	cur := dummy
	for cur.Next != nil && cur.Next.Next != nil {
		one := cur.Next
		two := cur.Next.Next
		three := cur.Next.Next.Next

		cur.Next = two
		two.Next = one
		one.Next = three

		cur = one
	}
	return dummy.Next
}
