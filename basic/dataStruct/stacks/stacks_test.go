package stacks

import (
	"testing"
)

func TestStacK_Len(t*testing.T){
	var myStack Stack
	myStack.Push(1)
	myStack.Push(2)
	myStack.Push("test")
	if myStack.Len() == 3 {
		t.Log("Pass Stack.Len")
	}else{
		t.Error("Failed Stack.Len")
	}
}

func TestStacK_Cap(t*testing.T){
	myStack := make(Stack,3)
	if myStack.Cap() == 3 {
		t.Log("Pass Stack.Len")
	}else{
		t.Error("Failed Stack.Len")
	}
}

func TestStacK_IsEmpty(t*testing.T){
	var myStack Stack
	if myStack.IsEmpty() {
		t.Log("Pass Stack.IsEmpty")
	}else{
		t.Log("Failed Stack.IsEmpty")
	}
}

func TestStacK_Pop(t*testing.T){
	var myStack Stack 
	// myStack.Push(0)
	// myStack.Push(1)
	value,err := myStack.Pop()
	if err == nil{
		t.Log("Pass Stack.Pop and value is ",value)
	}else{
		t.Error("Failed Stack.Pop and err is ",err)
	}
}

func TestStack_Push(t*testing.T){
	var myStack Stack
	myStack.Push(3)
	if myStack.Len() == 1 {
		t.Log("Pass Stack.Push")
	}else{
		t.Error("Failed Stack.Push")
	}
}

func TestStacK_Top(t*testing.T){
	var myStack Stack
	myStack.Push(0)
	myStack.Push(1)
	myStack.Pop()
	myStack.Push("test")

	value ,err := myStack.Top()
	if err == nil {
		t.Log("the top element of stack is ",value)
	}else {
		t.Error("Failed Stack.Top and err is ",err)
	}
}