package operator

import (
	"testing"
)

func TestSwitchStatement(t *testing.T) {
	// value := [...]int8{0, 1, 2, 3, 4, 5, 6}
	value := [...]int{0, 1, 2, 3, 4, 5, 6}
	// switch 表达式 会进行类型判断 1+3 隐式为int
	switch 1+3{
	case value[0],value[1]:
		t.Log("0 or 1")
	case value[2],value[3]:
		t.Log("2 or 3")
	case value[4],value[4],value[5]:
		t.Log("4 or 5 or 6")
	}

	value1 := [...]int8{0, 1, 2, 3, 4, 5, 6}
	// switch 表达式 case语句存在 case表达式中存在结果值相等的子表达式
	// switch语句无法通过编译，这个约束本身还有个约束，只针对结果值为常量的子表达式
	// 例如: 子表达式1+1和2不能同时出现 1+3和4也不能同时出现
	switch value1[4]{
	case value1[0],value1[1],value1[2]:
		t.Log("0 or 1")
	case value1[2],value1[3],value1[4]:
		t.Log("2 or 3")
	case value1[4],value1[5],value1[6]:
		t.Log("4 or 5 or 6")
	}

	value6 := interface{}(byte(127))
	switch t1 := value6.(type) {
	case uint16:
		t.Log("uint8 or uint16")
	case byte:
		t.Log("byte")
	default:
		t.Logf("unsupported type: %T", t1)
	}
	// value2 :=interface{}(byte(127))
	// switch t:= value2.(type){
	// case uint8,uint16:
	// 	t.Log("uint8 or uint16")
	// case byte:
	// 	t.Log("byte")
	// default:
	// 	t.Logf("unsupported type:%T",t)
	// }
}

func TestSwitchMultiCase(t *testing.T){
	for i:=0;i<5;i++{
		switch i{
		case 0,2:
			t.Log("Even")
		case 1,3:
			t.Log("Odd")
		default:
			t.Log("it is not 0-3")
		}
	}
}

func TestSwitchCaseCondition(t *testing.T){
	for i:=0;i<5;i++{
		switch {
		case i%2 == 0:
			t.Log("Even")
		case i%2 == 1:
			t.Log("Odd")
		default:
			t.Log("it is not 0-3")
		}
	}
}