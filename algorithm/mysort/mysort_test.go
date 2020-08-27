package mysort

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
)

// go语言的slice() 不仅可以对int类型的数组进行排序，也可以对struct类型的数组进行排序
// 排序函数如下
// 1.Slice() 排序不稳定
// 2.SliceStable() 稳定排序
// 3.SlicesSorted()判断是否已排序

type test struct {
	value int
	str   string
}

func TestSortSlices(t *testing.T) {
	s := make([]test, 5)
	s[0] = test{value: 4, str: "test1"}
	s[1] = test{value: 2, str: "test2"}
	s[2] = test{value: 3, str: "test3"}
	s[3] = test{value: 5, str: "test5"}
	s[4] = test{value: 1, str: "test4"}

	fmt.Println("初始化的结果")
	fmt.Println(s)
	// 从小到大不稳定排序
	// sort.Slice(s,func(i,j int)bool{
	// 	if s[i].value < s[j].value {
	// 		return true
	// 	}
	// 	return false
	// })
	// fmt.Println("从小到大排序的结果")
	// fmt.Println(s)

	// 打乱
	rand.Shuffle(len(s), func(i, j int) {
		s[i], s[j] = s[j], s[i]
	})
	fmt.Println("打乱的结果")
	fmt.Println(s)
	// 从小到大稳定排序
	sort.SliceStable(s, func(i, j int) bool {
		if s[i] == s[j] {
			fmt.Println("==")
			return s[i].value == s[j].value
		} else if s[i].value < s[j].value {
			fmt.Println("<")
			// return true
			return s[i].value < s[j].value
		}
		// return s[i].value > s[j].value
		return false
	})
	fmt.Println("从小到大排序的结果")
	fmt.Println(s)

	// 判断数组是否已排序
	bless := sort.SliceIsSorted(s, func(i, j int) bool {
		if s[i].value < s[j].value {
			return true
		}
		return false
	})
	fmt.Printf("S 是否已完成排序 %v\n", bless)
}
