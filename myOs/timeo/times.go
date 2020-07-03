package main

import (
	"fmt"
	"time"
)

// 创建定时任务
var (
	globalrounds int32
)
func main() {
	// var ch chan int

	// ticker := time.NewTicker(time.Second * 2)
	// go func(tickers *time.Ticker) {
	// 	for range tickers.C {
	// 		fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	// 	}
	// 	ch <- 1
	// }(ticker)
	// <-ch
	c := time.Tick(time.Second * 5)
	for {
		<-c
		go func() {
			fmt.Println(time.Now().Format("2006-01-02 15:04:05"),globalrounds)
			globalrounds += 1
		}()
	}
}
