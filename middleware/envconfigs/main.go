package main

// 设置环境变脸
//export MYAPP_DEBUG=false
//export MYAPP_PORT=8080
//export MYAPP_USER=Kelsey
//export MYAPP_RATE="0.5"
//export MYAPP_TIMEOUT="3m"
//export MYAPP_USERS="rob,ken,robert"
//export MYAPP_COLORCODES="red:1,green:2,blue:3"

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"log"
	"time"
)

type Specification struct {
	Debug      bool
	Port       int
	User       string
	Users      []string
	Rate       float32
	Timeout    time.Duration
	ColorCodes map[string]int
}

func main() {
	var s Specification
	// 根据前缀找到所有环境换辆
	err := envconfig.Process("myapp", &s)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println(s)
	format := "Debug:%v\nPort:%v\nUser:%v\nUsers:%v\nRate:%v\nTimeout:%v\nColorCodes:%v\n"
	_, err = fmt.Printf(format, s.Debug, s.Port, s.User, s.Users, s.Rate, s.Timeout, s.ColorCodes)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Users:")
	for _, u := range s.Users {
		fmt.Printf(" %v\n", u)
	}

	fmt.Println("Color Codes:")
	for k, v := range s.ColorCodes {
		fmt.Printf(" %s:%d\n", k, v)
	}
}
