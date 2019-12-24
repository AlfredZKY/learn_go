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
