package main

import (
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

var (
	globalrounds int32
	globalflag   bool
	lock         sync.Mutex

	runflag bool
)

func HandleConn(conn net.Conn) {
	// 函数调用完毕，要自动关闭
	defer conn.Close()

	// 获取客户端的网络地址信息 并转换成字符串
	addr := conn.RemoteAddr().String()

	fmt.Println(addr, "connect successful")
	buf := make([]byte, 2048)

	for {
		log.Println(time.Now().Format("handle 2006-01-02 15:04:05"), globalflag)
		if globalflag {
			runflag = true
			return
		}
		// 读取用户数据
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("handle err is : ", err)
			return
		}
		fmt.Printf("read buf=%s\n", string(buf[:n]))

		if "exit" == string(buf[:n-1]) {
			// 如果输入exit，则推出连接
			fmt.Println(addr, "exit")
			return
		}
		conn.Write([]byte(strings.ToUpper(string(buf[:n]))))
	}

}

func setRounds() {
	c := time.Tick(time.Second * 2)
	for {
		<-c
		go func() {
			if runflag{
				return
			}
			log.Println(time.Now().Format("2006-01-02 15:04:05"), globalrounds)
			lock.Lock()
			globalrounds++
			lock.Unlock()
			if globalrounds > 10 {
				globalflag = true
			}
		}()
	}
}

func tcpServer() {
	listen, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("main err is :", err)
		return
	}

	// 最后才执行关闭操作
	defer listen.Close()

	for {
		// log.Println(time.Now().Format("tcp 2006-01-02 15:04:05"), *flags)
		if globalflag {
			fmt.Println("tcp is over")
			return
		}
		// 阻塞等待用户连接
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("main for err is :", err)
			return
		}
		go HandleConn(conn)
	}

}

func main() {
	// 监听
	go setRounds()
	go setRounds()
	go tcpServer()
	time.Sleep(200 * time.Second)
}
