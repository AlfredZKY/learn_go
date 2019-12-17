package usecontext

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func addNum(numP *int32, id int, deferFunc func()) {
	defer func() {
		deferFunc()
	}()

	for i := 0; ; i++ {
		curNum := atomic.LoadInt32(numP)
		newNum := curNum + 1
		time.Sleep(time.Millisecond * 200)
		if atomic.CompareAndSwapInt32(numP, curNum, newNum) {
			fmt.Printf("The number:%d [%d-%d]\n", newNum, id, i)
			break
		} else {
			fmt.Printf("The CAS operation failed. [%d-%d]\n", id, i)
		}
	}
}

func TestCoordinateWithWaitGroup(t *testing.T) {
	total := 12
	stride := 3
	var num int32
	fmt.Printf("The number: %d [with sync.WaitGroup]\n", num)
	var wg sync.WaitGroup
	fmt.Printf("start loop...\n")
	for i := 1; i <= total; i = i + stride {
		wg.Add(stride)
		for j := 0; j < stride; j++ {
			go addNum(&num, i+j, wg.Done)
		}
		wg.Wait()
	}
	fmt.Printf("End.")
}

var (
	wg sync.WaitGroup
)

func work(ctx context.Context) error {
	defer wg.Done()

	for i := 0; i < 1000; i++ {
		select {
		case <-time.After(time.Second * 10):
			fmt.Println("Doing some work ", i)

		// we received the signal of cancelation in this channel
		case <-ctx.Done():
			fmt.Println("Cancel the context ", i)
			fmt.Println(ctx.Err())
			return ctx.Err()
		}
	}
	return nil
}

func TestWithTimeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*4)
	defer cancel()

	fmt.Println("Hey,I'm going to do some work")
	wg.Add(1)
	go work(ctx)
	wg.Wait()
	fmt.Println("Finished. I'm going home")
}

func TestWithDeadline(t *testing.T) {
	d := time.Now().Add(1 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()
	select {
	case <-time.After(2 * time.Second):
		fmt.Println("oversleep")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}

func handle(ctx context.Context, duration time.Duration) {
	select {
	case <-ctx.Done():
		fmt.Println("handle", ctx.Err())
	case <-time.After(duration):
		fmt.Println("Process request with", duration)
	}
}

func TestWithTimeout1(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()

	go handle(ctx, 1500*time.Millisecond)

	select {
		// ctx.Done() 返回的节点管道关闭而终止，也就是上下文超时
		// 多个goroutine同时订阅ctx.Done()管道中的消息，
		// 一旦接收到取消信号就停止当前正在执行的工作并提前返回
	case <-ctx.Done():
		fmt.Println("main", ctx.Err())
	}
}

// 传值很少用到 