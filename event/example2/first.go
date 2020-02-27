package main

import (
	"fmt"
)

// 创建一个事件响应map映射，事件名-事件函数切片
// 每个事件函数传递2个int64变量，返回bool和错误值
var evmap = map[string][]func(a, b int64) (bool, error){}

// func RegEv 事件注册 传递事件名，以及1个事件函数
func RegEv(name string, newf func(a, b int64) (bool, error)) {
	funclist := evmap[name]
	funclist = append(funclist, newf)
	evmap[name] = funclist
}

// DelEv 调用事件 传递事件名，事件函数的2个int64参数
func DelEv(name string) (bool, error) {
	delete(evmap, name)
	return true, nil
}

// CallEv 调用事件 传递事件名，事件函数的2个int64参数
func CallEv(name string, a, b int64) (bool, error) {
	funclist := evmap[name]
	if len(funclist) == 0 {
		return false, fmt.Errorf("没有%v的注册事件函数\n", name)
	}

	var br bool
	var err error
	for _, f := range funclist {
		br, err = f(a, b) // 调用事件函数
		if err != nil {
			break
		}
	}
	return br, err
}

func myadd(a, b int64) (bool, error) {
	if a == 0 || b == 0 {
		return false, fmt.Errorf("myadd函数调用失败:参数不能为0,a=%v\n,b=%v\n", a, b)
	}
	fmt.Println("myadd函数调用成功:", a+b)
	return true, nil
}

func mydiv(a, b int64) (bool, error) {
	if b == 0 {
		return false, fmt.Errorf("mydiv函数调用失败:除数不能为0,b=%v\n", b)
	}
	fmt.Println("mydiv函数调用成功:商->", a/b, "余数->", a%b)
	return true, nil
}

func main() {
	fmt.Println("hello world")
	RegEv("arithm1", myadd)
	RegEv("arithm1", mydiv)
	RegEv("arithm2", myadd)

	br, err := CallEv("arithm1", 11, 5)
	if !br {
		fmt.Println(err)
	}
	DelEv("arithm1")

	br, err = CallEv("arithm1", 11, 5)
	if !br {
		fmt.Println(err)
	}

	br, err = CallEv("arithm2", 29, 5)
	if !br {
		fmt.Println(err)
	}
}
