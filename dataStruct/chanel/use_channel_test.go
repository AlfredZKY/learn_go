package mygoroutine

import (
	"fmt"
	"testing"
	"time"
)

type Counter struct {
	count int
}

var strChan = make(chan string, 3)
var mapChan = make(chan map[string]int, 1)
var mapChan1 = make(chan map[string]Counter, 1)
var mapchan2 = make(chan map[string]*Counter, 1)

func useChannel() {
	syncChan1 := make(chan struct{}, 1)
	syncChan2 := make(chan struct{}, 2)

	go func() {
		<-syncChan1
		fmt.Println("Recived a sync signal and wait a second... [receiver]")
		time.Sleep(time.Second)

		for {
			if elem, ok := <-strChan; ok {
				fmt.Println("Recived: ", elem, "[receiver]")
			} else {
				break
			}
		}
		fmt.Println("Stopped. [receiver]")
		syncChan2 <- struct{}{}
	}()

	go func() {
		for _, elem := range []string{"a", "b", "c", "d"} {
			strChan <- elem
			fmt.Println("Sent: ", elem, "[sender]")
			if elem == "c" {
				syncChan1 <- struct{}{}
				fmt.Println("Sent a sync signal. [sender]")
			}
		}
		fmt.Println("Wait 2 seconds... [sender]")
		time.Sleep(time.Second * 2)
		close(strChan)
		syncChan2 <- struct{}{}
	}()
	<-syncChan2
	<-syncChan2
}

func useChanne2() {
	syncChan := make(chan struct{}, 2)
	go func() {
		for {
			if elem, ok := <-mapChan; ok {
				fmt.Println("enter mapChan", ok, elem)
				elem["count"]++
			} else {
				break
			}
		}
		fmt.Println("Stopped. [receiver]")
		syncChan <- struct{}{}
	}()

	go func() {
		countMap := make(map[string]int)
		fmt.Printf("The count map: %v. [sender]\n", countMap)
		for i := 0; i < 5; i++ {
			fmt.Println("enter countMap")
			mapChan <- countMap
			time.Sleep(time.Millisecond)
			fmt.Printf("The count map: %v. [sender]\n", countMap)
		}
		close(mapChan)
		syncChan <- struct{}{}
	}()
	<-syncChan
	<-syncChan
}

func useChanne3() {
	syncChan := make(chan struct{}, 2)
	go func() {
		for {
			if elem, ok := <-mapChan1; ok {
				counter := elem["count"]
				counter.count++
			} else {
				break
			}
		}
		fmt.Println("Stopped. [receiver]")
		syncChan <- struct{}{}
	}()

	go func() {
		countMap := map[string]Counter{
			"count": Counter{},
		}
		for i := 0; i < 5; i++ {
			mapChan1 <- countMap
			time.Sleep(time.Millisecond)
			fmt.Printf("The count map: %v. [sender]\n", countMap)
		}
		close(mapChan1)
		syncChan <- struct{}{}
	}()
	<-syncChan
	<-syncChan
}

func (counter *Counter) String() string {
	return fmt.Sprintf("{count:%d}", counter.count)
}

func useChanne4() {
	syncChan := make(chan struct{}, 2)
	go func() {
		for {
			if elem, ok := <-mapchan2; ok {
				counter := elem["count"]
				counter.count++
			} else {
				break
			}
		}
		fmt.Println("Stopped. [receiver]")
		syncChan <- struct{}{}
	}()

	go func() {
		countMap := map[string]*Counter{
			"count": &Counter{},
		}
		for i := 0; i < 5; i++ {
			mapchan2 <- countMap
			time.Sleep(time.Millisecond)
			fmt.Printf("The count map: %v -%v. [sender]\n", countMap["count"].String(),countMap["count"])
		}
		close(mapchan2)
		syncChan <- struct{}{}
	}()
	<-syncChan
	<-syncChan
}

func TestUseChannel(t *testing.T) {
	useChannel()
}

func TestUseChanne2(t *testing.T) {
	useChanne2()
}

func TestUseChanne3(t *testing.T) {
	useChanne3()
}

func TestUseChanne4(t *testing.T) {
	useChanne4()
}
