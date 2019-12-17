package mymap


import (
	"fmt"
)

// Go 语言规范规定，在键类型的值之间必须可以施加操作符==和!=。
// 换句话说，键类型的值必须要支持判等操作。由于函数类型、字典类型和切片类型的值并不支持判等操作，所以字典的键类型不能是这些类型。

// func OperatorMap
func OperatorMap(){
	aMap := map[string]int{
	"one":    1,
	"two":    2,
	"three": 3,
	//"four":nil,
	}
	k := "two"
	v, ok := aMap[k]
	if ok {
		fmt.Printf("The element of key %q: %d\n", k, v)
	} else {
		fmt.Println("Not found!")
	}

	// map 的键不可以是接口类型
	// var badMap2 = map[interface{}]int{
	// 	"1":   1,
	// 	[]int{2}: 2, // 这里会引发 panic。
	// 	3:    3,
	// }
	// k1 := "two"
	// v1, ok := badMap2[k1]
	// if ok {
	// 	fmt.Printf("The element of key %q: %d\n", k1, v1)
	// } else {
	// 	fmt.Println("Not found!")
	// }
}