package mutexrw

import (
	"log"
	"sync"
	"testing"
	"time"
)

// counter 代表计数器
type counter struct {
	num uint         //计数
	mu  sync.RWMutex //读写锁
}

// number 会返回当前的计数
func (c *counter) number() uint {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.num
}

func (c *counter) add(increment uint) uint {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.num += increment
	return c.num 
}

func count(c *counter) {
	// sign 用于传递信号
	sign := make(chan struct{}, 3)

	// 用于增加计数的函数
	go func() {
		defer func() {
			sign <- struct{}{}
		}()
		for i := 0; i <= 10; i++ {
			time.Sleep(time.Millisecond * 600)
			c.add(1)
		}
	}()

	go func() {
		defer func() {
			sign <- struct{}{}
		}()

		for j := 1; j <= 20; j++ {
			time.Sleep(time.Millisecond * 300)
			log.Printf("The number in counter:%d [%d-%d]", c.number(), 1, j)
		}
	}()
	
	go func () {
		defer func () {
			sign <- struct {}{}
		}()
		for k:=1;k<=20;k++{
			time.Sleep(time.Millisecond *300)
			log.Printf("The number is counter:%d [%d-%d]",c.number(), 2, k)
		}
	}()

	<-sign
	<-sign
	<-sign
}

func redundantUnlock(){
	var rwMu sync.RWMutex

	// 实例1 解锁未锁定的互斥锁会立即引发 panic。会引发panic
	// rwMu.Unlock()

	// 实例2 原因同上
	// rwMu.RUnlock()

	//实例3 
	rwMu.RLock()
	// rwMu.RUnlock() 重复解锁
	rwMu.RUnlock()

	// 实例4
	rwMu.Lock()
	// rwMu.Unlock()
	rwMu.Unlock()
}

func TestMuntexRW(t *testing.T) {
	c := counter{}
	count(&c)
	redundantUnlock()
}
