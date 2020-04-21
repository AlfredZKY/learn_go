package typetry

import "testing"

type MyInt int64

func TestImplicit(t *testing.T) {
	// 不支持隐式类型转换，包括别名也是不支持的
	var a int32 = 1
	var b int64
	//b = a // cannot use a (type int) as type int64 in assignment
	b = int64(a)
	var c MyInt
	//c = b	// cannot use b (type int64) as type MyInt in assignment
	c = MyInt(b)
	t.Log(a,b,c)
}

func TestPoint(t *testing.T){
	a :=1
	aPtr := &a
	// go中不支持指针运算
	//aPtr = aPtr + 1 // invalid operation: aPtr + 1 (mismatched types *int and int)
	t.Log(a,aPtr)
	t.Logf("%T %T",a,aPtr)
}

func TestString(t*testing.T){
	// go字符串默认是空字符串，而不是nil
	var a string
	t.Log("*" + a + "*")
	t.Log(len(a))
}
