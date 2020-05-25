package learnfunc

import (
	"fmt"
	"testing"
)

func TestReturnMultiValues(t *testing.T) {
	a, _ := ReturnMultiValues()
	t.Log(a)
	tsSF := timeSpent(slowFun)
	t.Log(tsSF(10))
}

func TestCalculate(t *testing.T) {
	op := func(x, y int) int {
		return x + y
	}
	a, b := 1, 4
	res, err := Calculate(a, b, op)
	if err == nil {
		t.Log(res)
	}
	t.Log(err)
}

func TestGenCalculator(t *testing.T) {
	op := func(x, y int) int {
		return x + y
	}
	x, y := 12, 23
	add := GenCalculator(op)
	res, err := add(x, y)
	if err == nil {
		t.Logf("The result:%d(err:%v)\n", res, err)
	}
}

func TestVarParam(t *testing.T) {
	t.Log(Sum(1, 2, 3, 4))
	t.Log(Sum(1, 2, 3, 4, 5))
}

func TestDefer(t *testing.T) {
	defer Clear()
	fmt.Println("start")
	panic("err") //panic下依然会执行defer修饰的函数
	//fmt.Println("End") // 执行不到的代码
}

func TestMapInterface(t *testing.T) {
	persons := make(map[interface{}]interface{})
	persons["张三"] = 20
	persons["李四"] = 23
	persons["王五"] = 26

	// EachFunc(persons, PrintInfo)
	// HandlerFunc(PrintInfo)不是函数的调用，而是函数的转型，同一类型是可以通过强制转换类型。
	Each(persons, HandlerFunc(PrintInfo))

	s := []int{}
	_= s
}
