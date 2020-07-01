package main

import (
	"os"
	"fmt"
	"net"
)

func main() {
	// 主动连接服务器
	conn, err := net.Dial("tcp", "127.0.0.1:8000") // 服务器的IP和端口
	if err != nil {
		fmt.Println("err is :", err)
		return
	}

	defer conn.Close()

	// 接受服务器回复的数据
	go func (){
		// 从键盘输入内容，并发送给服务器
		str := make([]byte,1024)
		for{
			// 从键盘上读取内容，并存储在str中
			n,err := os.Stdin.Read(str)
			if err != nil {
				fmt.Println("conn.Read err is :",err)
				return
			}
			conn.Write(str[:n])
		}
		
	}()

	// 切片缓冲
	buf := make([]byte,1024)
	for {
		n,err := conn.Read(buf)
		if err != nil {
			fmt.Println("conn.Read err= ",err)
			return
		}
		fmt.Println(string(buf[:n]))
	}
}
