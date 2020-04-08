package mygoroutine

import (
	"fmt"
	"testing"
	"time"
)

var strChan2 = make(chan string, 3)

func reciver(strChan <-chan string, syncChan1 <-chan struct{}, syncChan2 chan<- struct{}) {
	<-syncChan1
	fmt.Println("Recived a sync signal and wait a second... [receiver]")
	time.Sleep(time.Second)
	for {
		if elem, ok := <-strChan; ok {
			fmt.Println("Received:", elem, "[receiver]")
		} else {
			break
		}
	}
	fmt.Println("Stopped. [receiver]")
	syncChan2 <- struct{}{}
}

func send(strChan chan<- string, syncChan1 chan<- struct{}, syncChan2 chan<- struct{}) {
	for _, elem := range []string{"a", "b", "c", "d"} {
		strChan <- elem
		fmt.Println("Sent:", elem, "[sender]")
		if elem == "c" {
			syncChan1 <- struct{}{}
			fmt.Println("Sent a sync signal. [sender]")
		}
	}
	fmt.Println("Wait 2 senonds... [sender]")
	time.Sleep(time.Second)
	close(strChan)
	syncChan2 <- struct{}{}
}

func TestSignalChan(t *testing.T) {
	syncChan1 := make(chan struct{}, 1)
	syncChan2 := make(chan struct{}, 2)
	go reciver(strChan2, syncChan1, syncChan2)
	go send(strChan2, syncChan1, syncChan2)

	<-syncChan2
	<-syncChan2
}

func TestChannelType(t *testing.T) {
	var ok bool
	ch := make(chan int, 1)
	_, ok = interface{}(ch).(<-chan int)
	fmt.Println("chan int => <-chan int:",ok)

	_, ok = interface{}(ch).(chan<- int)
	fmt.Println("chan int => chan<- int:",ok)

	sch := make(chan<- int,1)
	_, ok = interface{}(sch).(chan int)
	fmt.Println("chan int => chan int:",ok)

	rch := make(<-chan int,1)
	_, ok = interface{}(rch).(chan int)
	fmt.Println("chan int => chan int:",ok)
}
