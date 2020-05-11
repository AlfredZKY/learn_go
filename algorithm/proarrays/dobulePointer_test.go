package proarrays

import "testing"


func TestDobulePointer(t *testing.T){
	nums := []int{-1,0,1,2,-1,-4}
	t.Log(nums)
	QuickSort(nums)
	t.Log(nums)
	res := RemoveDuplicates2(nums)
	t.Log(res)
}


func TestDobulePointerArea(t*testing.T){
	nums := []int{1,8,6,2,5,4,8,3,7}
	res := DobulePointerArea(nums)
	t.Log(res)
}

func TestMoveZeroesSentry(t*testing.T){
	nums := []int{0,1,0,3,12}
	t.Log(nums)
	MoveZeroesSentry(nums)
	t.Log(nums)
}

func TestMoveZeroesCount(t*testing.T){
	nums := []int{0,1,0,3,12}
	t.Log(nums)
	MoveZeroesCount(nums)
	t.Log(nums)
}


func TestTwoNumsSum(t*testing.T){
	nums := []int{2, 7, 11, 15}
	target := 26
	res := TwoNumsSum(nums,target)
	t.Log(res)
}

func TestRemoveDuplicates1(t*testing.T){
	nums := []int{0,0,1,1,1,2,2,3,3,4}
	length := RemoveDuplicates1(nums)
	t.Log(nums)
	t.Log(length)
}


func TestRemoveDuplicates2(t*testing.T){
	nums := []int{0,0,1,1,1,2,2,3,3,4}
	length := RemoveDuplicates2(nums)
	t.Log(nums)
	t.Log(length)
}

func TestRemoveDuplicates3(t *testing.T){
	nums := []int{0,0,1,1,1,2,2,3,3,4}
	length := RemoveDuplicates3(nums)
	t.Log(nums)
	t.Log(length)
}

func TestRotate1(t*testing.T){
	nums := []int {1,2,3,4,5,6,7}
	k := 3
	res := Rotate1(nums,k)
	t.Log(res)
}

func TestRotate2(t*testing.T){
	nums := []int {1,2,3,4,5,6,7}
	k := 3
	res := Rotate2(nums,k)
	t.Log(res)
}