package learnfunc

import (
	"fmt"
	"errors"
)

// 函数类型的声明
type Printer func(contents string)(n int,err error)

func PrintToStd(contents string)(bytesNum int,err error){
	return fmt.Println(contents)
}

type operate func(x int,y int)int

// func Calculate 方案1
func Calculate(x int,y int,op operate)(int ,error){
	if op == nil{
		return 0,errors.New("Invalid operation")
	}
	return op(x,y),nil
}

type calculateFunc func(x int,y int)(int,error)

// func GenCalculator 方案2
func GenCalculator(op operate) calculateFunc{
	return func(x int,y int)(int,error){
		if op == nil{
			return 0,errors.New("invalid operation")
		}
		return op(x,y),nil
	}
} 