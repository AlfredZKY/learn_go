package processtack

import "testing"

func TestIsValid(t *testing.T) {
	s := "()"
	res := IsValid(s)
	t.Log(res)
}


func TestLargestRectangleArea(t*testing.T){
	s := []int{2,1,5,6,2,3}
	res := LargestRectangleArea(s)
	t.Log(res)
}