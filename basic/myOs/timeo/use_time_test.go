package timeo

import (
	"fmt"
	"learn_go/logger"
	"testing"
	"time"
)

func TestTimer1(t *testing.T) {
	timer := time.NewTimer(10 * time.Second)
	fmt.Printf("Present time: %v.\n", time.Now().Format("2006-01-02 15:04:05"))
	expirationTime := <-timer.C
	fmt.Printf("Expiration time: %v.\n", expirationTime)
	fmt.Printf("Stop timer:%v\n", timer.Stop())
}

func TestTimeOut(t *testing.T) {
	intChan := make(chan int, 1)
	go func() {
		time.Sleep(time.Second)
		intChan <- 1
	}()
	select {
	case e := <-intChan:
		fmt.Printf("Received: %v\n", e)
	case <-time.NewTimer(time.Millisecond * 500).C:
		fmt.Println("Timeout!!!")
	case <-time.After(time.Millisecond * 500):
		fmt.Println("Timeout!!!")
	}
}

func TestComplexTimer(t *testing.T) {
	intChan := make(chan int, 1)
	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(time.Second)
			intChan <- i
		}
		close(intChan)
	}()

	timeout := time.Millisecond * 500
	var timer *time.Timer
	for {
		if timer == nil {
			timer = time.NewTimer(timeout)
		} else {
			timer.Reset(timeout)
		}
		select {
		case e, ok := <-intChan:
			if !ok {
				fmt.Println("End.")
				return
			}
			fmt.Printf("Received: %v\n", e)
		case <-timer.C:
			fmt.Println("Timeout!!!")
		}
	}
}

func TestTicker(t *testing.T) {
	intChan := make(chan int, 1)
	ticker := time.NewTicker(time.Second)
	go func() {
		for _ = range ticker.C {
			select {
			case intChan <- 1:
			case intChan <- 2:
			case intChan <- 3:
			}
		}
		fmt.Println("End. [sender]")
	}()

	var sum int
	for e := range intChan {
		fmt.Printf("Received: %v\n", e)
		sum += e
		if sum > 10 {
			fmt.Printf("Got: %v\n", sum)
			break
		}
	}
	ticker.Stop()
	fmt.Println("End. [receiver]")
}

func TestLoggerTime(t *testing.T) {
	start := time.Now().Format("2006-01-02 15:04:05 ")
	logger.DebugWithFilePath("./logger.log", "%v\n", start)
}
