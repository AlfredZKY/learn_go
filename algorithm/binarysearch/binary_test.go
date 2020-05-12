package binarysearch

import "testing"

func TestBinarySearch(t *testing.T) {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7}
	target := 3
	BinarySearch(nums, target)
}
