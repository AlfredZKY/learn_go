package except

import (
	"errors"
	"fmt"
	"testing"
)

func TestTryPanic(t *testing.T) {
	sInt := []int{1, 2, 3, 4, 5}
	_ = sInt
	//t.Log(sInt[5])
	// caller1()
	test()
}

func caller1() {
	fmt.Println("Enter function caller1.")
	caller2()
	fmt.Println("Enter function caller1.")
}

func test() {
	fmt.Println("Enter function test.")
	panic(errors.New("something wrong"))
	panic(fmt.Println)
	fmt.Println("Enter function test.")
}

func caller2() {
	fmt.Println("Enter function caller2.")
	sInt := []int{1, 2, 3, 4, 5}
	_ = sInt
	_ = sInt[5]
	fmt.Println("Exit function caller2.")
}

func TestRecoverWay(t *testing.T) {
	fmt.Println("Enter test func TestRecoverWay.")
	defer func() {
		fmt.Println("Enter defer function.")
		// recover函数的正确用法
		if p := recover(); p != nil {
			fmt.Printf("panic:%s\n", p)
		}
		fmt.Println("Exit defer function.")
	}()

	// recover函数的错误用法
	fmt.Printf("no panic:%v\n", recover())

	// 触发panic
	panic(errors.New("something wrong."))
	// caller2()
	p := recover()
	fmt.Printf("panic:%s\n", p)
	fmt.Println("Exit test func TestRecoverWay.")
}

func TestDerferMulti(t*testing.T){
	defer fmt.Println("first defer")
	for i:=0;i<3;i++{
		defer fmt.Printf("defer in for [%d]\n", i)
	}
	defer fmt.Println("last defer")
}