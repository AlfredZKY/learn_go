package learnfunc

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

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

func TestReturnMultiValues(t *testing.T) {
	a, _ := ReturnMultiValues()
	t.Log(a)
	tsSF:=timeSpent(slowFun)
	t.Log(tsSF(10))
}

func Sum(ops ...int)int{
	res := 0
	for _,op := range ops{
		res +=op
	}
	return res
}

func TestVarParam(t*testing.T){
	t.Log(Sum(1,2,3,4))
	t.Log(Sum(1,2,3,4,5))
}

func Clear(){
	fmt.Println("Clear resources.")
}

func TestDefer(t*testing.T){
	defer Clear()
	fmt.Println("start")
	panic("err")  //panic下依然会执行defer修饰的函数
	//fmt.Println("End") // 执行不到的代码
}