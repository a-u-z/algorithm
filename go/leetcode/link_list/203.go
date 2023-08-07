package main

import "fmt"

// 定義節點結構
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {

	// 為了處理開頭如果是需要被拿掉的 val
	// 所以建立一個 prev 來當做更上一層
	// 作用只是用來當更頭的節點
	prev := &ListNode{
		Val:  0,
		Next: head,
	}

	// 像是指針的概念，知道現在處理到哪一個物件
	// 直接修改 current 的 next
	current := prev
	for current.Next != nil { // 只要下一個不是空的就要繼續檢查
		// 符合題目要的，就更改 next
		if current.Next.Val == val {
			current.Next = current.Next.Next // =  nil 也可以
		} else {
			// 不符合題目要的，那就把 current = current 的下一個
			// 往下移動
			current = current.Next
		}
	}
	return prev.Next
}

// 打印链表
func printLinkedList(head *ListNode) {
	current := head
	for current != nil {
		fmt.Print(current.Val, " -> ")
		current = current.Next
	}
	fmt.Println("nil")
}

// 创建链表并返回头节点
func createLinkedList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	head := &ListNode{Val: nums[0]}
	current := head

	for i := 1; i < len(nums); i++ {
		newNode := &ListNode{Val: nums[i]}
		current.Next = newNode
		current = newNode
	}

	return head
}

// 創建連結串列
