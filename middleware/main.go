package main

import (
	"fmt"
	"os"

	// "learn_go/middleware/jiankong"
	// "learn_go/middleware/readconfig"

	"github.com/spf13/viper"
)

// CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go
func main() {
	// p,err := readconfig.ReadConf("/home/zky/go/src/middleware/config.toml")
	// if err != nil{
	// 	fmt.Printf("%v",err)
	// }
	// fmt.Printf("Person %s\n",p.MonitorUnit.LocalIP)
	//LocalIP := p.MonitorUnit.LocalIP

	// jiankong.Send(readconfig.LocalIP, "this is a messgae for testing")
	// allocated, ubytes := 123, 456"github.com/spf13/viper"
	// newValue := fmt.Sprintf("too much data in sector: %d > %d", allocated, ubytes)
	// fmt.Println(newValue)

	type MonitorUnit struct {
		LocalIP string
		Timeout string
	}
	// 使用环境变量
	// useEnvViper()
	v := viper.New()
	v.SetConfigFile("config.toml")
	// v.AddConfigPath("/home/zky/go/src/learn_go/middleware/")
	// v.SetConfigType("toml")

	if err := v.ReadInConfig(); err != nil { // 搜索路径，并读取配置数据
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	subv := v.Sub("MonitorUnit")
	fmt.Println(subv.Get("LocalIP"), subv.Get("Timeout"))
	v.SetDefault("Address", "0.0.0.0:9090")
	err := v.WriteConfig() //写入文件
	if err != nil {        // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	var monitor MonitorUnit
	if err := subv.Unmarshal(&monitor); err != nil {
		fmt.Printf("err:%s", err)
	}

	fmt.Println(monitor)

	viper.SetConfigFile("/home/zky/go/src/learn_go/middleware/hello2.toml")
	// viper.SetConfigFile("$GOPATH/src/learn_go/middleware/hello2.toml")
	viper.SetDefault("MonitorUnit.Address", "0.0.0.0:9090")
	viper.Set("Address", "0.0.0.0:9090") //统一把Key处理成小写 Address->address
	viper.SetDefault("notifyList", []string{"xiaohong", "xiaoli", "xiaowang"})
	err = viper.WriteConfig() //写入文件
	if err != nil {           // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	viper.SetConfigFile("/home/zky/go/src/learn_go/middleware/hello2.toml")
	err = viper.ReadInConfig() // 会查找和读取配置文件
	if err != nil {            // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	Address := viper.GetString("Address")
	//key取Address或者address都能取到值，反正viper转成小写处理
	fmt.Println(Address)

	fmt.Println("=========================================================")
	useNameViper()
}

func useEnvViper() {
	prefix := "PROJECTNAME"
	envs := map[string]string{
		"LOG_LEVEL":      "INFO",
		"MODE":           "DEV",
		"MYSQL_USERNAME": "root",
		"MYSQL_PASSWORD": "xxxx",
	}
	for k, v := range envs {
		os.Setenv(fmt.Sprintf("%s_%s", prefix, k), v)
	}

	v1 := viper.New()
	v1.SetEnvPrefix(prefix)
	v1.AutomaticEnv()

	for k, _ := range envs {
		fmt.Printf("env `%s` = %s\n", k, v1.GetString(k))
	}
}

func useNameViper() {
	viper.SetConfigName("hello4")
	// 该种用法必须要先创建出该配置文件才行，否则会报出无法找到该文件的错误
	viper.AddConfigPath("$LOTUS_STORAGE_PATH/")
	// viper.AddConfigPath(".")
	viper.SetConfigType("toml")
	viper.SetDefault("MonitorUnit.Address", "0.0.0.0:9090")
	viper.Set("Address", "0.0.0.0:9090") //统一把Key处理成小写 Address->address
	viper.SetDefault("notifyList", []string{"xiaohong", "xiaoli", "xiaowang"})
	err := viper.WriteConfig() //写入文件
	if err != nil {            // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	viper.SetConfigFile("hello3.toml")
	err = viper.ReadInConfig() // 会查找和读取配置文件
	if err != nil {            // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	Address := viper.GetString("Address")
	//key取Address或者address都能取到值，反正viper转成小写处理
	fmt.Println(Address)
}
