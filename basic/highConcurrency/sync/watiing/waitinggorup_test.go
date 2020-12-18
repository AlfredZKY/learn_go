package waiting

import (
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func addNum(numP *int32, id, max int32, deferFunc func()) {
	defer func() {
		// 不能忘记调用
		deferFunc()
	}()

	for i := 0; ; i++ {
		curNum := atomic.LoadInt32(numP)
		if curNum >= max {
			break
		}

		newNum := curNum + 2
		time.Sleep(time.Millisecond * 200)

		if atomic.CompareAndSwapInt32(numP, curNum, newNum) {
			fmt.Printf("The number: %d [%d-%d]\n", newNum, id, i)
		} else {
			fmt.Printf("the CAS operation failed. [%d-%d]\n", id, i)
		}
	}
}

func TestNoWaitingGroup(t *testing.T) {
	sign := make(chan struct{}, 2)
	num := int32(0)
	t.Logf("The number: %d [with chan struct{}]\n", num)
	max := int32(10)

	go addNum(&num, 1, max, func() {
		sign <- struct{}{}
	})
	go addNum(&num, 2, max, func() {
		sign <- struct{}{}
	})

	<-sign
	<-sign
}

func TestWaitingGroup(t *testing.T) {
	var wg sync.WaitGroup

	wg.Add(2)
	num := int32(0)
	fmt.Printf("the number: %d [with sync.WaitGroup]\n", num)
	max := int32(10)

	go addNum(&num, 1, max, wg.Done)
	go addNum(&num, 2, max, wg.Done)

	// 等待其他goroutine执行完毕 杜塞goroutine
	wg.Wait()
}

func testWaitGroup(t *testing.T, wg1 *sync.WaitGroup, wg2 *sync.WaitGroup) {
	n := 3
	wg1.Add(n)
	wg2.Add(n)
	exited := make(chan bool, n)
	for i := 0; i != n; i++ {
		go func() {
			wg1.Done()
			fmt.Println("wg1 is -- operation.")
			wg2.Wait()
			exited <- true
		}()
	}
	wg1.Wait()
	fmt.Println("wg1 is finished.")
	for i := 0; i != n; i++ {
		select {
		case <-exited:
			t.Fatal("WaitGroup released group too soon")
		default:
		}
		wg2.Done()
		fmt.Println("wg2 is -- operation.")
	}
	for i := 0; i != n; i++ {
		<-exited // Will block if barrier fails to unlock someone.
	}
}

func TestWaitGroup(t *testing.T) {
	wg1 := &sync.WaitGroup{}
	wg2 := &sync.WaitGroup{}

	// Run the same test a few times to ensure barrier is in a proper state.
	for i := 0; i != 100000; i++ {
		testWaitGroup(t, wg1, wg2)
	}
}

func TestWaitGroupMisuse(t *testing.T) {
	defer func() {
		err := recover()
		if err != "sync: negative WaitGroup counter" {
			t.Fatalf("Unexpected panic: %#v", err)
		}
	}()
	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func() {
		wg.Done()
	}()

	go func(){
		wg.Done()
	}()

	t.Fatal("Should panic")
}

func TestOnce(t *testing.T) {
	var counter uint32
	var once sync.Once

	go func() {
		once.Do(func() {
			// for i := 0; i != 10; i++ {
			// 	atomic.AddUint32(&counter, 1)
			// }
			atomic.AddUint32(&counter, 1)
			fmt.Println("enter in 1")
		})
	}()
	fmt.Printf("The counter:%d\n", counter)

	// 该函数没有被执行，只执行第一次传入的函数
	go func() {
		once.Do(func() {
			atomic.AddUint32(&counter, 2)
			fmt.Println("enter in 2")
		})
	}()
	time.Sleep(time.Second * 1)
	fmt.Printf("The counter:%d\n", counter)
	fmt.Println()
}

// go clean -cache 清楚测试的缓存数据
func TestOnce1(t *testing.T) {
	var once sync.Once
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		defer wg.Done()
		once.Do(func() {
			for i := 0; i < 3; i++ {
				fmt.Printf("Do task. [1-%d]\n", i)
				time.Sleep(time.Second)
			}
		})
		fmt.Println("Done. [1]")
	}()

	go func() {
		defer wg.Done()
		time.Sleep(time.Millisecond * 500)
		once.Do(func() {
			for i := 0; i < 3; i++ {
				fmt.Printf("Do task. [2-%d]\n", i)
			}
		})
		fmt.Println("Done. [2]")
	}()

	go func() {
		defer wg.Done()
		time.Sleep(time.Millisecond * 500)
		once.Do(func() {
			for i := 0; i < 3; i++ {
				fmt.Printf("Do task. [3-%d]\n", i)
			}
		})
		fmt.Println("Done. [3]")
	}()

	wg.Wait()
	fmt.Println()
}

// 多次运行或改变条件结果会不一样
func TestOnce2(t *testing.T) {
	var once sync.Once
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		defer func() {
			if p := recover(); p != nil {
				fmt.Printf("fatal error: %v\n", p)
			}
		}()
		once.Do(func() {
			fmt.Println("Do task. [1]")
			panic(errors.New("something wrong"))
			fmt.Println("Done. [1]")
		})
	}()

	go func() {
		defer wg.Done()
		//time.Sleep(time.Millisecond * 1)
		once.Do(func() {
			fmt.Println("Do task. [2]")
		})
		fmt.Println("Done. [2]")
	}()
	wg.Wait()
}

func addNum1(numP *int32, id int, deferFunc func()) {
	defer func() {
		deferFunc()
	}()

	for i := 0; ; i++ {
		curNum := atomic.LoadInt32(numP)
		newNum := curNum + 1
		time.Sleep(time.Millisecond * 200)
		if atomic.CompareAndSwapInt32(numP, curNum, newNum) {
			fmt.Printf("The number:%d [%d-%d]\n", numP, id, i)
			break
		} else {
			fmt.Printf("The CAS operation failed. [%d-%d]\n", id, i)
		}
	}
}

func TestBatchProcess(t *testing.T) {
	total := 12
	stride := 3
	var num int32
	var wg sync.WaitGroup

	fmt.Printf("The number: %d [with sync.WaitGroup]\n", num)
	for i := 1; i <= total; i = i + stride {
		wg.Add(stride)
		for j := 0; j < stride; j++ {
			go addNum1(&num, i+j, wg.Done)
		}
		wg.Wait()
	}
	fmt.Println("End.")
}
