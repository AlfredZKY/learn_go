package main

import (
	"fmt"
	"learn_go/logger"
	"reflect"
	"strconv"
	"sync"

	"github.com/spf13/viper"
)

var (
	apTask       taskAP      = taskAP{hostnameValuePair: make(map[string]int, 50)}
	pc1Task      taskAP      = taskAP{hostnameValuePair: make(map[string]int, 50)}
	pc2Task      taskAP      = taskAP{hostnameValuePair: make(map[string]int, 50)}
	c1Task       taskAP      = taskAP{hostnameValuePair: make(map[string]int, 50)}
	c2Task       taskAP      = taskAP{hostnameValuePair: make(map[string]int, 50)}
	mont         MonitorUnit = MonitorUnit{}
	lock         sync.Mutex
	global       int32
	SchedLogPath = "/opt/ns"
	taskTypes    = map[string]taskAP{"ap": apTask, "pc1": pc1Task, "pc2": pc2Task, "c1": c1Task, "c2": c2Task}
)

// MonitorUnit 监控
type MonitorUnit struct {
	Ap  taskAP
	Pc1 taskAP
	Pc2 taskAP
	C1  taskAP
	C2  taskAP
}

type taskAP struct {
	hostnameValuePair map[string]int
}

// LOTUS_MINER_PATH config path
const LOTUS_MINER_PATH = "/home/zky/project/"

func init() {
	v := viper.New()
	v.SetConfigFile(LOTUS_MINER_PATH + "worker_task_config.toml")

	if err := v.ReadInConfig(); err != nil { // 搜索路径，并读取配置数据
		// TODO 读取配置文件错误
		logger.DebugWithFilePath(SchedLogPath+"/read_conifg_err.log", "Fatal error config file: %s \n", err)
	}
	for taskName, taskType := range taskTypes {
		parseConfigToml(taskName, v, &taskType)
	}
	mont.Ap = apTask
	mont.Pc1 = pc1Task
	mont.Pc2 = pc2Task
	mont.C1 = c1Task
	mont.C2 = c2Task
}

func parseConfigToml(key string, v *viper.Viper, taskType *taskAP) {
	subAp := v.Sub(key)
	homeList := subAp.Get("notifylist")
	valueList := subAp.Get("value")
	defaultVaule := subAp.GetInt("default")

	switch reflect.TypeOf(homeList).Kind() {
	case reflect.Slice, reflect.Array:
		hostList := reflect.ValueOf(homeList)
		valueList := reflect.ValueOf(valueList)
		for i := 0; i < hostList.Len(); i++ {
			numberTmp := fmt.Sprintf("%v", valueList.Index(i))
			number, _ := strconv.Atoi(numberTmp)
			hostname := fmt.Sprintf("%s", hostList.Index(i))
			if defaultVaule < 0 {
				taskType.hostnameValuePair[hostname] = number
			} else {
				taskType.hostnameValuePair[hostname] = defaultVaule
			}
		}
	case reflect.String:
		s := reflect.ValueOf(homeList)
		fmt.Println(s.String(), "I am a string type variable.")
	case reflect.Int:
		s := reflect.ValueOf(homeList)
		t := s.Int()
		fmt.Println(t, " I am a int type variable.")
	}
}

// CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go
func main() {
	for index, value := range mont.Ap.hostnameValuePair {
		fmt.Println(index, value)
	}

}
