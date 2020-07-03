package main

import (
	"log"
	"net"
	"strings"
	"time"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	var buffer []byte = []byte("You are welcome. I'm server.")
	for {
		rremoteAddress := conn.RemoteAddr().String()
		index := strings.Index(rremoteAddress,":")
		time.Sleep(3 * time.Second)
		n, err := conn.Write(buffer)
		if err != nil {
			log.Println("write error:", err)
			break
		}
		log.Printf("send remote address %s ,ip address is %d ,count is %d:\n", rremoteAddress[:index],len(rremoteAddress[:index]), n)
	}
	log.Println("connection end")
}

func main() {
	addr := "0.0.0.0:9999"
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		log.Fatalf("net.ResolveTCPAddr fail :%s", addr)
	}

	listen, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Fatalf("listen %s fail:%s", addr, err)
	} else {
		log.Println("listening", addr)
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Println("listen.accept error:", err)
			continue
		}
		go handleConnection(conn)
	}
}
