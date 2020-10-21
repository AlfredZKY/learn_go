package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"log"
	"reflect"
	"time"
)

type tomlConfig struct {
	Title   string
	Owner   ownerInfo
	DB      database `toml:"database"`
	Servers map[string]server
	Clients clients
}

type ownerInfo struct {
	Name string
	Org  string `toml:"organization"`
	Bio  string
	Dob  time.Time
}

type database struct {
	Server  string
	Ports   []int
	ConnMax int `toml:"connection_max"`
	Enabled bool
}

type server struct {
	IP string
	DC string
}

type clients struct {
	Data  [][]interface{}
	Hosts []string
}

type tomlConfigs struct {
	Title   string
	Miners miner
	AP ap
	PC1 pc1
	PC2 pc2
	C1 c1
	C2 c2
}

type miner struct {
	HostName []string
	Value []int
	Default int
}

type ap struct {
	HostName []string
	Value []int
	Default int
}

type pc1 struct {
	HostName []string
	Value []int
	Default int
}
type pc2 struct {
	HostName []string
	Value []int
	Default int
}

type c1 struct {
	HostName []string
	Value []int
	Default int
}
type c2 struct {
	HostName []string
	Value []int
	Default int
}

func loadTomlConfig(){
	var config tomlConfigs
	if _,err := toml.DecodeFile("worker_task_config.toml",&config);err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(config.AP)
}

func testOriginConfig(){
	var config tomlConfig
	if _, err := toml.DecodeFile("config.toml", &config); err != nil {
		log.Fatal(err.Error())
	}

	log.Printf("Title: %s\n", config.Title)

	log.Printf("Owner: %s (%s, %s), Born: %s\n ", config.Owner.Name, config.Owner.Org, config.Owner.Bio, config.Owner.Dob)

	log.Printf("Database: %s %v (Max conn. %d),Enabled? %v\n", config.DB.Server, config.DB.Ports, config.DB.ConnMax, config.DB.Enabled)

	for serverName, server := range config.Servers {
		fmt.Printf("Server: %s (%s , %s)\n", serverName, server.IP, server.DC)
	}

	log.Printf("Clients data: %v type is %v\n", config.Clients.Data, reflect.TypeOf(config.Clients.Data))
	log.Printf("Clients hosts %v type is %v\n", config.Clients.Hosts, reflect.TypeOf(config.Clients.Hosts))

	fmt.Println("-------------------------")
}

func main() {
	testOriginConfig()
	loadTomlConfig()
	testStr := []string{"hi","hello"}
	fmt.Println(testStr)
}
