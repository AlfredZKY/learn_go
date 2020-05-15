package processlink

import "testing"

func TestCreateSingleLink(t *testing.T) {
	head := CreateSingleLink()
	head.Traverse()
}

func TestReverserSignleListIter(t *testing.T) {
	head := CreateSingleLink()
	head.Traverse()
	res := ReverserSignleListIter(head)
	res.Traverse()
}

func TestReverserSignleListRecursive(t *testing.T) {
	head := CreateSingleLink()
	head.Traverse()
	res := ReverserSignleListRecursive(head)
	res.Traverse()
}

func TestSwapPairs1(t *testing.T) {
	head := CreateSingleLink()
	head.Traverse()
	res := SwapPairsyOne(head)
	res.Traverse()
}

func TestSwapPairs2(t *testing.T) {
	head := CreateSingleLink()
	head.Traverse()
	res := SwapPairsyTwo(head)
	res.Traverse()
}

func TestSwapPairs3(t *testing.T) {
	head := CreateSingleLink()
	head.Traverse()
	res := SwapPairsyThree(head)
	res.Traverse()
}

func TestReverseGroup(t *testing.T) {
	head := CreateSingleLink()
	head.Traverse()
	res := ReverseGroup(head, 3)
	res.Traverse()
}

func TestReverseGroup1(t *testing.T) {
	head := CreateSingleLink()
	head.Traverse()
	res := ReverseGroupStack(head, 3)
	res.Traverse()
}

func TestHasCycle(t *testing.T) {
	head := CreateCycle()
	res := HasCyclePointer(head)
	t.Log(res)
}

func TestHasCycleMap(t *testing.T) {
	head := CreateCycle()
	res := HasCycleMap(head)
	t.Log(res)
}

func TestHasCycleSomeVal(t *testing.T) {
	head := CreateCycle()
	res := HasCycleSomeVal(head)
	t.Log(res)
}

func TestDetectCycle(t*testing.T){
	head := CreateCycle()
	res := DetectCycle(head)
	if res == nil {
		t.Error("no cycle link")
	}
	t.Log(res.Val)
	val,err := head.Get(res.Val)
	if err == nil {
		t.Log("find a element ", val)
	}else{
		t.Error(err)
	}
}
