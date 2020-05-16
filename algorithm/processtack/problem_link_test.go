package processtack

import "testing"

func TestLinkStack(t *testing.T) {
	stack := ConstructorLink()
	stack.Push(-2)
	stack.Push(0)
	stack.Push(-3)
	stack.Push(-4)
	stack.Push(1)
	res := stack.GetMin()
	t.Log(res)
	stack.Pop()
	res = stack.Top()
	t.Log(res)
	res = stack.GetMin()
	t.Log(res)
}
