package usecontext

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

var (
	wg sync.WaitGroup
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
		fmt.Println(ctx.Err()) // 打印出错误的原因 定时器到期 context deadline exceeded
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

func monitor(ctx context.Context, num int) {
	for {
		select {
		case v := <-ctx.Done(): // 不断的判断 <-ctx.Done() 是否可读，如果可读，则证明该goroutine可以取消
			fmt.Printf("监控器%v,接受的通道值为:%d,监控结束。\n", num, v)
			return
		default:
			fmt.Printf("监控器%v,正在监控中...\n", num)
			time.Sleep(2 * time.Second)
		}
	}
}

func TestWithCancel(t *testing.T) {
	// 以 context.Background()为parent context定义一个可取消的context
	ctx, cancel := context.WithCancel(context.Background())

	for i := 1; i <= 5; i++ {
		go monitor(ctx, i)
	}

	time.Sleep(1 * time.Second)

	// 关闭所有的通道
	cancel()

	time.Sleep(5 * time.Second)
	fmt.Printf("主程序退出")
}

// 传值很少用到
func monitorValue(ctx context.Context, num int) {
	for {
		select {
		case v := <-ctx.Done(): // 不断的判断 <-ctx.Done() 是否可读，如果可读，则证明该goroutine可以取消
			fmt.Printf("监控器%v,接受的通道值为:%d,监控结束。\n", num, v)
			return
		default:
			value := ctx.Value("item")
			fmt.Printf("监控器%v,正在监控 %v...\n", num,value)
			time.Sleep(2 * time.Second)
		}
	}
}

func TestContextWithValue(t*testing.T){
	ctx01,cancel := context.WithCancel(context.Background())
	ctx02,cancel2 := context.WithTimeout(ctx01,1*time.Second)
	ctx03 := context.WithValue(ctx02,"item","cpu")

	defer cancel()
	defer cancel2()

	for i :=1 ; i <=5;i++{
		go monitorValue(ctx03,i)
	}

	time.Sleep(5*time.Second)
	if ctx02.Err() != nil{
		fmt.Println("监控取消的原因: ",ctx02.Err())
	}
	fmt.Printf("主程序退出")
}