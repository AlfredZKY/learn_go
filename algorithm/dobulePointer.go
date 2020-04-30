package algorithm

// DobulePointer 双指针
func DobulePointer(nums []int) int {
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

// DobulePointer1 双指针
func DobulePointer1(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}

	slow :=  0
	for i := 1; i < n;  {
		if nums[slow] < nums[i] {
			slow++
			nums[slow] = nums[i]
		
		}
		i++
	}

	return slow + 1
}
