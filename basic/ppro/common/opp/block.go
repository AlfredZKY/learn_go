package opp

import (
	"math/rand"
	"sync"
	"time"
)

// BlockProFile 块阻塞
func BlockProFile() error {
	max := 10
	senderNum := max / 2
	receiverNum := max / 4
	ch1 := make(chan int, max/4)

	var senderGroup sync.WaitGroup
	senderGroup.Add(senderNum)
	repeat := 50000
	for j := 0; j < senderNum; j++ {
		go Sender(ch1, &senderGroup, repeat)
	}

	go func() {
		senderGroup.Wait()
		close(ch1)
	}()

	var receiverGroup sync.WaitGroup
	receiverGroup.Add(receiverNum)
	for j := 0; j < receiverNum; j++ {
		go Receiver(ch1, &receiverGroup)
	}

	receiverGroup.Wait()
	return nil
}

// Sender 发送消息
func Sender(ch1 chan int, wg *sync.WaitGroup, repeat int) {
	defer wg.Done()
	time.Sleep(time.Microsecond * 10)
	for k := 0; k < repeat; k++ {
		elem := rand.Intn(repeat)
		ch1 <- elem
	}
}

// Receiver 接收消息
func Receiver(ch1 chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for elem := range ch1 {
		_ = elem
	}
}
