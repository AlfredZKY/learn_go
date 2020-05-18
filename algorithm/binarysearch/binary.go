package binarysearch

import (
	"fmt"
)

// 题目
// 假设按照升序排序的数组在预先未知的某个点上进行了旋转。( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。
// 请找出其中最小的元素。你可以假设数组中不存在重复元素。

//BinarySearch 二分查找
func BinarySearch(nums []int, target int) {
	n := len(nums)

	// 取余表示怕数组不够长
	target = target % n

	// 旋转数组
	copy(nums, append(nums[len(nums)-target:], nums[:len(nums)-target]...))
	fmt.Println(nums)
	left, right := 0, len(nums)-1
	for left < right {
		// 左移相当于除二
		mid := (left + right) >> 1
		if nums[mid] < nums[right] {
			right = mid 
		} else {
			left = mid + 1
		}
	}
	fmt.Println(nums[left])
}
