package mygoroutine

import (
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
