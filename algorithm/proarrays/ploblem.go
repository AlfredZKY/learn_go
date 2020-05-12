package proarrays

import (
	"fmt"
	"learn_go/algorithm/sortfunc"
)

// 题目
// 给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
// 说明：你不能倾斜容器，且 n 的值至少为 2。
// 仔细读完题，可知，通过双指针，双边紧逼来， 求解最大区域，暴力破解法不再考虑

// DobulePointerArea 求出最大区域
func DobulePointerArea(nums []int) int {
	if len(nums) < 2 {
		return 0
	}

	// 设置头尾下标
	maxRes := 0
	for font, tail := 0, len(nums)-1; font < len(nums); {
		if nums[font] <= nums[tail] {
			maxRes = max(maxRes, (tail-font)*nums[font])
			font++
		} else {
			maxRes = max(maxRes, (tail-font)*nums[tail])
			tail--
		}

	}
	return maxRes
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

// 题目
// 给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
// 说明:
// 必须在原数组上操作，不能拷贝额外的数组。
// 尽量减少操作次数。

// MoveZeroesSentry 移动0，解题1.采用map存储的方式  解题2.采用数组计数的方法
func MoveZeroesSentry(nums []int) {
	j := -1 // j代表当前第一个0元素 -1代表没0
	for i := 0; i < len(nums); i++ {
		// 找到并初始化第一个0的位置
		if nums[i] == 0 && j == -1 {
			j = i
		}

		// 如果当前数非0 就与前面的0000第一个0交换，也就是j的位置,然后j++ 用于继续记录0的位置
		if nums[i] != 0 && j >= 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
}

// MoveZeroesCount 解题2.采用数组计数的方法
func MoveZeroesCount(nums []int) {
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			count++
		} else if count >= 0 {
			temp := nums[i]
			nums[i] = nums[i-count]
			nums[i-count] = temp
		}
	}
}

// 题目
// 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
// 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。
// 给定 nums = [2, 7, 11, 15], target = 9
// 方法1. 采用map存储 value和index

// TwoNumsSum nums 和一个目标值 target
func TwoNumsSum(nums []int, target int) []int {
	res := []int{}
	dic := make(map[int]int)
	for index, value := range nums {
		if value, exist := dic[target-value]; exist {
			res = append(res, value)
			res = append(res, index)
		}
		dic[value] = index
	}
	return res
}

// 题目
// 给定一个排序数组，你需要在 原地 删除重复出现的元素，使得每个元素只出现一次，返回移除后数组的新长度。
// 不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。
// 删除重复数组中重复的元素，并返回修改后数组的长度

// RemoveDuplicates1 方法1 后面往前走，因为已排序，所以不再考虑大小 方法2 采用双指针
func RemoveDuplicates1(nums []int) int {
	for length := len(nums) - 1; length > 0; length-- {
		if nums[length] == nums[length-1] {
			nums = append(nums[:length], nums[length+1:]...)
		}
	}
	return len(nums)
}

// RemoveDuplicates2 双指针
func RemoveDuplicates2(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}

	slow, quick := 0, 1
	for quick < n {
		if nums[slow] < nums[quick] {
			slow++
			nums[slow] = nums[quick]
		}
		quick++
	}
	return slow + 1
}

// RemoveDuplicates3 前后下标对比
func RemoveDuplicates3(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	j := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[j] {
			j++
			nums[j] = nums[i]
		}
	}
	return j + 1
}

// 题目
// 给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。
// 输入: [1,2,3,4,5,6,7] 和 k = 3

// Rotate1 旋转数组 方法1 copy 取余元素交换
func Rotate1(nums []int, k int) []int {
	n := len(nums)
	k = k % n
	copy(nums, append(nums[len(nums)-k:], nums[:len(nums)-k]...))
	return nums
}

// Rotate2 旋转数组 方法2 先整体反转，再在小数组中局部反转,需要反转三次
func Rotate2(nums []int, k int) []int {
	fmt.Println(nums)
	reverse(nums)
	fmt.Println(nums)
	reverse(nums[:k%len(nums)])
	fmt.Println(nums)
	reverse(nums[k%len(nums):])
	fmt.Println(nums)
	return nums
}

func reverse(nums []int) {
	for i := 0; i < len(nums)/2; i++ {
		nums[i], nums[len(nums)-i-1] = nums[len(nums)-i-1], nums[i]
	}
}

// 求取三数之和
//

// ThreadNums get nums
func ThreadNums(nums []int) [][]int {
	ret := make([][]int, 0, 0)
	if len(nums) < 3 {
		return ret
	}
	sortfunc.QuickSortArrays(nums)

	for i := 0; i < len(nums)-1; i++ {
		j, length := i+1, len(nums)-1
		if nums[i] > 0 || nums[j]+nums[i] > 0 {
			break
		}

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j < length {
			if j > i+1 && nums[j] == nums[j-1] {
				j++
				continue
			}
			if length < len(nums)-2 && nums[length] == nums[length+1] {
				length--
				continue
			}

			if nums[i]+nums[j]+nums[length] > 0 {
				length--
			} else if nums[i]+nums[j]+nums[length] < 0 {
				j++
			} else {
				ret = append(ret, []int{nums[i], nums[j], nums[length]})
				length--
				j++
			}
		}
	}

	return ret
}

// TreadSums get thread nums sums
func TreadSums(nums []int) [][]int {
	ret := make([][]int, 0, 0)
	if len(nums) < 3 {
		return ret
	}
	// 思路：采用先对数据进行排序，然后用下一个元素和最后一个元素求和，双边紧逼的方法 当前后索引相等时，即可完成所有的元素的求和
	// 先排序
	sortfunc.QuickSortArrays(nums)

	for i := 0; i < len(nums); i++ {
		// 判断元素只要大于0,就跳出循环，因为已经排序了，不可能存在三数之和等于0的情况
		j, length := i+1, len(nums)-1
		if nums[i] > 0 || nums[j]+nums[i] > 0 {
			break
		}

		// i 必须大于0的情况下 去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j < length {

			// 前元素去重
			if j > i+1 && nums[j] == nums[j+1] {
				j++
				continue
			}

			// 后边元素也要去重
			if length < len(nums)-2 && nums[length] == nums[length+1] {
				length--
				continue
			}
			if nums[i]+nums[j]+nums[length] > 0 {
				length--
			} else if nums[i]+nums[j]+nums[length] < 0 {
				j++
			} else {
				ret = append(ret, []int{nums[i], nums[j], nums[length]})
				j++
				length--
			}
		}

	}
	return ret
}
