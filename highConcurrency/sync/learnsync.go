package main

import (
	"sync/atomic"
	"fmt"
	"runtime"
	"sync"
)

func test(c chan int) {
	c <- 'A'
}

func testDeadLock(c chan int) {
	for {
		fmt.Println(<-c)
	}
}

var (
	count int32
	wg    sync.WaitGroup
)

func main() {
	// c := make(chan int)
	// go test(c)
	// go testDeadLock(c)

	// <- c
	// <- c

	wg.Add(2)
	go incCountAutomic()
	go incCountAutomic()
	wg.Wait()
	fmt.Println(count)
}

func incCount() {
	defer wg.Done()
	for i := 0; i < 2; i++ {
		value := atomic.LoadInt32(&count)
		// 是让当前goroutine暂停的意思，退回执行队列，让其他goroutine运行
		runtime.Gosched()
		value++
		atomic.StoreInt32(&count,value)
	}
}

func incCountAutomic() {
	defer wg.Done()
	for i := 0; i < 2; i++ {
		value := count
		// 是让当前goroutine暂停的意思，退回执行队列，让其他goroutine运行
		runtime.Gosched()
		value++
		count = value
	}
}
