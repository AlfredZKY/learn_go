package learnrpc

import (
	"bufio"
	"fmt"
	stProto "learn_go/learnrpcx/learnrpc/client/proto"
	"net"
	"os"
	"time"
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
	sender := bufio.NewScanner(os.Stdin)
	for sender.Scan() {
		cnt++
		stSend := &stProto.UserInfo{
			Message: "1",
			Length:  int32(2),
			Cnt:     int32(cnt),
		}
		_ = stSend
	}

}
