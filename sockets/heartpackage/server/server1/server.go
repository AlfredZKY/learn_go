package main

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
	pb "learn_go/sockets/heartpackage/secondtest/protocol"
	"log"
	"net"
	"os"
	"strconv"
	"time"

	"google.golang.org/protobuf/proto"
)

var (
	globalMainTable = &pb.MainTable{}
	notify          = make(chan struct{})
)

func GravelChannel(bytes []byte, message chan byte) {
	for index, v := range bytes {
		if index == 0 {
			message <- v
			// log.Println(string(v))
			break
		}
		// message <- v
	}

	close(message)
}

func gzipUnCompress(content *[]byte) []byte {
	var uncompressData bytes.Buffer
	uncompressData.Write(*content)
	r, _ := gzip.NewReader(&uncompressData)
	defer r.Close()
	undatas, _ := ioutil.ReadAll(r)
	return undatas
}

func HeartBeating(conn net.Conn, message chan byte, timeout int, data []byte) {
	select {
	case fk := <-message:
		if string(fk) == "h" {
			remoteAddress := conn.RemoteAddr().String()
			log.Printf("reciver %s of  心跳包 %s:\n", remoteAddress, string(data))
			conn.Write([]byte("main miner response"))
			conn.SetDeadline(time.Now().Add(time.Duration(timeout) * time.Second))
			break
		} else if string(fk) == "r" {
			log.Println("enter")
			maininfo := &pb.MainInfo{}
			proto.Unmarshal(data[1:], maininfo)
			elementExist := false
			// TODO 单测 所以不能判断ip是否一致
			// for _, v := range globalMainTable.Maintable {
			// 	if maininfo.Ip == v.Ip {
			// 		log.Println("the some ip update some info")
			// 		v.Status = maininfo.Status
			// 		v.Weight = maininfo.Weight
			// 		elementExist = true
			// 		break
			// 	}
			// }
			if !elementExist {
				globalMainTable.Maintable = append(globalMainTable.Maintable, maininfo)
				conn.Write([]byte("receive a register info"))
				log.Printf("receiver register %d info\n", len(globalMainTable.Maintable))
				// notify <- struct{}{}
			}
			// for _, v := range globalMainTable.Maintable {
			// 	log.Println(v.Weight, v.Status, v.Ip, v.Port)
			// }

			// TODO
			// 当从miner信息收集完毕后广播给所有的从miner
		} else {
			log.Println("defaule")
		}
	case <-time.After(5 * time.Second):
		conn.Close()
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			log.Println("server read err:", err)
			return
		}

		Data := buffer[:n]
		message := make(chan byte)

		// 心跳计时
		go HeartBeating(conn, message, 15, Data)

		// 每次检测心跳是否有数据传来
		go GravelChannel(Data[:1], message)
	}
}

func server() {
	server := "127.0.0.1:7373"
	netListen, err := net.Listen("tcp", server)
	if err != nil {
		log.Println("connect error:", err)
		os.Exit(1)
	}

	log.Println("waiting for client...")

	for {
		conn, err := netListen.Accept()
		if err != nil {
			log.Println(conn.RemoteAddr().String(), "fatal err:", err)
			continue
		}
		// 设置短连接 10秒钟
		//conn.SetReadDeadline(time.Now().Add(time.Duration(20) * time.Second))
		go handleConnection(conn)
	}
}

// Init 初始化注册信息
func Init() *pb.MainInfo {
	sub1 := &pb.MainInfo{
		Weight: 90,
		Status: false,
		Ip:     "192.168.10.1",
		Port:   7373,
	}

	// maintable := &pb.MainTable{}
	// maintable.Maintable = append(maintable.Maintable, sub1)
	return sub1
}

func sendHeart(conn *net.TCPConn) {
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

func sendRegister(conn *net.TCPConn) {
	maininfo := Init()
	data, _ := proto.Marshal(maininfo)
	prefix := []byte{'r'}
	newData := []byte{}
	newData = append(newData, prefix...)
	newData = append(newData, data...)
	flag := false
	for {
		sendCount, err := conn.Write(newData)
		if err != nil {
			log.Println("client write err")
			return
		}
		buf := make([]byte, 1024)
		for {
			n, err := conn.Read(buf)
			if err != nil {
				log.Println("conn.Read err= ", err)
				return
			}
			if n > 0 {
				flag = true
				break
			}
		}
		if flag {
			log.Println("register conn write success", sendCount)
			break
		}
	}
}

func sender(conn *net.TCPConn) {
	defer func() {
		log.Println("client close")
		conn.Close()
	}()
	// go sendRegister(conn)
	sendHeart(conn)
}

func client() {
	// <- notify
	for {
		socket := make(map[string]bool)
		for _, v := range globalMainTable.Maintable {
			log.Println(v.Weight, v.Status, v.Ip, v.Port)
			server := v.Ip + ":" + strconv.Itoa(int(v.Port))
			if socket[server] {
				continue
			}
			// server := "127.0.0.1:737"
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

			sender(conn)
			// time.Sleep(time.Second * 4)
			socket[server] = true
		}

		// time.Sleep(time.Second * 2)
	}
}

func main() {
	server()
	time.Sleep(time.Second * 3600)
}
