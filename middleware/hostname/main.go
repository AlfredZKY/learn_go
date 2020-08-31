package main

import (
	"fmt"
	"reflect"
	"strconv"
	"sync"

	"github.com/spf13/viper"
)

var (
	lock   sync.Mutex
	global int32
)

// MonitorUnit 监控
type MonitorUnit struct {
	Ap  taskAP
	Pc1 taskPC1
	Pc2 taskPC2
	C1  taskC1
	C2  taskC2
}

type taskAP struct {
	hostnameValuePair map[string]int
}

type taskPC1 struct {
	hostnameValuePair map[string]int
}

type taskPC2 struct {
	hostnameValuePair map[string]int
}

type taskC1 struct {
	hostnameValuePair map[string]int
}

type taskC2 struct {
	hostnameValuePair map[string]int
}

// CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go
func main() {
	v := viper.New()
	v.SetConfigFile("config.toml")
	// v.AddConfigPath("/home/zky/go/src/learn_go/middleware/")
	// v.SetConfigType("toml")

	if err := v.ReadInConfig(); err != nil { // 搜索路径，并读取配置数据
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	// subv := v.Sub("MonitorUnit")
	subAp := v.Sub("Ap")
	homeList := subAp.Get("notifylist")
	valueList := subAp.Get("value")
	fmt.Println(homeList, reflect.TypeOf(homeList), valueList)

	fmt.Println(subAp.GetBool("default"), reflect.TypeOf(subAp.GetBool("default")))
	fmt.Println("=========================================================")
	var apTask taskAP = taskAP{hostnameValuePair:make(map[string]int,50)}

	var mont MonitorUnit = MonitorUnit{}
	
	switch reflect.TypeOf(homeList).Kind() {
	case reflect.Slice, reflect.Array:
		hostList := reflect.ValueOf(homeList)
		valueList := reflect.ValueOf(valueList)
		for i := 0; i < hostList.Len(); i++ {
			numberTmp := fmt.Sprintf("%v", valueList.Index(i))
			fmt.Println(numberTmp, reflect.TypeOf(numberTmp))
			number,_ := strconv.Atoi(numberTmp)
			hostname := fmt.Sprintf("%s", hostList.Index(i))
			fmt.Println(hostname, reflect.TypeOf(hostname))
			fmt.Println(number, reflect.TypeOf(number))
			apTask.hostnameValuePair[hostname] = number

		}
	case reflect.String:
		s := reflect.ValueOf(homeList)
		fmt.Println(s.String(), "I am a string type variable.")
	case reflect.Int:
		s := reflect.ValueOf(homeList)
		t := s.Int()
		fmt.Println(t, " I am a int type variable.")
	}
	// newHost := homeList.(reflect.Array)
	// valueHost := homeList.(reflect.Array)
	// var mont *MonitorUnit = &MonitorUnit{}
	// for index, value := range newHost {
	// 	mont.Ap.hostnameValuePair[value] = valueHost[index]
	// }
	mont.Ap = ap_task
	for index,value := range mont.Ap.hostnameValuePair {
		fmt.Println(index,value)
	}

	subPc1 := v.Sub("pc1")
	fmt.Println(subPc1.Get("notifylist"), subPc1.Get("value"))
	fmt.Println(subPc1.GetBool("default"), reflect.TypeOf(subPc1.GetBool("default")))
	fmt.Println("=========================================================")
}
