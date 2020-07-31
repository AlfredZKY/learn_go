package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"time"

	dingtalk_robot "github.com/JetBlink/dingtalk-notify-go-sdk"
	"github.com/spf13/viper"
	"github.com/fsnotify/fsnotify"
)

var (
	mainminercount    string = "mainminercount"
	globalcount       int64  = 0
	newminerip        string = "newminerip"
	lockstatus        string = "lockstatus"
	lastCOunt         int64
	flags             bool
	lockFilePath      string = "/home/zky/go/src/learn_go/sockets/simpleheart/test.txt"
	lotusStorageMiner string = "/home/zky/go/src/learn_go/sockets/simpleheart"
	configPath        string = "/home/zky/go/src/learn_go/sockets/simpleheart"
	configName        string = "config"
	configType        string = "toml"
)

// 协程可以在配置文件中自增，来让从miner判断自己是否宕机
func modifyVariableConfig(key string, value interface{}) {
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

// 正常关闭lotus-storage-miner stop
func closeMinerProcess() error {
	args := []string{lotusStorageMiner, " stop"}
	cmd := exec.Command("bash", args...)
	err := cmd.Start()

	if err != nil {
		return err
	}
	err = cmd.Wait()
	return err
}

// 实际上就是删除锁文件
func removefile() error {
	// TODO 首先判断文件是否存在
	_, err := os.Stat(lockFilePath)
	if err != nil {
		log.Printf("%s has not exist!\n", lockFilePath)
		return nil
	}
	args := []string{" -rf ", lockFilePath}
	cmd := exec.Command("rm", args...)
	err = cmd.Start()
	if err != nil {
		return err
	}
	err = cmd.Wait()
	log.Println(err)
	return err
}

func checkMainIP() string {
	viper.SetConfigName(configName)
	viper.AddConfigPath(configPath)
	viper.SetConfigType(configType)
	err := viper.ReadInConfig() // 会查找和读取配置文件
	if err != nil {             // Handle errors reading the config file
		log.Println(fmt.Errorf("Fatal read error config file: %s", err))
	}
	newminerips := viper.GetString(newminerip)
	return newminerips
}

// 检测配置文件中新主IP的产生，如果产生，正常关闭lotus-storage-miner进程
func checkNewMinerIP() {
	newminerip := checkMainIP()
	if newminerip == "" {
		// 执行正常的关闭动作，否则继续检测
		for {
			err := closeMinerProcess()
			if err != nil {
				log.Println("lotus-storage-miner has been close success!!!")
				break
			} else {
				log.Println("continue close lotus-storage-miner")
				time.Sleep(time.Second * 2)
				continue
			}
			// TODO 删除锁文件
		}
		// 正常关闭后，更新锁状态，此时就无锁了，从miner可以启动了。
		for {
			err := removefile()
			if err != nil {
				log.Println("lock file has been del")
				break
			}
			time.Sleep(time.Second * 2)
		}
		modifyVariableConfig(lockstatus, false)
		return
	}
	log.Println("main miner is running")
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
			if ip == nil || ip[0] != 192 {
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
		// globalcount++
		// log.Println(time.Now().Format("2006-01-02 15:04:05"), globalcount)
		// go modifyVariableConfig(mainminercount, globalcount)
		if flags {
			log.Println("runing sub mienr script")
			checkNewMinerIP()
		} else {
			log.Println("runing sub mienr script")
			localIP, _ := externalIP()
			newminerip := checkMainIP()
			if newminerip != "" && localIP.String() == newminerip {
				flags = true
			}
		}
	}
}

func useFlag() {
	// 参数地址 参数名 参数的默认值 参数的含义(简短的说明)
	// flag.StringVar(&name, "name", "everyone", "The greeting object.") 对应的是地址
	flag.BoolVar(&flags, "flags", true, "默认是主标志")

	// 使无参或者自定义参数 对flag.Usage重新赋值,flag.Usage的类型是func(),即一种无参数声明且无结果声明的函数类型
	// flag.Usage = func(){
	// 	fmt.Fprintf(os.Stderr,"Usage of %s:\n","question")
	// 	flag.PrintDefaults()
	// }

	// 把用户传递的命令行参数解析为对应变量的值
	flag.Parse()

	fmt.Println("cliFlag=", flags)

}

var (
	apiFile string = "/home/zky/go/src/learn_go/sockets/simpleheart/api"
)

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
	time.Sleep(time.Second * 20)
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

// TimeWrite 定时更新轮数
func TimeWrite() {
	c := time.Tick(time.Second * 15)
	for {
		<-c
		globalcount++
		log.Println(time.Now().Format("2006-01-02 15:04:05"), globalcount)
		writeConfigs(globalcount)
	}
}

func writeConfigs(count int64) {
	viper.SetConfigName(configName)
	viper.AddConfigPath(configPath)
	viper.SetConfigType(configType)
	viper.ReadInConfig()
	localIP, _ := externalIP()
	ListenAddress := "/ip4/" + localIP.String() + "/tcp/s/http" 
	RemoteListenAddress := localIP.String() + "/tcp/s/http"
	viper.Set("API.ListenAddress", ListenAddress)
	viper.Set("API.RemoteListenAddress", RemoteListenAddress)
	err := viper.WriteConfig() //写入文件
	if err != nil {            // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

func watchConfig() {
    viper.WatchConfig()
    viper.OnConfigChange(func(e fsnotify.Event) {
        fmt.Printf("Config file changed: %sn", e.Name)
    })
}

func main() {
	// go TimeWrite()
	// useFlag()
	// go setHeartCount()
	// err := removefile()
	// if err != nil {
	// 	log.Println("normal delete")
	// }
	// time.Sleep(time.Second * 315360000)
	viper.SetConfigName(configName)
	viper.AddConfigPath(configPath)
	viper.SetConfigType(configType)
	viper.ReadInConfig()
	// if err != nil {
	// 	log.Println(err)
	// }
	log.Println(viper.GetBool("lockstatus"))
}
