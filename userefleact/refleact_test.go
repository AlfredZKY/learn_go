package usereflect

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReflectFirst(t *testing.T) {
	// 由一个变量得知它所携带的信息
	author := "testing"
	fmt.Println("TypeOf author:", reflect.TypeOf(author))
	fmt.Println("ValueOf author",reflect.ValueOf(author))

	// 通过反射改变i的值,地址不变
	i := 1
	fmt.Println(&i)
	v := reflect.ValueOf(&i)
	v.Elem().SetInt(10)
	fmt.Println(i)
	fmt.Println(&i)
}
