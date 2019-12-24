package condition

import (
	"log"
	"time"
	"sync"
	"testing"
)

func TestCondition(t*testing.T){
	// mailbox 代表信箱  0代表信箱为空  1代表信箱中有数据
	var mailbox uint8 

	// lock 代表信箱上的锁
	var lock sync.RWMutex
	
	// sendCond 代表专用于发信的条件变量 recvCond代表专用于收信的条件变量
	sendCond :=sync.NewCond(&lock)
	recvCond := sync.NewCond(lock.RLocker()) 	//底层本来就是指针

	// sign 用于传递显示完成的信号
	sign := make(chan struct {},3)
	max := 5

	// 发信
	go func(max int){
		defer func(){
			sign <- struct {}{}
		}()
		for i:=0; i < max; i++{
			time.Sleep(time.Millisecond*500)
			lock.Lock()  // 不是锁上锁，而是有打开锁的权利
			for mailbox == 1{
				// 等待被唤醒
				sendCond.Wait()
			}
			log.Printf("sender [%d]: the mailbox is empty.",i)
			mailbox = 1
			log.Printf("sender [%d]: the letter has been sent.",i)
			lock.Unlock()

			// 主动单向唤醒
			recvCond.Signal()
		}
	}(max)

	// 收信
	go func (max int) {
		defer func(){
			sign <- struct {}{}
		}()

		for j := 0; j < max; j++{
			time.Sleep(time.Millisecond*500)
			lock.RLock()
			for mailbox == 0 {
				// 等待被唤醒
				recvCond.Wait()
			}
			log.Printf("receiver [%d]:the mailbox is full",j)
			mailbox=0
			log.Printf("receiver [%d]:the letter has been received.",j)
			lock.RUnlock()

			// 主动单向唤醒
			sendCond.Signal()
		}
	}(max)
	
	<- sign
	<- sign
}