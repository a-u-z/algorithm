package main

import "fmt"

type MyLinkedList struct {
	Head *ListNode
	Size int
}

func main() {
	list := Constructor()

	// 測試 AddAtHead 和 Get
	list.AddAtHead(1)
	list.AddAtHead(2)
	list.AddAtHead(3)
	fmt.Println("Current List: ", list.Head) // 預期輸出：3->2->1
	fmt.Println("Get(0): ", list.Get(0))     // 預期輸出：3
	fmt.Println("Get(1): ", list.Get(1))     // 預期輸出：2

	// 測試 AddAtTail 和 DeleteAtIndex
	list.AddAtTail(4)
	fmt.Println("Current List: ", list.Head) // 預期輸出：3->2->1->4
	list.DeleteAtIndex(2)
	fmt.Println("Current List: ", list.Head) // 預期輸出：3->2->4
}

func Constructor() MyLinkedList {
	return MyLinkedList{} // 初始化一個空的鏈表
}

func (this *MyLinkedList) GetV2(index int) int {
	// edge case
	if index < 0 || index > this.Size-1 {
		return -1
	}
	current := this.Head

	for i := 0; i < index; i++ {
		current = current.Next
	}
	return current.Val
}

func (this *MyLinkedList) AddAtHeadV2(val int) {
	this.Size++

	// edge case
	if this.Head == nil {
		this.Head = &ListNode{Val: val}
		return
	}

	new := &ListNode{
		Val:  val,
		Next: this.Head,
	}
	this.Head = new
}
func (this *MyLinkedList) AddAtTaiV2l(val int) {
	this.Size++
	new := &ListNode{Val: val}

	// this.Size 空
	if this.Head == nil {
		this.Head = new
	}
	// this.Size 有東西
	// range 他
	current := this.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = new
}

func (this *MyLinkedList) AddAtIndexV2(index int, val int) {

	// index <0 index >=this.Size
	if index < 0 || index >= this.Size {
		return
	}
	// index == 0
	if index == 0 {
		this.AddAtHeadV2(val)
		return
	}

	// index == this.Size
	if index == this.Size {
		this.AddAtTaiV2l(val)
		return
	}
	if index > 0 && index < this.Size {
		this.Size++
		current := this.Head
		for i := 0; i < index-1; i++ {
			current = current.Next
		}
		new := &ListNode{
			Val:  val,
			Next: current.Next,
		}
		current.Next = new
		return
	}
	// index 在之間
}
func (this *MyLinkedList) DeleteAtIndexV2(index int) {
	// index <0 index >=this.Size
	if index < 0 || index >= this.Size {
		return
	}
	this.Size--
	if index == 0 {
		this.Head = this.Head.Next
		return
	}
	current := this.Head
	for i := 0; i < index-1; i++ {
		current = current.Next
	}
	current.Next = current.Next.Next
	// index 在之間
}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.Size {
		return -1 // 索引超出範圍，返回 -1 表示無效索引
	}

	curr := this.Head // 用一個指針 curr 遍歷鏈表
	for i := 0; i < index; i++ {
		curr = curr.Next // 移動指針到下一個節點，直到達到指定索引位置
	}

	return curr.Val // 返回指定索引位置節點的值
}

func (this *MyLinkedList) AddAtHead(val int) {
	new := &ListNode{
		Val:  val,       // 創建一個新的節點，設置值為 val
		Next: this.Head, // 新節點的 Next 指向鏈表的頭節點
	}
	this.Head = new // 更新鏈表的頭節點為新節點
	this.Size++     // 增加鏈表的大小
}

func (this *MyLinkedList) AddAtTail(val int) {
	new := &ListNode{
		Val: val, // 創建一個新的節點，設置值為 val
	}
	// edgecase，如果鏈表是空的，直接將頭節點指向新節點
	if this.Size == 0 {
		this.Head = new
		this.Size++
		return
	}

	current := this.Head // 用一個指針 current 遍歷鏈表
	for current.Next != nil {
		current = current.Next // 移動指針到下一個節點，直到找到最後一個節點
	}
	current.Next = new // 將最後一個節點的 Next 指向新節點，即將新節點插入到尾部
	this.Size++        // 增加鏈表的大小
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	// 分析 index 在開頭
	if index == 0 {
		this.AddAtHead(val) // 在頭部插入節點
		return
	}

	// 分析 index 在結尾
	// this.Size 不用減一的原因是 index 是要成為第X個
	if index == this.Size {
		this.AddAtTail(val) // 在尾部插入節點
		return
	}

	// 分析 index 在中間
	if index > 0 && index < this.Size {
		current := this.Head // 用一個指針 current 遍歷鏈表
		for i := 0; i < index-1; i++ {
			current = current.Next // 移動指針到指定索引前一個節點
		}
		new := &ListNode{
			Val:  val,          // 創建一個新的節點，設置值為 val
			Next: current.Next, // 新節點的 Next 指向原來的下一個節點
		}
		current.Next = new // 將原來節點的 Next 指向新節點，即在指定位置插入新節點
		this.Size++        // 增加鏈表的大小
		return
	}
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	// edgecase
	if index < 0 || index >= this.Size {
		return // 索引超出範圍，不需要刪除節點，直接返回
	}
	if index == 0 {
		this.Head = this.Head.Next // 刪除頭節點，直接將頭節點指向下一個節點
		this.Size--                // 減少鏈表的大小
		return
	}
	current := this.Head // 用一個指針 current 遍歷鏈表
	for i := 0; i < index-1; i++ {
		current = current.Next // 移動指針到指定索引前一個節點
	}
	current.Next = current.Next.Next // 刪除指定位置的節點，將前一個節點的 Next 指向下下個節點
	this.Size--                      // 減少鏈表的大小
}
