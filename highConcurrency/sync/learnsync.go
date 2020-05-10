package main

import (
	"fmt"
)

func test(c chan int) {
	c <- 'A'
}

func testDeadLock(c chan int) {
	for {
		fmt.Println(<-c)
	}
}

func main() {
	c := make(chan int)
	go test(c)
	go testDeadLock(c)
	<- c
	<- c
}
