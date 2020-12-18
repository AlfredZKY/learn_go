package typejudgment

import (
	"reflect"
	"fmt"
	"testing"
)

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
	t.Log(a, b, c)
}

func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	// go中不支持指针运算
	//aPtr = aPtr + 1 // invalid operation: aPtr + 1 (mismatched types *int and int)
	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr)
}

func TestString(t *testing.T) {
	// go字符串默认是空字符串，而不是nil
	var a string
	t.Log("*" + a + "*")
	t.Log(len(a))
}

type Bag struct {
	Key string
}

type Bag2 struct {
	Key int
}

func TestVariableType(t *testing.T) {
	var b1 interface{}
	var b2 interface{}

	b1 = Bag{Key: "1"}
	b2 = Bag2{Key: 0}

	// 方法1 通过接口去判断
	types1, ok := b1.(Bag)
	if ok {
		t.Logf("Types is %q\n", types1)
	}

	types2, ok := b2.(Bag2)
	if ok {
		t.Logf("Types is %q\n", types2)
	}

	// 方法2 通过swith判断
	switch v := b1.(type) {
	case Bag:
		t.Logf("Types is %q\n", v)
	case Bag2:
		t.Logf("Types is %q\n", v)
	default:
		t.Log("Types is Unknow")
	}

}

type point struct {
	x, y int
}

func TestPrintParam(t *testing.T) {
	// go 提供了几种打印格式，用来格式话一般的go值，例如%v 打印一个point结构体的对象的值
	p := point{1, 2}
	fmt.Printf("%v\n", p)

	// 如果格式化的值是一个结构体对象，那么%+v 的格式化输出 将包括结构体的成员名称和值
	fmt.Printf("%+v\n", p)

	// '%#v'格式化出处将输出一个值的go语法表示方式
	fmt.Printf("%#v\n", p)

	// %T 用于输出一个值的数据类型
	fmt.Printf("%T\n", p)

	// %t 用于输出布尔型变量
	fmt.Printf("%t\n", true)

	// 数值型
	{
		// %d 以十进制来输出整型的方式
		fmt.Printf("%d\n", 123)

		// %b 以二进制输出整型的方式
		fmt.Printf("%b\n", 14)

		// %x 以十六进制打印
		fmt.Printf("%x\n", 456)

		// %f 以浮点型输出
		fmt.Printf("%f\n", 78.9)

		// %c 输出整型数值对应的字符
		fmt.Printf("%c\n", 99)

		// %e %E 科学计数法输出整型
		fmt.Printf("%e\n", 12340000000.0)
		fmt.Printf("%E\n", 12340000000.0)
	}

	// 字符串
	{
		// %s 输出基本的字符串
		fmt.Printf("%s\n", "\"strings\"")

		// %q 输出带有双引号的字符串
		fmt.Printf("%q\n", "\"strings\"")
		fmt.Printf("%x\n", "hex this")
	}
}

func TestTypeJudgement(t*testing.T){
	var i interface{} = "kk"
	j,err := i.(int)
	if err {
		t.Logf("%T->%v\n",j,j)
	}else{
		t.Log("err is ",err)
	}
	t.Logf("%T->%v\n",j,j)
}

func TestMakeNewType(t*testing.T){
	var i *int = new(int) 
	*i = 10
	t.Log(*i,"\t",reflect.TypeOf(i))

	var j = make([]int,7)
	t.Log(j,"\t",reflect.TypeOf(j))
}