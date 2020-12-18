package slice

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

func TestSliceInit(t *testing.T) {
	var s0 []int
	t.Log(len(s0), cap(s0))

	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0))

	s1 := []int{1, 2, 3, 4}
	t.Log(len(s1), cap(s1))

	s2 := make([]int, 3, 5)
	t.Log(len(s2), cap(s2))
	// panic: runtime error: index out of range [recovered]
	// panic: runtime error: index out of range
	// t.Log(s2[0],s2[1],s2[2],s2[3],s2[4])
	t.Log(s2[0], s2[1], s2[2])
	s2 = append(s2, 1)
	t.Log(s2[0], s2[1], s2[2], s2[3])
	t.Log(len(s2), cap(s2))
	s2 = append(s2, 1)
	t.Log(len(s2), cap(s2))
	s2 = append(s2, 1)
	t.Log(len(s2), cap(s2))
}

func TestSliceGrowing(t *testing.T) {
	s := []int{}
	for i := 0; i < 1025; i++ {
		s = append(s, i)
		t.Log(len(s), cap(s))
	}
}

func TestSliceShareMemory(t *testing.T) {
	year := []string{
		"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec",
	}
	Q2 := year[3:6]
	t.Log(Q2, len(Q2), cap(Q2))

	summer := year[5:8]
	t.Log(summer, len(summer), cap(summer))
	summer[0] = "UnKnow"
}

func TestSliceComparing(t *testing.T) {
	a := []int{1, 2, 3, 4}
	b := []int{1, 2, 3, 4}
	// invalid operation: a == b (slice can only be compared to nil)
	_ = a
	_ = b

	// nil 切片
	var nilSlice []int

	// 空切片
	slice := []int{}

	t.Log(reflect.TypeOf(nilSlice))
	t.Log(reflect.TypeOf(slice))

	nilSlice = append(nilSlice, 1)
	slice = append(slice, 1)

	t.Log(nilSlice)
	t.Log(slice)
}

func modify(slice []int) {
	fmt.Printf("%p\n", &slice)
	slice[1] = 10
}

// 切片的底层是三个字段构成的结构类型，所以在函数间以值得方式传递的时候，占用的内存非常小，成本很低，在传递复制切片的时候
// 其底层数组不会被复制，也不会受影响，复制只是复制切片本身，不涉及底层数组
// 在函数间传递切片非常高效，而且不需要传递指针和处理复杂的语法，只需要复制切片，然后根据自己
// 业务修改，最后传递回一个新的切片副本即可，这也是为什么函数间传递参数，使用切片，而不是数组的原因

// 在函数中传入数组指针也可以达到省略内存的目的，但是每次传入时，数组指针的地址都不会改变，所以万一数组的指针指向变了，
// 那么函数中的指针执行也会更改，但是切片不会，它会重新拷贝一份。

func TestFuncSlice(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Printf("%p\n", &slice)
	modify(slice)
	t.Log(slice)
}

func array() [1024]int {
	var x [1024]int
	for i := 0; i < len(x); i++ {
		x[i] = i
	}
	return x
}

func slice() []int {
	x := make([]int, 1024)
	for i := 0; i < len(x); i++ {
		x[i] = i
	}
	return x
}

func BenchmarkArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		array()
	}
}

// 对切片和数组进行性能测试，并且禁用内联和优化，来观察切片的堆上内存分配的情况
// go test -bench . -benchmem -gcflags "-N -l"

// BenchmarkArray-6     521054      2340 ns/op     0 B/op       0 allocs/op
// BenchmarkSlice-6     379929      2969 ns/op     8192 B/op    1 allocs/op
// 代表循环次数 	平均执行时间ns 每次执行堆上分配内存总量是0，分配次数也是0

// 结果并非所有时候都是和用切片代替数组，因为切片底层数组可能会在堆上分配内存，
// 而且小数组在栈上拷贝的消耗也未必比make消耗大

func BenchmarkSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		slice()
	}
}

func TestCreateSlice(t *testing.T) {
	s := make([]byte, 20)
	ptr := unsafe.Pointer(&s[0])
	t.Log(ptr)

	length := len(s)

	var ptr1 unsafe.Pointer
	var s1 = struct {
		addr uintptr
		len  int
		cap  int
	}{(uintptr)(ptr1), length, length}
	sptr := *(*[]byte)(unsafe.Pointer(&s1))
	// var b []byte = []byte("s")

	// sptr = append(sptr, b...)
	t.Log(sptr)

	var o []byte
	sliceHeader := (*reflect.SliceHeader)((unsafe.Pointer(&o)))
	sliceHeader.Cap = length
	sliceHeader.Len = length
	sliceHeader.Data = uintptr(ptr)
	t.Log(sliceHeader)
	o = append(o, []byte("ssss")...)
	o[2] = 1
	t.Log(sliceHeader)
	t.Log(o)
}

func TestSliceCopy(t *testing.T) {
	slice := []int{10, 20, 30, 40}
	// 利用range对切片进行索引时,value的地址其实是切片里面的值的拷贝，所以每次打印出来的地址不变
	// 由于value是值拷贝，并非引用传递，所以每次修改value是达不到更改原切片值得目的的，需要通过&slice[index]获取真实的地址
	for index, value := range slice {
		fmt.Printf("value= %d ,value-addr = %x,slice-addr = %x\n", value, &value, &slice[index])
	}

	type Duration int64
	var dur Duration
	dur = Duration(100)
	t.Log(dur)
}
