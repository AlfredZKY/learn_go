package operator

import (
	"testing"
)

func TestForStatement(t *testing.T) {
	number := []int{1, 2, 3, 4, 5, 6}
	for i := range number {
		if i == 3 {
			number[i] |= i
		}
	}
	t.Log(number)
}

// range表达式只会在for语句开始执行时被求值一次，无论后边会有多少次迭代；
// range表达式的求值结果会被复制，也就是说，被迭代的对象是range表达式结果值的副本而不是原值。

func TestForStatement1(t *testing.T) {
	number := [...]int{1, 2, 3, 4, 5, 6}
	maxIndex := len(number) -1
	t.Log(number)
	// range 语句只初始化一次，不在改变
	for i,e := range number{
		if i == maxIndex{
			number[0] +=e
		}else{
			number[i+1] +=e
		}
		t.Log(number)
	}

	number1 := []int{1, 2, 3, 4, 5, 6}
	maxIndex1 := len(number1) -1
	t.Log(number1)
	// 切片是引用
	for i,e := range number1{
		if i == maxIndex1{
			number1[0] +=e
		}else{
			number1[i+1] +=e
		}
		t.Log(number1)
	}

}
