package linkstruct

import "testing"

func TestLinkMoudle(t *testing.T) {
	linkList := New(0)
	linkList.Insert(1, 9)
	linkList.Insert(1, 99)
	linkList.Append(55)
	linkList.Append(66)
	linkList.Append(77)
	linkList.Append(88)
	linkList.Traverse()
	//linkList.Delete(2)
	linkList.DeleteTail()
	linkList.Traverse()
	linkList.Append(88)
	linkList.Traverse()
	res, err := linkList.Get(0)
	if err == nil {
		t.Logf("res is %d\n", res)
	} else {
		t.Logf("err is %s\n", err)
	}
}

func TestLinkMoudleLoop(t *testing.T) {
	head := New(0)
	headNode := head.Loop()
	headNode.Traverse()
}
