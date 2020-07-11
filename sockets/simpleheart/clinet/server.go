package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"reflect"

	// "strconv"
	"syscall"
	"time"
	"unsafe"

	"github.com/spf13/viper"
)

var (
	mainminercount  string = "mainminercount"
	localClientFlag bool   = false
	lastCOunt       int64
	lockFilePath    string = "/home/zky/go/src/learn_go/sockets/simpleheart/test.txt"
	configPath      string = "/home/zky/go/src/learn_go/sockets/simpleheart/"
	configName      string = "config"
	configType      string = "toml"
)

type configVal struct {
	localIPMain    string // 主mienr的IP:端口
	localIPSub     string // 从mienr的IP:端口
	lockStatus     bool   // ceph下文件锁的状态
	mainMinerCount int64  // 用来检测主miner所在服务器的状态
	subMinerCount  int64  // 用来检测从miner所在服务器的状态
	newMinerIP     string // 新主miner的IP，主miner上脚本看到这个标志，即可以执行自己的功能
	normalBoot     bool   // 从miner变成主mienr的，记录启动成功的标志
}

// TCPSubServer tcp server
func tcpServerSub() {
	server := "127.0.0.1:7373"
	netListen, err := net.Listen("tcp", server)
	if err != nil {
		log.Fatal("create sub socker error:", err)
		os.Exit(1)
	}

	defer netListen.Close()

	log.Println("main miner waiting for client...")
	for {
		// 如果自己被选为新主后，自己不会在和从miner进行心跳，新主会启用自己的心跳包
		conn, err := netListen.Accept()
		if err != nil {
			log.Println(conn.RemoteAddr().String(), "fatal err", err)
			continue
		}
		// 设置连接等待时间
		conn.SetReadDeadline(time.Now().Add(time.Duration(18) * time.Second))
		go handleConnectionTCPSub(conn)
	}
}

// HandleConnectionTCP 处理tcp的连接
func handleConnectionTCPSub(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			// TODO 超时等待后，现在此写入自己的IP,并在此判断主mienr是否已关闭，和文件锁状态。都为true时，即可正常执行
			log.Println("server read err:", err)
			log.Println("local server become new main miner")
			localIP, _ := externalIP()
			modifyConfig("newminerip", localIP.String())
			time.Sleep(time.Second * 4)
			// readValFromConfig()
			for {
				viper.SetConfigName("config")
				viper.AddConfigPath("/home/zky/go/src/learn_go/sockets/simpleheart/")
				viper.SetConfigType("toml")
				err := viper.ReadInConfig() // 会查找和读取配置文件
				if err != nil {             // Handle errors reading the config file
					log.Println(fmt.Errorf("Fatal read error config file: %s", err))
				}
				if viper.GetBool("lockStatus") {
					log.Println("file lock has been del!!!")
					break
				}
			}

			time.Sleep(time.Second * 10)
		}

		Data := buffer[:n]
		message := make(chan byte)

		// 心跳计时
		go heartBeatingSub(conn, message, 19, Data)
		go channelNotify(Data, message)
	}
}

// SubHeartBeating 从miner的心跳处理
func heartBeatingSub(conn net.Conn, message chan byte, timeout int, data []byte) {
	select {
	case fk := <-message:
		if string(fk) == "h" {
			remoteAddress := conn.RemoteAddr().String()
			// index := strings.Index(remoteAddress, ":")
			// 消息主题
			message := data[1:]
			// TODO 单测 所以不能判断ip是否一致 最后删除该字段
			// elementExist := false
			// if !elementExist {
			// 	// TODO scoketsub map 临时有用，后面在正式服务器上会删掉

			// }
			log.Printf("from %s receive a heat message %s", remoteAddress, message)
			// 重置超时时间
			conn.SetDeadline(time.Now().Add(time.Duration(timeout) * time.Second))
			conn.Write([]byte("main miner response"))
		} else if string(fk) == "p" {
			// TODO 扩展消息处理
		} else if string(fk) == "s" {
			// TODO 扩展消息处理
		}
	case <-time.After(time.Second * 5):
		// 超时关闭客户端的连接
		conn.Close()
		// 在此成为新主
		log.Println("heart local server become new main miner")

	}
}

// GravelChannel chan 信道通知
func channelNotify(bytes []byte, message chan byte) {
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

// sendHeartMain 发送心跳报文到主miner
func sendHeartMain(conn *net.TCPConn) {
	defer conn.Close()
	for {
		now := time.Now()
		// 设置写超时
		begin := now.Local().UnixNano() / (1000 * 1000)
		if _, err := conn.Write([]byte("hhello")); err != nil {
			// 这个错误是不能转换*net.OpError
			if err == syscall.EINVAL {
				return
			}
			// 转换成*net.OpError
			opErr := (*net.OpError)(unsafe.Pointer(reflect.ValueOf(err).Pointer()))
			if opErr.Err.Error() == "i/o timeout" {
				end := time.Now().Local().UnixNano() / (1000 * 1000)
				log.Printf("Write timeout! end: %d, begin: %d, timeOut: %dms", end, begin, end-begin)
				return
			}
		}

		time.Sleep(time.Second * 10)
	}
}

// sender 处理心跳和注册消息
func sender() {
	conn := clientMain()
	go sendHeartMain(conn)
}

// clientMain 创建客户端 与主miner通信
func clientMain() *net.TCPConn {
	server := ""
	for {
		for {
			time.Sleep(time.Second * 4)
			// TODO 到时候会改成配置文件的具体路径 采用热更新的方式进行
			// viper.SetConfigFile("/home/zky/go/src/learn_go/sockets/simpleheart/config.toml")
			viper.SetConfigName("config")
			viper.AddConfigPath("/home/zky/go/src/learn_go/sockets/simpleheart/")
			viper.SetConfigType("toml")
			err := viper.ReadInConfig() // 会查找和读取配置文件
			if err != nil {             // Handle errors reading the config file
				log.Println(fmt.Errorf("Fatal read error config file: %s", err))
				continue
			}
			// subv := viper.Sub("MonitorUnit")
			// log.Println("address is ", subv.GetString("localipsub"), *subv)
			// server = subv.GetString("localipsub")
			log.Println("address is ", viper.GetString("localIPSub"))
			server = viper.GetString("localIPSub")
			if server == "" {
				time.Sleep(time.Second * 2)
				continue
			} else {
				break
			}

		}
		// server := "127.0.0.1:7374"
		// tcpAddr, err := net.ResolveTCPAddr("tcp4", server)
		log.Println("server is ", server)
		tcpAddr, err := net.ResolveTCPAddr("tcp4", server)
		//time.Sleep(5 * time.Second)
		if err != nil {
			log.Println(os.Stderr, "fatal error:", err)
			time.Sleep(2 * time.Second)
			continue
		}
		log.Println("address is ", server, tcpAddr)
		conn, err := net.DialTCP("tcp4", nil, tcpAddr)
		if err != nil {
			log.Println("fatal error:", err)
			time.Sleep(time.Second * 2)
			continue
		}

		log.Println(conn.RemoteAddr().String(), "connection success!")
		return conn
	}
}

func modifyConfig(key string, value interface{}) {
	viper.SetConfigName(configName)
	viper.AddConfigPath(configPath)
	viper.SetConfigType(configType)
	viper.ReadInConfig()

	// 根据类型进行判断
	switch v := value.(type) {
	case bool:
		{
			log.Println("bool v is:", v)
			viper.Set(key, v)
		}
	case string:
		{
			log.Println("string v is:", v)
			viper.Set(key, v)
		}
	}

	err := viper.WriteConfigAs(configPath + configName + "." + configType) //写入文件
	if err != nil {                                                        // Handle errors reading the config file
		log.Fatalf("%s \n", err)
	}
}

func readValFromConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath("/home/zky/go/src/learn_go/sockets/simpleheart/")
	viper.SetConfigType("toml")
	err := viper.ReadInConfig() // 会查找和读取配置文件
	if err != nil {             // Handle errors reading the config file
		log.Println(fmt.Errorf("Fatal read error config file: %s", err))
	}

	var configv configVal
	if err := viper.Unmarshal(&configv); err != nil {
		fmt.Printf("err:%s", err)
	}
	log.Println("address is ", viper.GetString("localIPMain"))
	log.Println("address is ", viper.GetString("localIPSub"))
	log.Println("address is ", viper.GetString("lockStatus"))
	log.Println("address is ", viper.GetString("mainMinerCount"))
	log.Println("address is ", viper.GetString("subMinerCount"))
	log.Println("address is ", viper.GetString("normalBoot"))
	log.Println("address is ", viper.GetString("newMinerIP"))
}

// 当从miner判定主mienr有网络问题时，从miner的脚本会检测主miner是否在更新count,如果不更新则自己删除锁文件并更新锁状态
func checkMainCount() {
	viper.SetConfigName(configName)
	viper.AddConfigPath(configPath)
	viper.SetConfigType(configType)
	err := viper.ReadInConfig() // 会查找和读取配置文件
	if err != nil {             // Handle errors reading the config file
		log.Println(fmt.Errorf("Fatal read error config file: %s", err))
	}
	count := 0
	mainminercount := viper.GetInt64(mainminercount)
	if mainminercount-lastCOunt > 0 {
		// 加一个容错值
		count++
		log.Println("main miner server online")
	} else {
		// 如果在容错值内还没有更新就认为时宕机，由从miner删除锁文件后，让从miner正式启动
		for {
			// err := removefile()
		}
		// 正常关闭后，更新锁状态，此时就无锁了，从miner可以启动了。
		// modifyVariableConfig(lockstatus, false)
	}
	lastCOunt = mainminercount

}

// 解析出本地Ip
func externalIP() (net.IP, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return nil, err
		}
		for _, addr := range addrs {
			ip := getIPFromAddr(addr)
			if ip == nil {
				continue
			}
			return ip, nil
		}
	}
	return nil, errors.New("connected to the network?")
}

func getIPFromAddr(addr net.Addr) net.IP {
	var ip net.IP
	switch v := addr.(type) {
	case *net.IPNet:
		ip = v.IP
	case *net.IPAddr:
		ip = v.IP
	}
	if ip == nil || ip.IsLoopback() {
		return nil
	}
	ip = ip.To4()
	if ip == nil {
		return nil // not an ipv4 address
	}

	return ip
}

func main() {
	go modifyConfig("newminerip", "")
	time.Sleep(time.Second * 5)
	localIP, _ := externalIP()
	modifyConfig("newminerip", localIP.String())
	// 启动协程
	//go tcpServerSub()
	//go sender()
	time.Sleep(1000 * time.Second)
}
