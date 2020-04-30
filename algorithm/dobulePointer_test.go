package algorithm

import "testing"


func TestDobulePointer(t *testing.T){
	nums := []int{-1,0,1,2,-1,-4}
	t.Log(nums)
	QuickSort(nums)
	t.Log(nums)
	res := DobulePointer(nums)
	t.Log(res)
}



func TestDobulePointer1(t *testing.T){
	nums := []int{-1,0,1,2,-1,-4}
	t.Log(nums)
	QuickSort(nums)
	t.Log(nums)
	res := DobulePointer1(nums)
	t.Log(res)
}