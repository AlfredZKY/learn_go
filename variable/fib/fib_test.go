package fib

import (
	"testing"
)

func TestFibList(t *testing.T) {
	var a int = 1
	var b int = 1
	// ii := 0
	// _ = &(ii++)
	t.Log(a)
	for i := 0; i < 5; i++ {
		t.Log(b)
		tmp := a
		a = b
		b = tmp + b
	}
}

func TestExchange(t *testing.T) {
	// a := 1
	// b := 2
	a,b:=1,2
	t.Log(a, b)
	// tmp := a
	// a = b
	// b = tmp
	b,a=a,b
	t.Log(a, b)
}
