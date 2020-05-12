package sortfunc

import (
	"fmt"
	"errors"
)

// 快速排序法
// 快速排序（Quicksort）是对冒泡排序的一种改进。由C. A. R. Hoare在1962年提出。它的基本思想是：通过一趟排序将要排序的数据分割成独立的两部分，
// 其中一部分的所有数据都比另外一部分的所有数据都要小，然后再按此方法对这两部分数据分别进行快速排序，整个排序过程可以递归进行，以此达到整个数据变成有序序列。

// 设要排序的数组是A[0]……A[N-1]，首先任意选取一个数据（通常选用第一个数据）作为关键数据，然后将所有比它小的数都放到它前面，所有比它大的数都放到它后面，这个过程称为一趟快速排序。
// 值得注意的是，快速排序不是一种稳定的排序算法，也就是说，多个相同的值的相对位置也许会在算法结束时产生变动。

// 一趟快速排序的算法是：
// 1）设置两个变量i、j，排序开始的时候：i=0，j=N-1；
// 2）以第一个数组元素作为关键数据，赋值给key，即 key=A[0]；
// 3）从j开始向前搜索，即由后开始向前搜索（j -- ），找到第一个小于key的值A[j]，A[i]与A[j]交换；
// 4）从i开始向后搜索，即由前开始向后搜索（i ++ ），找到第一个大于key的A[i]，A[i]与A[j]交换；
// 5）重复第3、4、5步，直到 I=J； (3,4步是在程序中没找到时候j=j-1，i=i+1，直至找到为止。找到并交换的时候i， j指针位置不变。另外当i=j这过程一定正好是i+或j-完成的最后令循环结束。）

// QuickSortArrays 对切片数组快速排序
func QuickSortArrays(nums []int) error {
	fmt.Println(nums)
	n := len(nums)
	if n < 2 {
		return errors.New("array elements number is too little")
	}

	// 设置哨兵元素，默认左边的第一个元素开始
	sentry := nums[0]
	i, length := 0, len(nums)-1
	for k := 1; k <= length; {
		if nums[k] < sentry {
			nums[k],nums[i] = nums[i],nums[k]
			k++
			i++
		}else{
			nums[k],nums[length] = nums[length],nums[k]
			length--
		}
	}
	QuickSortArrays(nums[:length])
	QuickSortArrays(nums[length+1:])

	return nil
}
