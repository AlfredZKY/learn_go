package conf

import (
	"fmt"
	"io/ioutil"
	"runtime"

	"gopkg.in/yaml.v2"
)

//位置文件结构

type Conf struct {
	Port int `yaml:"port"`
	Eos  struct {
		Host    string `yaml:"host"`
		Account string `yaml:"account"`
	} `yaml:"eos"`
	Accounts []struct {
		Key      string   `yaml:"key"`
		Name     string   `yaml:"name"`
		Betvalue []string `yaml:"betvalue"`
		Bettime  []int    `yaml:"bettime"`
	} `yaml:"accounts"`
	Randtime []int `yaml:"randtime"`
	Testnet  bool  `yaml:"testnet"`
}

/*
yaml文件内容如下：

port: 8088
eos:
  host: http://114.67.37.198:8000
  account: eosbullfight
accounts:
  - key: user1
    name: name1
    betvalue: ["0.5000 EOS","1.0000 EOS","5.0000 EOS"]
    bettime: [2] #[1,3,5,3]
  - key: user2
    name: name2
    betvalue: ["0.5000 EOS","1.0000 EOS","5.0000 EOS"] #下注金额范围随机取一个
    bettime: [1] #[1,1,5,6] #下注次数随机取一个
randtime: [100] #[200,120,30,254,89,99,45,66,73,82,102,113,188,122,156,57,133,77,145,144,166,185,174] #随机间隔时间

*/

func (c *Conf) GetConfig() *Conf {
	var path = "chain.yml"
	if runtime.GOOS == "windows" {
		path = `C:\` + path
	}

	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		fmt.Printf("Unmarshal: %v", err.Error())
	}
	// fmt.Println("yfile", string(yamlFile))
	return c
}

var installc *Conf

func init() {
	install()
}

func install() {
	if installc == nil {
		xx := &Conf{}
		installc = xx.GetConfig()
	}
}

func CCFrom() *Conf {
	install()
	return installc
}

func _CCC() {
	var _ = ""
}
