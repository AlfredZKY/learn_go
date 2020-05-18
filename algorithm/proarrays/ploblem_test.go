package proarrays

import (
	"learn_go/algorithm/sortfunc"
	"sort"
	"testing"
)

func TestDobulePointerArea(t *testing.T) {
	nums := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	res := DobulePointerArea(nums)
	t.Log(res)
}

func TestMoveZeroesSentry(t *testing.T) {
	nums := []int{0, 1, 0, 3, 12}
	t.Log(nums)
	MoveZeroesSentry(nums)
	t.Log(nums)
}

func TestMoveZeroesCount(t *testing.T) {
	nums := []int{0, 0, 1, 0, 3, 12}
	t.Log(nums)
	MoveZeroesCount(nums)
	t.Log(nums)
}

func TestTwoNumsSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	res := TwoNumsSum(nums, target)
	t.Log(res)
}

func TestDobulePointer(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}
	t.Log(nums)
	sortfunc.QuickSortArrays(nums)
	t.Log(nums)
	res := RemoveDuplicates2(nums)
	t.Log(res)
}

func TestRemoveDuplicates1(t *testing.T) {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	length := RemoveDuplicates1(nums)
	t.Log(nums)
	t.Log(length)
}

func TestRemoveDuplicates2(t *testing.T) {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	length := RemoveDuplicates2(nums)
	t.Log(nums)
	t.Log(length)
}

func TestRemoveDuplicates3(t *testing.T) {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	length := RemoveDuplicates3(nums)
	t.Log(nums)
	t.Log(length)
}

func TestRotate1(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	res := Rotate1(nums, k)
	t.Log(res)
}

func TestRotate2(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	res := Rotate2(nums, k)
	t.Log(res)
}

func TestThreadNums(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}
	res := ThreadNums(nums)
	t.Log(res)
}

func TestSort(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}
	t.Log(nums)
	sort.Ints(nums)
	t.Log(nums)
}

func TestTreadSums(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}
	ret := ThreadNums(nums)
	t.Log(ret)
}

func TestClimbStairs(t *testing.T) {
	res := ClimbStairs(6)
	t.Log(res)
}

func TestPlusOne(t *testing.T) {
	// nums := []int{ 1, 9, 9}
	nums := []int{9}
	res := PlusOne(nums)
	t.Log(res)
}
