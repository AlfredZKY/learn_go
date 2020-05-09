package mutexs

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"sync"
)

// protecting 用于指示是否使用互斥锁来保护数据写入
// 若值等于0则表示不使用，若值值大于0则表示使用

var protecting uint

func init() {
	flag.UintVar(&protecting, "protecting", 1, "It indicates whether to use a mutex to protect data writing.")
}

// MutexChannel 互斥锁和信道的使用
func MutexChannel() {
	flag.Parse()

	// buffer 代表缓冲区
	var buffer bytes.Buffer

	// 设置一些常量
	const (
		max1 = 5  // 代表启用goroutine的数量
		max2 = 10 // 代表gorotine需要写入的数据块的数量
		max3 = 10 // 代表每个数据块中需要有多个重复的数字
	)

	// mu 代表互斥锁
	var mu sync.Mutex

	// sign 代表信号的通道
	signFlag := make(chan struct{}, max1)

	for i := 1; i <= max1; i++ {
		go func(id int, writer io.Writer) {
			defer func() {
				signFlag <- struct{}{}
			}()

			for j := 1; j <= max2; j++ {
				// 准备数据
				header := fmt.Sprintf("\n[id:%d,iteration:%d]", id, j)
				data := fmt.Sprintf(" %d", id*j)

				// 写入数据
				if protecting > 0 {
					mu.Lock()
				}

				_, err := writer.Write([]byte(header))
				if err != nil {
					log.Printf("error :%s [%d]", err, id)
				}

				for k := 0; k < max3; k++ {
					_,err := writer.Write([]byte(data))
					if err != nil{
						log.Printf("error:%s [%d]",err,id)
					}
				}
				if protecting > 0{
					mu.Unlock()
				}
			}

		}(i, &buffer)
	}

	for i := 0; i < max1; i++ {
		<-signFlag
	}

	data, err := ioutil.ReadAll(&buffer)
	if err != nil {
		log.Fatalf("fatal error:%s", err)
	}
	log.Printf("The contents:\n%s", data)
}
