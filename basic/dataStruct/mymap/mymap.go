package mymap


import (
	"fmt"
)

// Go 语言规范规定，在键类型的值之间必须可以施加操作符==和!=。
// 换句话说，键类型的值必须要支持判等操作。由于函数类型、字典类型和切片类型的值并不支持判等操作，所以字典的键类型不能是这些类型。

// OperatorMap operator Map
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
	var badMap2 = map[interface{}]int{
		"1":   1,
		[]int{2}: 2, // 这里会引发 panic。
		3:    3,
	}
	k1 := "two"
	v1, ok := badMap2[k1]
	if ok {
		fmt.Printf("The element of key %q: %d\n", k1, v1)
	} else {
		fmt.Println("Not found!")
	}
}

// Mapfuncvalue testing value
func Mapfuncvalue(){
	array := make(map[int]func()int)

	// 可以通过
	array[func()int{return 10}()] = func ()int {
		return 12
	}

	// 类型检查就直接报错了
	// array1 := make(map[func()int]int)
	// // 
	// array[func()int{return 10}] =  12

	fmt.Println(array)
}

// Mapvaluenil testing value nil
func Mapvaluenil(){

	// 声明一个map,下面操作会产生第一个元素对值为nil的map
	var valuenil map[string]int

	// make函数创建的map本身就是一个非nil
	valuenil2 := make(map[string]int)

	key := "two"
	elem,ok:=valuenil["two"]
	fmt.Printf("The element paired with key %q in nil map: %d (%v)\n",
		key, elem, ok)

	fmt.Printf("The length of nil map:%d\n",len(valuenil))

	fmt.Printf("Delete the key-element pair by key%q...\n",key)
	delete(valuenil,key)
	elem = 2
	fmt.Printf("Add a key key-element pair to nil map...")
	valuenil["third"]= elem
	valuenil2["third"]= elem
	fmt.Printf("The length of nil map:%d\n",len(valuenil))
}

// Mapvaluenil1 testing nils
func Mapvaluenil1(){
	var m map[string]int

	key := "two"
	elem, ok := m["two"]
	fmt.Printf("The element paired with key %q in nil map: %d (%v)\n",
		key, elem, ok)

	fmt.Printf("The length of nil map: %d\n",
		len(m))

	fmt.Printf("Delete the key-element pair by key %q...\n",
		key)
	delete(m, key)

	elem = 2
	fmt.Println("Add a key-element pair to a nil map...")
	m["two"] = elem // 这里会引发panic。
}