package useerror

import (
	"errors"
	"fmt"
	"testing"
)




func TestPanic(t*testing.T){
	// 访问切片超出边界，引发运行恐慌
	array := []int{1,2,3,4}
	t.Log(array[2])
	panic("手动触发")
}

func divide(a,b int)(res int,err error){
	func(){
		defer func(){
			if rec := recover();rec != nil{
				err = fmt.Errorf("%s",rec)
			}
		}()
		res = a/b
	}()
	return
}

func TestDivide(t*testing.T){
	res,err := divide(1,0)
	t.Log(res,err)
}


func TestDeferSort(t*testing.T){
	// 这个队列与该defer语句所属的函数是对应的，并且，它是先进后出（FILO）的，相当于一个栈。
	defer fmt.Println("first defer")
	for i := 0; i < 3; i++ {
		defer fmt.Printf("defer in for [%d]\n", i)
	}
	defer fmt.Println("last defer")
}

func TestDeferCallPanic(t*testing.T){
	fmt.Println("Enter func main")
	defer func() {
		fmt.Println("Enter defer func")
		panic(errors.New("again some errors"))
		if p := recover(); p != nil {
			fmt.Printf("panic:%s\n", p)
		}
		fmt.Println("Exit defer")
	}()
	fmt.Println("Begin panic")
	panic(errors.New("some errors"))
	fmt.Println("Exit main")
}