package main

import (
	"os"
	"io"
	"fmt"
)

var test = "test_variable"

// Printt print some info
func Printt(){
	fmt.Println("hello world~!!!")
}

// VariableDefineMulti 变量的重声明
func VariableDefineMulti(){
	var err error
	n, err := io.WriteString(os.Stdout, "Hello, everyone!\n")
	if err != nil{
		fmt.Println(n)
	}
}

var contain = []string{"zero","one","two"}

// 一对不包裹任何东西的花括号，除了可以代表空的代码块之外，还可以用于表示不包含任何内容的数据结构（或者说数据类型）。
// 一定要记住，当整数值的类型的有效范围由宽变窄时，只需在补码形式下截掉一定数量的高位二进制数即可。
// 类似的快刀斩乱麻规则还有：当把一个浮点数类型的值转换为整数类型值时，前者的小数部分会被全部截掉。
// 第二，虽然直接把一个整数值转换为一个string类型的值是可行的，但值得关注的是，
// 被转换的整数值应该可以代表一个有效的 Unicode 代码点，否则转换的结果将会是"�"（仅由高亮的问号组成的字符串值


// 第三个知识点是关于string类型与各种切片类型之间的互转的。
func conversionString(){
	fmt.Println(string([]byte{'\xe4', '\xbd', '\xa0', '\xe5', '\xa5', '\xbd'}))
	fmt.Println(string([]rune{'\u4F60', '\u597D'})) // 你好
}

func main(){
	// 不同类型的重名变量，不受类型的限制
	contain := map[int]string{0:"zero",1:"one",2:"two"}
	fmt.Printf("The Element is %q.\n",contain[1])

	// go 判断一个变量的类型 使用类型断言表达式
	// value,ok := interface{}(contain).([]string)
	value,ok := interface{}(contain).(map[int]string)
	if ok{
		fmt.Println(value)
	}

	Printt()
	VariableDefineMulti()
	fmt.Println(string(-1))
	conversionString()
}

