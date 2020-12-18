package main

import (
	"fmt"
	strProto "learn_go/learnrpcx/learnrpc/proto"
	"net"
	"os"

	"github.com/golang/protobuf/proto"
)

func main() {
	// 监听
	listen, err := net.Listen("tcp", "localhost:6600")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			panic(err)
		}
		fmt.Println("new connect", conn.RemoteAddr())
		go readMessage(conn)
	}
}

func readMessage(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 4096, 4096)
	for {
		// 读消息
		cnt, err := conn.Read(buf)
		if err != nil {
			panic(err)
		}

		// 一个空结构体用反序列化消息
		strReceive := &strProto.UserInfo{}
		pData := buf[:cnt]

		// protobuf 解码
		err = proto.Unmarshal(pData, strReceive)
		if err != nil {
			panic(err)
		}

		fmt.Println("receive", conn.RemoteAddr(), strReceive)
		if strReceive.Message == "stop" {
			os.Exit(1)
		}
	}
}
