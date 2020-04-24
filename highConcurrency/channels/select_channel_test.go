package channels

import (
	"fmt"
	"testing"
	"time"
)

var numbers = []int{1, 2, 3, 4, 5}
var intChan1 chan int // 未初始化，会阻塞
var intChan2 chan int // 同上
var channels = []chan int{intChan1, intChan2}

func TestSelect(t *testing.T) {
	select {
	case getChan(0) <- getNumber(0):
		fmt.Println("1th case is selected.")
	case getChan(1) <- getNumber(1):
		fmt.Println("2nd case is selected.")
	default:
		fmt.Println("Default")
	}
}

func getChan(i int) chan int {
	fmt.Printf("channels[%d]\n", i)
	return channels[i]
}

func getNumber(i int) int {
	fmt.Printf("numbers[%d]\n", i)
	return numbers[i]
}

func TestSortSelect(t *testing.T) {
	chanCap := 5
	intChan := make(chan int, chanCap)
	for i := 0; i < chanCap; i++ {
		select {
		case intChan <- 1:
		case intChan <- 2:
		case intChan <- 3:
		}
	}
	for i := 0; i < chanCap; i++ {
		fmt.Printf("%d\n", <-intChan)
	}
}

func TestSelectFor(t *testing.T) {
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}
	close(intChan)
	syncChan := make(chan struct{}, 1)
	go func() {
	Loop:
		for {
			select {
			case e, ok := <-intChan:
				if !ok {
					fmt.Println("End.")
					break Loop
				}
				fmt.Printf("Received: %v\n", e)
			}
		}
		syncChan <- struct{}{}
	}()
	<-syncChan
}

func TestChannelCap(t *testing.T) {
	sendingInterval := time.Second
	receptionInterval := time.Second * 2
	intChan := make(chan int, 0)
	go func() {
		var ts0, ts1 int64
		for i := 1; i <= 5; i++ {
			intChan <- i
			ts1 = time.Now().Unix()
			if ts0 == 0 {
				fmt.Println("Sent:", i)
			} else {
				fmt.Printf("Sent:%d [interval :%d s]\n", i, ts1-ts0)
			}
			ts0 = time.Now().Unix()
			time.Sleep(sendingInterval)
		}
		close(intChan)
	}()
	var ts0, ts1 int64
Loop:
	for {
		select{
		case v,ok:= <- intChan:
			if !ok{
				break Loop
			}
			ts1 = time.Now().Unix()
			if ts0 == 0{
				fmt.Println("Received:",v)
			}else{
				fmt.Printf("Received: %d [interval: %d s]\n",v,ts1-ts0)
			}
			ts0 = time.Now().Unix()
			time.Sleep(receptionInterval)
		}
	}
	fmt.Println("End.")
}
