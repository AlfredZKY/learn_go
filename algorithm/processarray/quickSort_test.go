package algorithm


import (
	"sort"
	"testing"
)

func TestQuickSort(t *testing.T){
	nums := []int{-1,0,1,2,-1,-4}
	t.Log(nums)
	QuickSort(nums)
	t.Log(nums)
}

func TestThreadNums(t *testing.T){
	nums := []int{-1,0,1,2,-1,-4}
	res := ThreadNums(nums)
	t.Log(res)
}

func TestSort(t*testing.T){
	nums := []int{-1,0,1,2,-1,-4}
	t.Log(nums)
	sort.Ints(nums)
	t.Log(nums)
}

func TestQucikSortArray(t*testing.T){
	nums := []int{-1,0,1,2,-1,-4}
	t.Log(nums)
	QuickSortArray(nums)
	t.Log(nums)
}

func TestTreadSums(t*testing.T){
	nums := []int{-1,0,1,2,-1,-4}

	ret := ThreadNums(nums)
	t.Log(ret)
	
}