package typeconversion

import (
	"fmt"
	"testing"
)

func TestConvertType(t *testing.T) {
	var general interface{}
	general = 6.6
	typeCast(general)
}

func typeCast(param interface{}) {
	switch param.(type) {
	case int:
		fmt.Println("the param type is int")
		newInt, ok := param.(int)
		if ok == false{
			panic("type cast failed")
		}
		fmt.Println("newInt value:",newInt)
		newInt += 2
		fmt.Println("+2 is",newInt)

		newInt -= 4
		fmt.Println("-4 is",newInt)
		fmt.Println()
		fmt.Println()
	case float32:
		fmt.Println("param type is float32")
		newfloat32, ok := param.(float32)
		if ok == false{
			panic("type cast failed")
		}
		fmt.Println("newfloat32 value:",newfloat32)
		newfloat32 += 2.2
		fmt.Println("+2.2 is",newfloat32)

		newfloat32 -= 4.1
		fmt.Println("-4 is",newfloat32)
		fmt.Println()
		fmt.Println()
	case float64:
		fmt.Println("param type is float32")
		newfloat64, ok := param.(float64)
		if ok == false{
			panic("type cast failed")
		}
		fmt.Println("newfloat32 value:",newfloat64)
		newfloat64 += 2.21
		fmt.Println("+2.21 is",newfloat64)

		newfloat64 -= 4.11
		fmt.Println("-4.11 is",newfloat64)
		fmt.Println()
		fmt.Println()
	default:
		fmt.Println("unkown type")
	}
}

func TestConvertTypes(t*testing.T){
	// 当类型不兼容时，是无法转换的
	var var1 int = 7
	t.Logf("%T->%v\n",var1,var1)

	var2 := float32(var1)
	var3 := int64(var1)
	t.Logf("%T->%v\n",var2,var2)
	t.Logf("%T->%v\n",var3,var3)

	// 类型不兼容
	// var4 := []int(var1)
	// var5 := []int64(var1)
	// t.Logf("%T->%v\n",var4,var4)
	// t.Logf("%T->%v\n",var5,var5)

	var6 := new(int32)
	t.Logf("%T->%v\n",var6,var6)
	var7 := (*int32)(var6)
	t.Logf("%T->%v\n",var7,var7)
}
