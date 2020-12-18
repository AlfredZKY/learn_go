package main

import (
	"bytes"
	"log"
	"net"
	"os"
	"time"

	//"os"
	"io/ioutil"
	//"fmt"

	"compress/gzip"
	pb "learn_go/sockets/heartpackage/secondtest/protocol"

	"google.golang.org/protobuf/proto"
)

func gzipCompress(content *[]byte) []byte {
	var compressData bytes.Buffer
	gzipWriter := gzip.NewWriter(&compressData)
	defer gzipWriter.Close()
	gzipWriter.Write(*content)
	gzipWriter.Flush()
	return compressData.Bytes()
}

func gzipUnCompress(content *[]byte) []byte {
	var uncompressData bytes.Buffer
	uncompressData.Write(*content)
	r, _ := gzip.NewReader(&uncompressData)
	defer r.Close()
	undatas, _ := ioutil.ReadAll(r)
	return undatas
}

// Init 初始化注册信息
func Init() *pb.MainInfo {
	sub1 := &pb.MainInfo{
		Weight:    20,
		Status:    false,
		Ip:        "192.168.10.2",
		Reserfile: "client2",
	}
	// sub2 := &pb.MainInfo{
	// 	Weight:    80,
	// 	Status:    false,
	// 	Ip:        "192.168.10.21",
	// 	Reserfile: "recover",
	// }

	// maintable := &pb.MainTable{}
	// maintable.Maintable = append(maintable.Maintable, sub1)
	return sub1
}

func senderHeartInfo(conn *net.TCPConn){
	for i := 0; i < 12; i++ {
		_, err := conn.Write([]byte("hello"))
		if err != nil {
			log.Println("client write err")
			return
		}
		//log.Println("conn write success")
		time.Sleep(2 * time.Second)
	}
}

func sendRegister(conn *net.TCPConn){
	maininfo := Init()
	data, _ := proto.Marshal(maininfo)
	// for _, v := range maintable.Maintable {
	// 	log.Println(v.Weight, v.Status, v.Ip)
	// }

	prefix := []byte{'r'}
	newData := []byte{}
	newData = append(newData,prefix...)
	newData = append(newData,data...)

	_, err := conn.Write(newData)
	if err != nil {
		log.Println("client write err")
		return
	}

	log.Println("register conn write success")
}

// 心跳包和注册信息发送模块
func sender1(conn *net.TCPConn) {
	defer func() {
		log.Println("client close")
		conn.Close()
	}()

	// go sendRegister(conn)
	maininfo := Init()
	data, _ := proto.Marshal(maininfo)
	// for _, v := range maintable.Maintable {
	// 	log.Println(v.Weight, v.Status, v.Ip)
	// }

	prefix := []byte{'r'}
	newData := []byte{}
	newData = append(newData,prefix...)
	newData = append(newData,data...)

	sendCount, err := conn.Write(newData)
	if err != nil {
		log.Println("client write err")
		return
	}

	log.Println("register conn write success",sendCount)

	for i := 0; i < 12; i++ {
		_, err := conn.Write([]byte("hello"))
		if err != nil {
			log.Println("client write err")
			return
		}
		//log.Println("conn write success")
		time.Sleep(2 * time.Second)
	}
}
func sender(conn *net.TCPConn) {
	defer func() {
		log.Println("client close")
		conn.Close()
	}()
	// go sendRegister(conn)
	maininfo := Init()
	data, _ := proto.Marshal(maininfo)
	// for _, v := range maintable.Maintable {
	// 	log.Println(v.Weight, v.Status, v.Ip)
	// }

	prefix := []byte{'r'}
	newData := []byte{}
	newData = append(newData,prefix...)
	newData = append(newData,data...)
	flag := false
	for {
		sendCount, err := conn.Write(newData)
		if err != nil {
			log.Println("client write err")
			return
		}
		buf := make([]byte,1024)
		for {
			n,err := conn.Read(buf)
			if err != nil {
				log.Println("conn.Read err= ",err)
				return
			}
			if n > 0 {
				flag = true
				break
			}
		}
		if flag{
			log.Println("register conn write success",sendCount)
			break
		}
	}

	for i := 0; i < 12; i++ {
		_, err := conn.Write([]byte("hello"))
		if err != nil {
			log.Println("client write err")
			return
		}
		//log.Println("conn write success")
		time.Sleep(time.Second)
	}
}

func main() {
	server := "127.0.0.1:7373"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", server)

	if err != nil {
		log.Println(os.Stderr, "fatal error:", err)
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Println("fatal error:", err)
		os.Exit(1)
	}

	log.Println(conn.RemoteAddr().String(), "connection success!")

	// go func() {
	// 	i := 0
	// 	for {
	// 		// fmt.Println(i)
	// 		time.Sleep(time.Second)
	// 		i++
	// 	}
	// }()

	
	sender(conn)
	time.Sleep(time.Second * 4)
	log.Println("send over")
}
