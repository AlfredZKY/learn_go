package useselect

import (
	"fmt"
	"testing"
	"time"
)

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
			fmt.Println(x, y)
		case <-quit:
			fmt.Println("quit", x, y)
			return
		}
	}
}

func TestSelect(t *testing.T) {
	a := make(chan int)
	b := make(chan int)
	b <- 1
	go fibonacci(a, b)
	a <- 1

}

func TestNotBlock(t *testing.T) {
	// 没有阻塞直接返回了
	ch := make(chan int)
	select {
	case i := <-ch:
		fmt.Println(i)
	default:
		fmt.Println("default")
	}
}

func TestRand(t *testing.T) {
	ch := make(chan int)
	go func() {
		for range time.Tick(1 * time.Second) {
			ch <- 0
		}
	}()
	for {
		// 遇到两个同时响应，会随机选取一个执行
		select {
		case <-ch:
			fmt.Println("case1")
		case <-ch:
			fmt.Println("case2")
		}
	}
}
