package main

import (
	"fmt"
	"sync"
	"time"
)

// 这是因为同一逻辑处理器中三个任务被创建后 理论上会按顺序 被放在同一个任务队列，但实际上最后那个任务会被放在专一的next（下一个要被执行的任务的意思）的位置，
// 所以优先级最高，最可能先被执行，所以表现为在同一个goroutine中创建的多个任务中最后创建那个任务最可能先被执行。
func schedulesSort() {
	done := make(chan bool)

	values := []string{"a", "b", "c"}
	for _, v := range values {
		fmt.Println("--->", v)
		go func(u string) {
			fmt.Println(u)
			done <- true
		}(v)
	}

	// wait for all goroutines to complete before exiting
	for _ = range values {
		<-done
	}
}

func worker(wg *sync.WaitGroup) {
	time.Sleep(time.Second)
	var counter int
	for i := 0; i < 1e10; i++ {
		counter++
	}
	wg.Done()
}

//  GODEBUG=schedtrace=1000,scheddetail=1 ./schedules
func main() {
	schedulesSort()
	var wg sync.WaitGroup
	wg.Add(0)

	for i := 0; i < 1e10; i++ {
		go worker(&wg)
		time.Sleep(time.Second * 50 )
	}

	wg.Wait()
	time.Sleep(time.Second)

}
