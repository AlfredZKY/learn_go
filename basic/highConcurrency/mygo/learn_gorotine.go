package mygo

import (
	"sync/atomic"
	"fmt"
	//"runtime"
	"time"
)

func println() {
	fmt.Println("hello world!!!")
}

func usemain() {
	names := []string{"Eric", "Harry", "Robet", "Jim", "Mark"}
	for _, name := range names {
		go func(who string) {
			fmt.Printf("Hello, %s!\n", who)
		}(name) // 带上参数就是可重入函数
	}
	// 并发运行，不知道谁先执行结束
	// go println()

	// name := "Eric"
	// go func() {
	// 	fmt.Printf("Hello,%s!\n", name)
	// }()

	// 等待函数阻塞一下 位置很主要
	time.Sleep(time.Millisecond)
	// name = "Harry"

	//fmt.Println(runtime.GOMAXPROCS)
}

// AsynExecute goroutine is aync execute
func AsynExecute(){
	// 多次运行会打印出不同显示，主要是有系统调度造成，主goroutine执行完毕后，就退出了，不会有打印
	// 由于go异步并发执行，go语句执行完毕后，go程序不会等待go函数的执行，它会立即执行后面的语句
	for i := 0; i < 10; i++ {
		go func(a int) {
			fmt.Println(a)
		}(i)
	}

	// go func() {
	// 	fmt.Println("My first goroutine")
	// }()
	// fmt.Println("hello")
}

// GorotineChannel 通过channel闲置goroutine的数量，限制并发的数量,但是不能保证goroutine按制定的顺序执行
func GorotineChannel(){
	num := 10
	sign := make(chan struct{}, num)
	for i := 0; i < num; i++ {
		go func(a int) {
			fmt.Println(a)
			// 把一个空结构体传入通道
			sign <- struct{}{}
		}(i)
	}

	for j := 0; j < num; j++ {
		<-sign
	}
}


// GorotineOrderExecute 让gorotinr按照制定的顺序执行
func GorotineOrderExecute(){
	
	// 不停的获取一个名叫count的变量的值，并判断该值是否与参数i相等，如果相等就调用fn
	// 然后把count的+1，最后显式的退出当前循环

	// 会被多个goroutine调用，产生竟态
	var count uint32

	// func() 即无参数声明，也无结果声明的函数类型
	trigger := func(i uint32,fn func()){
		for{

			// 自选锁，spinning 除非自己满足条件，否则一直检查
			if n:= atomic.LoadUint32(&count);n == i{
				fn()
				atomic.AddUint32(&count,1)
				break
			}
			time.Sleep(time.Nanosecond)
		}
		
	}

	for i:= uint32(0);i< 10;i++{
		go func(a uint32){
			fn := func(){
				fmt.Println(a)
			}
			trigger(a,fn)
		}(i)
	}

	// 阻塞等待其他的goroutine执行完毕后，再结束主goroutine
	trigger(10,func(){})
}
