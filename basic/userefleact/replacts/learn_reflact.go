package replacts

import (
	"fmt"
	"reflect"
)

// Enum 定义一个Enum类型
type Enum int

const (
	// Zero constant variable
	Zero Enum = 0
)

// Student 学生信息
type Student struct {
	Name string
	Age  int
}

// StudentReflact 操作结构体
func StudentReflact() {
	var stu Student
	typeOfStu := reflect.TypeOf(stu)
	// 打印出变量的名字 变量的字面量
	fmt.Println(typeOfStu.Name(), typeOfStu.Kind())

	// *reflect.rtype *reflect.rtype replacts.Student
	fmt.Println(reflect.TypeOf(stu), reflect.ValueOf(stu))

	// 通过 reflect.TypeOf() 直接获取反射类型对象。
	typeOfZero := reflect.TypeOf(Zero)
	fmt.Println(typeOfZero.Name(), typeOfZero.Kind())

	var pstu = &Student{Name: "zky", Age: 20}
	ptypeOfStu := reflect.TypeOf(pstu)
	fmt.Println(ptypeOfStu)
	// 在go中指针变量的类型名称是空
	fmt.Printf("Name:'%v',Kind:'%v'\n", ptypeOfStu.Name(), ptypeOfStu.Kind())

	// 取类型的元素的指针
	ptypeOfStu = ptypeOfStu.Elem()
	// 显示反射类型对象名称和种类
	fmt.Printf("element name:'%v' , element kind:'%v'\n", typeOfStu.Name(), typeOfStu.Kind())
}

// ReflactStruct reflact 调整结构体
func ReflactStruct() {
	type cat struct {
		Name string
		Type int `json:"type" id:"100"`
	}

	// 创建cat实例
	ins := cat{Name: "mini", Type: 1}

	// 获取结构体实例的反射类型对象
	typeOfCat := reflect.TypeOf(ins)

	fmt.Printf("struct reflact '%v' kind:'%v'", typeOfCat.Name(), typeOfCat.Kind())

	// 遍历结构体所有成员 NumField获得一个结构体共有多少字段，如果不是结构体类型，会触发宕机错误
	// panic: reflect: NumField of non-struct type int [recovered]
	for i := 0; i < typeOfCat.NumField(); i++ {
		// 获取每个成员的结构体字段类型
		filedType := typeOfCat.Field(i)

		// 输出成员名和tag
		fmt.Printf("name: %v tag:%v\n", filedType.Name, filedType.Tag)
	}

	// 通过字段名，找到字段类型信息
	if catType, ok := typeOfCat.FieldByName("Type"); ok {
		// 从tag中中取出需要的tag
		fmt.Println(catType.Tag.Get("json"), catType.Tag.Get("id"))
	}
}

//ReflectVaule 通过反射操作值
func ReflectVaule() {
	// 声明并赋值一个变量
	var a int = 1024

	// 获取变量反射的值
	valueOfA := reflect.ValueOf(a)

	// 获取interface{}类型的值，通过类型断言转换
	var getA int = valueOfA.Interface().(int)

	// 获取64位的值，强制类型转换为int类型
	var getB int = int(valueOfA.Int())

	// 获取bool的值 强制类型转换为bool类型
	var getBool bool = bool(valueOfA.Bool())
	fmt.Println(getA,getB,getBool)
}
