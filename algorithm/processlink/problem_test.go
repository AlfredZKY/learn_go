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

func TestReverserSignleListReceive(t*testing.T){
	head := CreateSingleLink()
	head.Traverse()
	res := ReverserSignleListReceive(head)
	res.Traverse()
}

func TestSwapPairs(t*testing.T){
	head := CreateSingleLink()
	head.Traverse()
	res := SwapPairsyOne(head)
	res.Traverse()
}


