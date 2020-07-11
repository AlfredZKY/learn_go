package main

import (
	"errors"
	"fmt"
	"net"
	"os"

	"log"
	"time"

	dingtalk_robot "github.com/JetBlink/dingtalk-notify-go-sdk"
	"github.com/spf13/viper"
)

var (
	mainminercount string = "mainminercount"
	lockstatus     string = "lockstatus"
	mainFlag       bool   = true
	lastCount      int64
	faultTolerant  int32 = 5
	faultcount     int32
	globalcount    int64

	lockFilePath string = "/home/zky/go/src/learn_go/sockets/simpleheart/test.txt"
	configPath   string = "/home/zky/go/src/learn_go/sockets/simpleheart/"
	configName   string = "config"
	configType   string = "toml"

	apiFile string = "/home/zky/go/src/learn_go/sockets/simpleheart/api"
)

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
	case int64:
		{
			log.Println("int64 v is:", v)
			viper.Set(key, v)
		}
	}

	err := viper.WriteConfigAs(configPath + configName + "." + configType) //写入文件
	if err != nil {                                                        // Handle errors reading the config file
		log.Fatalf("%s \n", err)
	}
}

// Send is to send some messages
func Send(LocalIP, messages string) {
	msg := map[string]interface{}{
		"msgtype": "text",
		"text": map[string]string{
			"content": LocalIP + messages,
		},
		"at": map[string]interface{}{
			"atMobiles": []string{},
			"isAtAll":   false,
		},
	}
	// https://oapi.dingtalk.com/robot/send?access_token=a15a0c483228898166cbad1a07c475fc9bab5891bf069adc8d1db3db9d87235f
	robot := dingtalk_robot.NewRobot("a15a0c483228898166cbad1a07c475fc9bab5891bf069adc8d1db3db9d87235f", "SEC8229ba11fe5487a3579d37321d88c0dc1ee5621afd0b34583956745e7fe66156")

	if err := robot.SendMessage(msg); err != nil {
		fmt.Println(err)
	}
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
	mainminercount := viper.GetInt64(mainminercount)
	if mainminercount-lastCount > 0 {
		log.Println("main miner server online")
	} else {
		// 如果在容错值内还没有更新就认为时宕机，让从miner正式启动
		// 加一个容错值
		faultcount++
	}
	if faultcount > faultTolerant {
		time.Sleep(time.Second * 5)
		localIP, _ := externalIP()
		modifyConfig("newminerip", localIP.String())

		time.Sleep(time.Second * 5)

		// TODO 再次检查检查锁状态
		count := 0
		for {
			viper.SetConfigName(configName)
			viper.AddConfigPath(configPath)
			viper.SetConfigType(configType)
			err := viper.ReadInConfig() // 会查找和读取配置文件
			if err != nil {             // Handle errors reading the config file
				log.Println(fmt.Errorf("Fatal read error config file: %s", err))
			}
			lockstatus := viper.GetBool(lockstatus)
			if !lockstatus {
				// 无锁 从miner可以正常启动
				break
			} else {
				count++
				time.Sleep(time.Minute * 1)
			}
			if count > 5 {
				Send("", "main miner process don't normal close")
				time.Sleep(time.Minute * 1)
			}
		}
		modifyAPI()
		mainFlag = false
		lastCount = 0
		faultcount = 0
	}
	lastCount = mainminercount

}

// 修改api文件
func modifyAPI() {
	//OpenFile指定文件打开方式，只读，只写，读写和权限
	file4, err7 := os.OpenFile(apiFile, os.O_RDWR, 0666)
	defer file4.Close()
	if err7 != nil {
		log.Fatal(file4)
	}
	//向文件写入数据
	localIP, _ := externalIP()
	log.Println("local ip", localIP.String())

	api := "/ip4/" + localIP.String() + "/tcp/2345/http"
	file4.Write([]byte(api))
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

// 定时更新轮数
func setHeartCount() {
	c := time.Tick(time.Second * 15)
	for {
		<-c
		if mainFlag {
			globalcount++
			log.Println(time.Now().Format("2006-01-02 15:04:05"), globalcount)
			modifyConfig("mainminercount", globalcount)
		} else {
			checkMainCount()
			log.Println("sub miner has been become main mainer")
		}
	}
}

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	// 启动协程
	go setHeartCount()
	// modifyAPI()
	//Send("","main miner process don't normal close")
	time.Sleep(1000 * time.Second)
}
