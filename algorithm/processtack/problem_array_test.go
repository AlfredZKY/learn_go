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

func TestMaxSlidingWindow(t*testing.T){
	nums :=[]int{1,3,-1,-3,5,3,6,7}
	k:=3
	res := MaxSlidingWindow(nums,k)
	t.Log(res)
}

func TestDubleCycleQueue(t*testing.T){
	dq := ConstructorDCQ(3)
	t.Log(dq)
	res := dq.InsertFront(1)
	res = dq.InsertLast(2)
	res = dq.InsertFront(3)
	res = dq.InsertFront(4)
	t.Log(res)
	resNum := dq.GetRear()
	t.Log(resNum)
	res = dq.IsFull()
	t.Log(res)
	res = dq.DeleteLast()
	t.Log(res)
	res = dq.InsertFront(4)
	t.Log(res)
	resNum = dq.GetFront()
	t.Log(resNum)
}