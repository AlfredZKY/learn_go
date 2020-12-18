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


// MutexChannelBuffer 互斥锁和信道
func MutexChannelBuffer(){
	flag.Parse()

	const(
		max1 = 5 	// 设置goroutine的最大数量
		max2 = 10 	// 设置每个goroutine需要写入的数据块的数量
		max3 = 10   // 代表每个数据块中写入多少个重复的数字
	)

	// 设置chann 用于gorotine 的通信
	singFlag := make(chan struct{},max1)

	// 设置buffer 用于数据的接受
	var buff bytes.Buffer

	// 设置互斥锁，保证每个goroutine能够单独对数据块进行操作
	var mu sync.Mutex

	for i:=1 ;i <= max1;i++{
		go func(id int,writer io.Writer){
			defer func(){
				singFlag <- struct{}{}
			}()
			for j:=1;j<max3;j++{
				if protecting > 0{
					mu.Lock()
				}
				// 准备数据 "\n[id:%d,iteration:%d]", id, j)
				head := fmt.Sprintf("\n[id:%d,iteration:%d]",id,j)
				data := fmt.Sprintf("\t%d\t",j*id)
				_,err := writer.Write([]byte(head))
				if err != nil{
					log.Printf("error:%s id:%d",err,id)
				}
				for k:=0;k<max3;k++{
					_,err := writer.Write([]byte(data))
					if err != nil{
						log.Printf("error:%s id:%d",err,id)
					}
				}

				if protecting > 0 {
					mu.Unlock()
				}
			}
			
		}(i,&buff)
	}

	for i:=0;i<max1;i++{
		<- singFlag
	}

	// 从buff中读取数据
	data,err := ioutil.ReadAll(&buff)
	if err != nil{
		log.Printf("The error is %s\n",err)
	}
	log.Printf("The contents:\n %s",data)
}