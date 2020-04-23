package algorithm


import (
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
