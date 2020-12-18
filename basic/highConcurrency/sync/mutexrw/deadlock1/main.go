package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

var (
	rwMutex sync.RWMutex
)

func ReadGo(idx int, in chan int) {
	for {
		time.Sleep(time.Millisecond * 500)
		rwMutex.RLock()
		num := <-in
		fmt.Printf("%dth 读go程，读到:%d\n", idx, num)
		rwMutex.RUnlock()
	}
}

func WriteGo(idx int, out chan int) {
	for {
		num := rand.Intn(500)
		rwMutex.Lock()
		out <- num
		fmt.Printf("------%dth 写go程，写入:%d\n", idx, num)
		rwMutex.Unlock()
		time.Sleep(time.Millisecond * 200)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	ch := make(chan int, 5)

	for i := 0; i < 5; i++ {
		go ReadGo(i+1, ch)
	}

	for i := 0; i < 5; i++ {
		go WriteGo(i+1, ch)
	}

	for {
		runtime.GC()
	}
}
