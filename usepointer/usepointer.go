package main

import (
	"unsafe"
	"fmt"
)

// Dog define a struct variable
type Dog struct{
	name string //名字
}

// New a object of Dog
func New(name string)Dog{
	return Dog{name}
}

// SetName set variable of name 
func (dog *Dog)SetName(name string){
	dog.name = name
}

// Name return a variable of name
func (dog Dog)Name()string{
	return dog.name
}

// Named interface
type Named interface {
	Name() string // 用于获取名字
}

func testAddress(){
	const num =123
	// _ = &num // 常量不可寻址

	// _ = &(123) // 基本类型值的字面量

	// 字符串变量
	var str = "abc"
	_ = str
	//_ = &(str[0]) // 对字符串变量的索引结果值不可寻址
	//_ = &(str[1:2]) // 对字符串变量的切片结果值不可寻址
	//_ = &(str[1:2]) // 对字符串变量的切片结果值不可寻址
	str1 := str[0] 
	_ = &str1	// 这样寻址是合法的，切片索引值的结果值，底层维护了一个切片数组

	//_ = &(123 + 456) //算术操作的结果值不可寻址
	num2 := 456
	_ = num2
	// _ = &(num2*2) 

	// _ = &([3]int{1,2,3}[0])	 // 对数组字面量的索引结果值不可寻址
	// _ = &([3]int{1,2,3}[0:1]) // 对数组字面量的切片结果值不可寻址
	_ = &([]int{1,2,3}[0])		 // 对切片字面量的索引结果值却是可寻址的
	//_ = &([]int{1,2,3}[0:1])     // 对切片字面量的切片结果值不可寻址
	// _ = &(map[int]string{1:'a'}[0]) // 对字典字面量的索引结果值不可寻址

	var arr = []int{1,2,3}
	_ = &(arr[0])			// 对数组变量的索引值是可以寻址的
	//_ = &(arr[1:2])		// 对数组变量的切片值是不可寻址的
	arrTemp := arr[1:2]
	_ = &arrTemp			// 对数组切片结果值的变量值是可以寻址的

	var map1 = map[int]string{1:"a",2:"b",3:"c"}
	_ = map1
	//_ = &(map1[0]) // 对字典变量的索引结果值不可寻址

	// _ = &(func(x,y int)int{ // 字面量代表的函数不可寻址
	//	return x+y
	//})
	//_ = &(fmt.Sprintf)	// 标识符代表的函数不可寻址
	//_ = &(fmt.Sprintln("abc")) // 对函数的调用结果值不可寻址

	dog := Dog{"little pig"}
	_ = dog

	//_ = &(dog.Name) // 标识符代表的函数不可寻址
	// _=&(dog.Name()) // 对方法的调用的结果值不可寻址

	// _ = &(Dog{"little pid"}.name) // 对结构体字面量的字段不可寻址

	// _ = & (interface{}(dog))	// 类型转化表达式的结果值不可寻址
	dogI := interface{}(dog)
	_ = dogI
	//_ = &(dogI.(Named))		// 类型断言表达式的结果值不可寻址
	named := dogI.(Named)
	_ = named
	//_ = &(named.(Dog))			// 类型断言表达式的结果值不可寻址

	var chan1 = make(chan int,1)
	chan1 <- 1
	// _ = &(<-chan1) // 接受表达式的结果值不可寻址

	// 通过链式手法调用SetName方法,New出来的对象是临时的，不可寻址，
	// 但是我们在一个基本类型的值上调用它的指针方法,这是因为Go语言会自动帮我们转译
	// 更具体地说，对于一个Dog类型的变量dog来说，调用表达式dog.SetName("monster")会被自动地
	// 转译为(&dog).SetName("monster"),即:先取dog的指针值，再在该指针上调用SetName方法。
	// New("little dog").SetName("monster")

	//Go 语言中的++和--并不属于操作符，而分别是自增语句和自减语句的重要组成部分。
	//只要在++或--的左边添加一个表达式，就可以组成一个自增语句或自减语句，
	//但是，它还明确了一个很重要的限制，那就是这个表达式的结果值必须是可寻址的。
	//这就使得针对值字面量的表达式几乎都无法被用在这里。
	i := 0
	_ = i
	//_ = &(i++)
	//_ = &(map1[i++])
}

func useUnsafe(){
	dog := Dog{"little pig"}
	dogP := &dog
	// 经过两次类型转化，并赋值给变量dogPtr 转化规则：
	// 1.一个指针值(*Dog类型的值)可以被转换为一个unsafe.Pointer类型的值，反之亦然
	// 2.一个uintptr类型的值也可以被转换为一个unsafe.Pointer类型的值，反之亦然
	// 3.一个指针值无法被直接转换成一个uintptr类型的值，反过来也是如此
	dogPtr := uintptr(unsafe.Pointer(dogP))
	_ = dogPtr
	namePtr := dogPtr + unsafe.Offsetof(dogP.name)

	// 尽量不使用不安全的指针类型操作
	nameP := (*string)(unsafe.Pointer(namePtr))
	fmt.Println(nameP)
}


func main(){
	fmt.Println("start use learn pointer")
	dog := Dog{"litter dog"}
	dog.SetName("monster")
	//testAddress()
	useUnsafe()
}