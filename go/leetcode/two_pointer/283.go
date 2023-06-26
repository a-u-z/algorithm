package main

import "fmt"

func moveZeroes(nums []int) {
	if len(nums) == 0 {
		return
	}

	left := 0 // 左指针，指向当前已处理好的序列的尾部
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			// 非零元素与当前已处理好的序列的尾部交换
			nums[left], nums[i] = nums[i], nums[left]
			left++
		}
		fmt.Println(nums) // 打印每一轮的结果
	}
}
