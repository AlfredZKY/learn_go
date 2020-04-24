package mygo


import (
	"fmt"
	"sync/atomic"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	time.Sleep(1 * time.Second)
}

func TestGoRoutineProcess(t *testing.T) {
	// 多次运行会打印出不同显示，主要是有系统调度造成
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i)
		}()
	}

	// go func() {
	// 	fmt.Println("My first goroutine")
	// }()
	fmt.Println("hello")
}

func TestGoRoutineBlocked(t *testing.T) {
	num := 100
	sign := make(chan struct{}, num)
	for i := 0; i < num; i++ {
		go func() {
			t.Log(i)
			sign <- struct{}{}
		}()
	}

	for j := 0; j < num; j++ {
		<-sign
	}
}

func TestOrderExecute(t *testing.T) {
	for i := uint32(0); i < 10; i++ {
		go func(i uint32) {
			//t.Log(i)
			fmt.Println(i)
		}(i)
	}
}

func TestSyncRoutine(t *testing.T) {
	var count uint32
	// func() 即无参数声明，也无结果声明的函数类型
	trigger := func(i uint32, fn func()) {
		for {
			// 不停的获取一个名叫count的变量的值，并判断该值是否与参数i相等，如果相等就调用fn
			// 然后把count的+1，最后显式的退出当前循环
			if n := atomic.LoadUint32(&count); n == i {
				fn()
				atomic.AddUint32(&count, 1)
				break
			}
		}
	}
	for i := uint32(0); i < 10; i++ {
		go func(i uint32) {
			fn := func() {
				fmt.Println(i)
			}
			trigger(i, fn)
		}(i)
	}
	// 阻塞等待其他的goroutine执行完毕后，再结束主goroutine
	trigger(10, func() {})
}
