package operator

import "testing"

const(
	Readable = 1<< iota
	Writable
	Executable
)

func TestBitClear(t*testing.T){
	a:=7 // 0111
	// &^ 按位清零运算符,只考虑右边，将运算符左边数据相异的位保留，相同位清零。
	// 右边为0，左边数据保持不变 右边为1，左边数据清零
	t.Log(Readable,Writable,Executable)
	a = a &^ Readable
	t.Log(a)
	t.Log(a&Readable)
	a = a &^ Executable
	t.Log(a)
	t.Log(a&Executable==Executable)
}

func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 2, 34, 5}
	c := [...]int{1, 2, 3, 4, 5}
	d := [...]int{1, 2, 3, 4}

	t.Log(a == b)
	_ = c 
	//t.Log(a == c) //invalid operation: a == c (mismatched types [4]int and [5]int)
	t.Log(a == d)
}
