package learnfunc

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Printer 函数类型的声明
type Printer func(contents string) (n int, err error)

// PrintToStd Print Str
func PrintToStd(contents string) (bytesNum int, err error) {
	return fmt.Println(contents)
}

type operate func(x int, y int) int

// Calculate 方案1
func Calculate(x int, y int, op operate) (int, error) {
	if op == nil {
		return 0, errors.New("Invalid operation")
	}
	return op(x, y), nil
}

// CalculateFunc func
type CalculateFunc func(x int, y int) (int, error)

// GenCalculator 方案2
func GenCalculator(op operate) CalculateFunc {
	return func(x int, y int) (int, error) {
		if op == nil {
			return 0, errors.New("invalid operation")
		}
		return op(x, y), nil
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
func Clear() {
	fmt.Println("Clear resources.")
}

// Sum func
func Sum(ops ...int) int {
	res := 0
	for _, op := range ops {
		res += op
	}
	return res
}

// Handler 处理函数接口
type Handler interface {
	// Do为接口函数，要有具体的实现操作才行，否则就是空接口
	Do(k, v interface{})
}

// HandlerFunc k v两个接口类型 优化了接口类型 既可以是接口又可以是方法
type HandlerFunc func(k, v interface{})

// Do 具体的Do实现函数
func (f HandlerFunc) Do(k, v interface{}) {
	f(k, v)
}

// Each 具体的执行函数
func Each(m map[interface{}]interface{}, h Handler) {
	if m != nil && len(m) > 0 {
		for k, v := range m {
			h.Do(k, v)
		}
	}
}

// EachFunc 执行Each函数
func EachFunc(m map[interface{}]interface{}, f func(k, v interface{})) {
	Each(m, HandlerFunc(f))
}

// PrintInfo 信息展示 接口函数的具体实现，相当于做了映射关系，不仅仅局限于当前接口函数
func PrintInfo(k, v interface{}) {
	fmt.Printf("大家好，我叫%s,今年%d岁\n", k, v)
}
