package useatomic

import (
	"errors"
	"fmt"
	"io"
	"os"
	"reflect"
	"sync/atomic"
	"testing"
)

type atomicValue struct {
	v atomic.Value
	t reflect.Type
}

// NewAtomicValue 新申请一个原子值 可以存储任意类型的值 指针
func NewAtomicValue(example interface{}) (*atomicValue, error) {
	if example == nil {
		return nil, errors.New("atomic value:nil example")
	}

	return &atomicValue{
		t: reflect.TypeOf(example),
	}, nil
}

func (av *atomicValue) Store(v interface{}) error {
	// 存储值 会产生一个完全分离的新值，相当于被复制那个值的快照，两者是相互不影响的
	// 1.不能用原子值存储nil,否则就会引发panic
	// 2.向原子值存储的第一个值，决定了今后只能存储这个类型的值
	if v == nil {
		return errors.New("atomic value:nil value")
	}

	t := reflect.TypeOf(v)
	if t != av.t {
		return fmt.Errorf("atomic value:wrong type:%s", t)
	}

	av.v.Store(v)
	return nil
}

func (av *atomicValue) Load() interface{} {
	return av.v.Load()
}

func (av *atomicValue) TypeOfValue() reflect.Type {
	return av.t
}

func TestAtomicOther(t *testing.T) {
	// 实例1 原子值在使用前可以被复制
	// atomic.Value可以存储任何值，但是第一次存储的类型确定后，只能存储该类型的值
	var box atomic.Value
	fmt.Println("Copy box to box2")
	box2 := box
	box3 := box2
	v1 := [...]int{1, 2, 3}
	fmt.Printf("Store %v to box.\n", v1)
	box.Store(v1)
	fmt.Printf("The value load from box is %v.\n", box.Load())
	fmt.Printf("The value load from box2 is %v.\n", box2.Load())
	fmt.Println()

	// 实例2
	v2 := "123"
	fmt.Printf("Store %q to box2.\n", v2)
	box2.Store(v2) // 这里并不会引发panic
	fmt.Printf("The value load from box is %v.\n", box.Load())
	fmt.Printf("The value load from box2 is %v.\n", box2.Load())

	// 实例3
	fmt.Println("Copy box to box3")
	// box3 = box //原子值在真正使用后不应该被复制
	fmt.Printf("The value load from box3 is %v.\n", box3.Load())
	v3 := 123
	fmt.Printf("Store %d to box2.\n", v3)
	box3.Store(v3)	// 引发一个panic 此时atomic.Value的类型改变了

	
}

func TestAtomicOther2(t*testing.T){
	// 实例4
	var box4 atomic.Value
	v4 := errors.New("something wrong")
	fmt.Printf("Store an error with message %q to box4.\n", v4)
	box4.Store(v4)
	v41 := io.EOF
	fmt.Println("Store a value of the same type to box4")
	box4.Store(v41)

	// 判断某个变量的类型
	v42, ok := interface{}(&os.PathError{}).(error)
	fmt.Printf("v42 is %v and ok is %v\n", v42, ok)
	if ok {
		fmt.Printf("Store a value of type %T that implements error interface to box4\n", v42)
		// box4.Store(v42) // 此处类型不一致导致panic
	}
	fmt.Println()

	// 实例5
	box5, err := NewAtomicValue(v4)
	if err != nil {
		fmt.Printf("error:%s\n", err)
	}

	fmt.Printf("The legal type in box5 is %s.\n", box5.TypeOfValue())
	fmt.Println("Store a value of the same type to box5")
	err = box5.Store(v42)
	if err != nil {
		fmt.Printf("error:%s\n", err)
	}
	fmt.Println()

	// 实例6
	var box6 atomic.Value
	v6 := []int{1, 2, 3}
	fmt.Printf("Store %v to box6.\n", v6)
	box6.Store(v6)
	v6[1] = 4 // 注意，此处的操作并不是并发安全的
	fmt.Printf("The value load from box6 is %v.\n", box6.Load())

	// 正确的做法下
	v6 = []int{1, 2, 3}
	store := func(v []int) {
		replica := make([]int, len(v))
		copy(replica, v)
		box6.Store(replica)
	}

	fmt.Printf("Store %v to box6.\n", v6)
	store(v6)
	v6[2] = 5 //此处是安全的
	fmt.Printf("The value load from box6 is %v.\n", box6.Load())
	fmt.Printf("v6 is %v.\n", v6)
}
