package sortfunc


import "testing"

func TestQuickSortArrays(t*testing.T){
	// nums := []int{-1,0,1,2,-1,-4}
	nums := []int{6,0,1,2,-1,-4}
	t.Log(nums)
	err := QuickSortArrays(nums)
	if err == nil{
		t.Log(nums)
	}	
}