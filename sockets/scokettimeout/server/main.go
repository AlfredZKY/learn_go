package main

import (
	"time"
	"log"
	"net"
)

func handleConnection(conn net.Conn){
	defer conn.Close()

	var buffer []byte = []byte("You are welcome. I'm server.")
	for {
		time.Sleep(3*time.Second)
		n,err := conn.Write(buffer)
		if err != nil {
			log.Println("write error:",err)
			break
		}
		log.Println("send:",n)
	}
	log.Println("connection end")
}

func main(){
	addr:="0.0.0.0:9999"
	tcpAddr,err := net.ResolveTCPAddr("tcp",addr)
	if err != nil {
		log.Fatalf("net.ResolveTCPAddr fail :%s",addr)
	}

	listen,err := net.ListenTCP("tcp",tcpAddr)
	if err != nil {
		log.Fatalf("listen %s fail:%s",addr,err)
	}else{
		log.Println("listening",addr)
	}

	for {
		conn ,err := listen.Accept()
		if err != nil {
			log.Println("listen.accept error:",err)
			continue
		}
		go handleConnection(conn)
	}
}