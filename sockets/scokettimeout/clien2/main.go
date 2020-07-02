package main

import (
	"log"
	"net"
	"os"
	"time"
)

func main() {
	connTimeout := 3 * time.Second

	// 连接设置3秒超时
	conn, err := net.DialTimeout("tcp", "127.0.0.1:9999", connTimeout)
	if err != nil {
		log.Println("dial failed:", err)
		os.Exit(1)
	}

	defer conn.Close()
	readTimeout := 12 * time.Second
	buffer := make([]byte, 512)

	for {
		err = conn.SetDeadline(time.Now().Add(readTimeout))
		if err != nil {
			log.Println("setReadDeadline failed:", err)
		}

		n, err := conn.Read(buffer)
		if err != nil {
			log.Println("read failed:", err)
			return 
		}
		log.Println("count:", n, "msg:", string(buffer))
		// time.Sleep(2 * time.Second)
	}
}
