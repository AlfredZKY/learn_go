package learnfunc

import (
	"fmt"
	"testing"

)

func TestReturnMultiValues(t *testing.T) {
	a, _ := ReturnMultiValues()
	t.Log(a)
	tsSF:=timeSpent(slowFun)
	t.Log(tsSF(10))
}

func TestCalculate(t*testing.T){
	op := func(x, y int) int {
		return x + y
	}
	a,b := 1,4
	res,err := Calculate(a,b,op)
	if err == nil{
		t.Log(res)
	}
	t.Log(err)
}

func TestGenCalculator(t*testing.T){
	op := func(x, y int) int {
		return x + y
	}
	x,y := 12,23
	add := GenCalculator(op)
	res,err := add(x,y)
	if err == nil{
		t.Logf("The result:%d(err:%v)\n",res,err)
	}
}


func TestVarParam(t*testing.T){
	t.Log(Sum(1,2,3,4))
	t.Log(Sum(1,2,3,4,5))
}


func TestDefer(t*testing.T){
	defer Clear()
	fmt.Println("start")
	panic("err")  //panic下依然会执行defer修饰的函数
	//fmt.Println("End") // 执行不到的代码
}