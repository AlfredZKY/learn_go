package main

import (
	"fmt"
	"learn_go/logger"
	"reflect"
	"strconv"

	"github.com/spf13/viper"
)

var (
	apTaskBak  taskPair = taskPair{hostnameValuePair: make(map[string]int, 50)}
	pc1TaskBak taskPair = taskPair{hostnameValuePair: make(map[string]int, 50)}
	pc2TaskBak taskPair = taskPair{hostnameValuePair: make(map[string]int, 50)}
	c1TaskBak  taskPair = taskPair{hostnameValuePair: make(map[string]int, 50)}
	c2TaskBak  taskPair = taskPair{hostnameValuePair: make(map[string]int, 50)}

	apTask  taskPair = taskPair{hostnameValuePair: make(map[string]int, 50)}
	pc1Task taskPair = taskPair{hostnameValuePair: make(map[string]int, 50)}
	pc2Task taskPair = taskPair{hostnameValuePair: make(map[string]int, 50)}
	c1Task  taskPair = taskPair{hostnameValuePair: make(map[string]int, 50)}
	c2Task  taskPair = taskPair{hostnameValuePair: make(map[string]int, 50)}

	SchedLogPath  = "/opt/ns"
	taskTypesBack = map[string]taskPair{"ap": apTaskBak, "pc1": pc1TaskBak, "pc2": pc2TaskBak, "c1": c1TaskBak, "c2": c2TaskBak}

	taskTypes    = map[string]taskPair{"ap": apTask, "pc1": pc1Task, "pc2": pc2Task, "c1": c1Task, "c2": c2Task}
	scheduleTask = map[string]taskPair{"seal/v0/addpiece": apTask, "seal/v0/precommit/1": pc1Task, "seal/v0/precommit/2": pc2Task, "seal/v0/commit/1": c1Task, "seal/v0/commit/2": c2Task}
)

type taskPair struct {
	hostnameValuePair map[string]int
}

// LOTUS_MINER_PATH config path
const LOTUS_MINER_PATH = "/home/qh/zhou_project/"

func init() {
	initConfig()
}

func initConfig() {
	v := viper.New()
	v.SetConfigFile(LOTUS_MINER_PATH + "worker_task_config.toml")

	if err := v.ReadInConfig(); err != nil { // 搜索路径，并读取配置数据
		// TODO 读取配置文件错误
		logger.DebugWithFilePath(SchedLogPath+"/read_conifg_err.log", "Fatal error config file: %s \n", err)
	} else {
		for taskName, taskType := range taskTypesBack {
			parseConfigToml(taskName, v, &taskType)
		}
	}

}

func parseConfigToml(key string, v *viper.Viper, taskType *taskPair) {
	subAp := v.Sub(key)
	homeList := subAp.Get("hostnames")
	valueList := subAp.Get("value")
	defaultVaule := subAp.GetInt("default")

	switch reflect.TypeOf(homeList).Kind() {
	case reflect.Slice, reflect.Array:
		hostList := reflect.ValueOf(homeList)
		valueList := reflect.ValueOf(valueList)
		for i := 0; i < hostList.Len(); i++ {
			hostname := fmt.Sprintf("%s", hostList.Index(i))
			if defaultVaule < 0 {
				numberTmp := fmt.Sprintf("%v", valueList.Index(i))
				number, _ := strconv.Atoi(numberTmp)
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

func removeWorkerFromTaskWorkerMap(hostname string) {
	logger.DebugWithFilePath(SchedLogPath+"/remove_worker.log", "removing worker to map: %v\n", hostname)
	for taskname, taskPairs := range taskTypesBack {
		for hostnameexist := range taskPairs.hostnameValuePair {
			if hostnameexist == hostname {
				delete(taskTypes[taskname].hostnameValuePair, hostname)
			}
		}
	}
}

func addWorkerFromTaskWorkerMap(hostname string) {
	logger.DebugWithFilePath(SchedLogPath+"/new_worker.log", "trying to add new worker: %v\n", hostname)
	var exist = false
	for _, taskPairs := range taskTypesBack {
		for hostnameexist := range taskPairs.hostnameValuePair {
			if hostnameexist == hostname {
				exist = true
			}
		}
	}
	if !exist {
		logger.DebugWithFilePath(SchedLogPath+"/new_worker.log", "%v is not in current record, re-initializing record from config file...\n", hostname)
		initConfig()
	}

	logger.DebugWithFilePath(SchedLogPath+"/new_worker.log", "trying to add new worker to map: %v\n", hostname)

	for taskname, taskPairs := range taskTypesBack {
		for hostnameexist, value := range taskPairs.hostnameValuePair {
			if hostnameexist == hostname {
				taskTypes[taskname].hostnameValuePair[hostnameexist] = value
				logger.DebugWithFilePath(SchedLogPath+"/new_worker.log", "this worker can do: %v %v\n", taskname, value)
			}
		}
	}
}

func main() {
	addWorkerFromTaskWorkerMap("idc21")
	fmt.Println("======================================")
	for _, taskPairs := range taskTypes {
		for hostname, value := range taskPairs.hostnameValuePair {
			fmt.Println(hostname, " ", value)
		}
	}
	removeWorkerFromTaskWorkerMap("idc23")
	fmt.Println("======================================")
	for _, taskPairs := range taskTypes {
		for hostname, value := range taskPairs.hostnameValuePair {
			fmt.Println(hostname, " ", value)
		}
	}
}
