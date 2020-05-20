package main

import (
	"bufio"
	"fmt"
	stProto "learn_go/learnrpcx/learnrpc/proto"
	"net"
	"os"
	"time"

	"github.com/golang/protobuf/proto"
	// "github.com/gogo/protobuf/proto" 同上
)

func main() {
	strIP := "localhost:6600"
	var conn net.Conn
	var err error

	// 连接服务器
	for conn, err = net.Dial("tcp", strIP); err != nil; conn, err = net.Dial("tcp", strIP) {
		fmt.Println("connect", strIP, "fail")
		time.Sleep(time.Second)
		fmt.Println("reconnect")
	}
	fmt.Println("connect", strIP, "success")
	defer conn.Close()

	// 发送消息
	cnt := 0

	// 从控制台读取输入消息
	sender := bufio.NewScanner(os.Stdin)
	for sender.Scan() {
		cnt++
		stSend := &stProto.UserInfo{
			Message: "1",
			Length:  *proto.Int(len(sender.Text())),
			Cnt:     *proto.Int(cnt),
		}
		// 利用protobuf编码
		pData, err := proto.Marshal(stSend)
		if err != nil {
			panic(err)
		}

		// 发送消息
		conn.Write(pData)
		if sender.Text() == "stop" {
			return
		}
	}
}
