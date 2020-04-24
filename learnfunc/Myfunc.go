package learnfunc

import (
	"math/rand"
	"time"
	"fmt"
	"errors"
)

// Printer 函数类型的声明
type Printer func(contents string)(n int,err error)

// PrintToStd Print Str
func PrintToStd(contents string)(bytesNum int,err error){
	return fmt.Println(contents)
}

type operate func(x int,y int)int

// Calculate 方案1
func Calculate(x int,y int,op operate)(int ,error){
	if op == nil{
		return 0,errors.New("Invalid operation")
	}
	return op(x,y),nil
}

// CalculateFunc func
type CalculateFunc func(x int,y int)(int,error)

// GenCalculator 方案2
func GenCalculator(op operate) CalculateFunc{
	return func(x int,y int)(int,error){
		if op == nil{
			return 0,errors.New("invalid operation")
		}
		return op(x,y),nil
	}
} 

// ReturnMultiValues return multi values
func ReturnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

func slowFun(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

// Clear func
func Clear(){
	fmt.Println("Clear resources.")
}

// Sum func
func Sum(ops ...int)int{
	res := 0
	for _,op := range ops{
		res +=op
	}
	return res
}