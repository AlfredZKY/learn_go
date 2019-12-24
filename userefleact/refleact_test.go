package usereflect

import (
	"fmt"
	"math/rand"
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestReflectFirst(t *testing.T) {
	// 由一个变量得知它所携带的信息
	author := "testing"
	fmt.Println(&author)
	fmt.Println("TypeOf author:", reflect.TypeOf(author))
	fmt.Println("ValueOf author:", reflect.ValueOf(author))

	// 通过反射改变i的值,地址不变
	i := 1
	fmt.Println(&i)
	v := reflect.ValueOf(&i)
	v.Elem().SetInt(10)
	fmt.Println(i)
	fmt.Println(&i)

	s := reflect.ValueOf(&author)
	temp := "hello world"

	// 判断是否可以通过地址来重设新值
	fmt.Println(s.CanSet())

	// s.Elem().Set(reflect.ValueOf(&s))
	s.Elem().SetString(temp)
	fmt.Println(s)
	fmt.Println(&s)
}

func TestReflectKindOrName(t *testing.T) {
	type MyInt int
	var x MyInt = 7
	v := reflect.ValueOf(x)
	_ = v
	// Kind()返回静态类型 Type返回类型和type出的类型名
	fmt.Println(v.Kind(), v.Type())
	temp := 1
	v = reflect.ValueOf(temp)
	fmt.Println(v.Kind(), v.Type())
}

type Student struct {
	Pid    string
	Openid string
	Name   string
	Age    string
}

func (s Student) PrintPid() {
	fmt.Println(s.Age)
}

func (s Student) SetOpenid(id string) {
	s.Openid = id
	fmt.Println(s.Openid)
}

func (s Student) ReAge() string {
	return s.Age
}

func SetRand() string {
	rand.Seed(time.Now().UnixNano())
	answers := []string{
		"It is certain",
		"It is decidedly so",
		"Without a doubt",
		"Yes definitely",
		"You may rely on it",
		"As I see it yes",
		"Most likely",
		"Outlook good",
		"Yes",
		"Signs point to yes",
		"Reply hazy try again",
		"Ask again later",
		"Better not tell you now",
		"Cannot predict now",
		"Concentrate and ask again",
		"Don't count on it",
		"My reply is no",
		"My sources say no",
		"Outlook not so good",
		"Very doubtful",
	}
	time.Sleep(time.Millisecond)
	return answers[rand.Intn(len(answers))]
}

// 使用反射，查看字段并且赋值
func reflectStructField(sPtr interface{}) {
	switch sPtr.(type) {
	case *Student:
		fmt.Println(("*Struct type"))
	default:
		fmt.Println("None type")
	}

	// 反射出一个字段
	rValue := reflect.ValueOf(sPtr)

	// 取出指针值
	rE := rValue.Elem()

	// 判断结构体中字段是否可以更改
	fmt.Println("struct element's is:", rE.NumField())

	for i := 0; i < rE.NumField(); i++ {
		f := rE.Field(i)
		if f.CanSet() {
			f.SetString(SetRand())
		}
	}
}

func reflectStructMethod(sPtr interface{}) {
	if sPtr == nil {
		fmt.Println("sPtr is nil")
		return
	}

	// 结构体的函数名根据ASCII码进行排序
	rValue := reflect.ValueOf(sPtr).Elem()
	for i := 0; i < rValue.NumMethod(); i++ {
		if i == 0 {
			// 无参数
			// fmt.Println(rValue.Method(i))
			rValue.Method(i).Call(nil)
		} else if i == 1 {
			// 有返回值
			data := rValue.Method(i).Call(nil)
			fmt.Println("return is:", data)
		} else if i == 2 {
			// 有参数
			var v []reflect.Value
			v = append(v, reflect.ValueOf("天下无双"))
			rValue.Method(i).Call(v)
		} else {
			fmt.Println("Don't run here")
		}
	}
}

func TestLearnReflectStruct(t *testing.T) {
	fmt.Println("learn reflect...")
	var num float64 = 1.23456
	pointer := reflect.ValueOf(&num)
	value := reflect.ValueOf(num)
	fmt.Println(pointer, value)

	// 防止断言错误，出现panic
	convertPointer, ok := pointer.Interface().(*float64)
	if ok {
		fmt.Println("convertPointer", convertPointer)
	}
	convertValue, ok := pointer.Interface().(float64)
	if ok {
		fmt.Println("convertValue", convertValue)
	}

	convertValue = 90.0
	pointer.Elem().SetFloat(convertValue)
	fmt.Println("num", num)

	// 反射结构体字段的运用
	s := &Student{"a", "b", "c", "d"}
	reflectStructField(s)
	fmt.Println(*s)
	// 反射结构体方法的运用
	reflectStructMethod(s)
}

func SwapTwoValue() {
	swap := func(in []reflect.Value) []reflect.Value {
		return []reflect.Value{in[1], in[0]}
	}

	makeSwap := func(fPtr interface{}) {
		fn := reflect.ValueOf(fPtr).Elem()
		// 函数原型，函数实现
		v := reflect.MakeFunc(fn.Type(), swap)
		fn.Set(v)
	}

	var intSwap func(int, int) (int, int)
	makeSwap(&intSwap)
	fmt.Println(intSwap(0, 1))
}

func AddTwoValue() {
	add := func(args []reflect.Value) (results []reflect.Value) {
		if len(args) == 0 {
			return nil
		}

		var r reflect.Value

		switch args[0].Kind() {
		case reflect.Int:
			n := 0
			for _, a := range args {
				n += int(a.Int())
			}
			r = reflect.ValueOf(n)
		case reflect.String:
			ss := make([]string, 0, len(args))
			for _, s := range args {
				ss = append(ss, s.String())
			}
			r = reflect.ValueOf(strings.Join(ss, ""))
		}
		results = append(results, r)
		return results
	}
	makeAdd := func(fPtr interface{}) {
		fn := reflect.ValueOf(fPtr).Elem()
		// 函数原型，函数实现
		v := reflect.MakeFunc(fn.Type(), add)
		fn.Set(v)
	}
	//var addInt func(int,int)(int)
	var addstring func(string,string)(string)
	makeAdd(&addstring)
	fmt.Println(addstring("10", "1"))
}

func TestLearnReflectFunc1(t *testing.T) {
	SwapTwoValue()
	AddTwoValue()
}
