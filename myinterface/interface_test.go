package myinterface

import "testing"

type Programer interface{
	WriteHelloWorld() string
}

type GoProgrammer struct{
}

func (g * GoProgrammer) WriteHelloWorld()string{
	return "fmt.Println(\"Hello World\")"
}

func TestClient(t*testing.T){
	var p Programer
	p = new(GoProgrammer)
	t.Log(p.WriteHelloWorld())

}

type Rect struct{
	x,y float64
	width,height float64
}

func TestNewObject(t*testing.T){
	rect1 := new(Rect)
	rect2 := &Rect{}
	rect3 := &Rect{0,0,100,200}
	rect4 := &Rect{width:100,height:200}
	// 以上变量全部指向Rect结构体的指针(指针变量),因为使用了new()函数和&操作符
	t.Log(rect1)
	t.Log(rect2)
	t.Log(rect3)
	t.Log(rect4)

	// a变量则和上面完全不一样，它是一个Rect对象的变量
	a := Rect{}
	a.x = 10
	t.Logf("%[1]v %[1]T", a)

	rect5 := &Rect{0,0,100,200}
	rect5.height = 20
	t.Logf("%[1]v %[1]T", rect5)
}