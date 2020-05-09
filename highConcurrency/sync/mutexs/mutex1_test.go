package mutexs

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"sync"
	"testing"
	"time"
)

// singleHandler 代表单次处理函数的类型
type singleHandler func() (data string, n int, err error)

// 代表处理流程配置的类型
type handlerConfig struct {
	handler   singleHandler // 单次处理函数
	goNum     int           // 需要启用的goroutine的数量
	number    int           // 单个goroutine中处理的次数
	interval  time.Duration // 单个goroutine中的处理间隔时间
	counter   int           // 数据量的计数器，以字节为单位
	counterMu sync.Mutex    // 数据量计数器专用的互斥锁
}

// count 会增加计数器的值，并会返回增加后的计数
func (hc *handlerConfig) count(increment int) int {
	hc.counterMu.Lock()
	defer hc.counterMu.Unlock()
	hc.counter += increment
	return hc.counter
}

func TestMutexRW(t *testing.T) {
	// mu 代表一下流程要使用互斥锁,在函数中直接使用，不要传递
	var mu sync.Mutex

	// genWriter 代表的用于生成写入函数的函数
	genWriter := func(writer io.Writer) singleHandler {
		return func() (data string, n int, err error) {
			// 准备数据
			data = fmt.Sprintf("%s\t", time.Now().Format(time.StampNano))

			// 写入数据
			mu.Lock()

			defer mu.Unlock()
			n, err = writer.Write([]byte(data))
			return
		}
	}

	// genReader 代表的用于生成读取函数的函数
	genReader := func(reader io.Reader) singleHandler {
		return func() (data string, n int, err error) {
			buffer, ok := reader.(*bytes.Buffer)
			if !ok {
				err = errors.New("unsupported reader")
				return
			}

			// 读取数据
			mu.Lock()
			defer mu.Unlock()
			data, err = buffer.ReadString('\t')
			n = len(data)
			return
		}
	}

	// buffer 代表缓冲区
	var buffer bytes.Buffer

	// 数据写入配置
	writingConfig := handlerConfig{
		handler:  genWriter(&buffer),
		goNum:    5,
		number:   4,
		interval: time.Millisecond * 100,
	}

	// 数据读取配置
	readingConfig := handlerConfig{
		handler:  genReader(&buffer),
		goNum:    5,
		number:   4,
		interval: time.Millisecond * 100,
	}

	// sign 代表信号的通道
	sign := make(chan struct{}, writingConfig.goNum + readingConfig.goNum)

	// 启用多个goroutine 对缓冲区进行多次数据写入
	for i := 1; i <= writingConfig.goNum; i++ {
		go func(i int) {
			defer func() {
				sign <- struct{}{}
			}()

			for j := 1; j <= writingConfig.number; j++ {
				time.Sleep(writingConfig.interval)
				data, n, err := writingConfig.handler()
				if err != nil {
					log.Printf("writer [%d-%d]: error: %s", i, j, err)
					continue
				}

				total := writingConfig.count(n)
				log.Printf("writer [%d-%d]: %s (total: %d)", i, j, data, total)
			}
		}(i)
	}

	// 启动多个goroutine对缓冲区进行多次数据的读取
	for i := 1; i <= readingConfig.goNum; i++ {
		go func(i int) {
			defer func() {
				sign <- struct{}{}
			}()

			for j := 1; j <= readingConfig.number; j++ {
				time.Sleep(readingConfig.interval)

				var data string
				var n int
				var err error
				for {
					data, n, err = readingConfig.handler()
					if err != nil || err != io.EOF {
						break
					}

					// 如果读比写快(读时会发生EOF错误),那就等一会再读
					time.Sleep(readingConfig.interval)
				}
				if err != nil {
					log.Printf("reader [%d-%d]: error:%s", i, j, err)
					continue
				}
				total := readingConfig.count(n)
				log.Printf("reader [%d-%d]: %s (total :%d)", i, j, data, total)
			}
		}(i)
	}

	// signNumber 代表要接受的信号的数量
	signNumber := writingConfig.goNum + readingConfig.goNum

	// 等待上面启用的所有的goroutine的运行全部接受
	for j := 0; j < signNumber; j++ {
		<-sign
	}
}

func TestMutexSort(t *testing.T) {
	var mutex sync.Mutex
	fmt.Println("Lock the lock. (main)")
	mutex.Lock()
	fmt.Println("The lock is locked. (main)")
	for i := 1; i <= 3; i++ {
		go func(i int) {
			fmt.Printf("Lock the lock. (g%d)\n", i)
			mutex.Lock()
			fmt.Printf("The lock is locked. (g%d)\n", i)
		}(i)
	}
	// 休眠允许goroutine能够执行，
	time.Sleep(time.Second)
	fmt.Println("Unlock the lock. (main)")

	// 解锁之后，唤醒三个goroutine进行竞争,
	mutex.Unlock()
	fmt.Println("The lock is unlocked. (main)")

	// 休眠 允许goroutine打印出结果
	// time.Sleep(time.Second * 2)
	time.Sleep(time.Second)
}

func TestMultiUnLock(t *testing.T) {
	// 恢复在之前低版本是可以恢复的
	defer func() {
		fmt.Println("Try to recover the panic.")
		if p := recover(); p != nil {
			fmt.Printf("Recovered the panic(%#v).\n", p)
		}
	}()
	var mutex sync.Mutex
	fmt.Println("Lock the lock")
	mutex.Lock()
	fmt.Println("The lock is locked.")
	fmt.Println("Unlock the lock.")
	mutex.Unlock()
	fmt.Println("The lock is unlocked.")
	fmt.Println("Unlock the lock again.")
	// 重复对同一个互斥锁解锁会引发一个运行恐慌，且不可恢复。
	mutex.Unlock()
}
