package main

import (
	"fmt"
	"learn_go/logger"
	"os"
	"reflect"
	"strconv"
	"sync"

	"github.com/spf13/viper"
)

var (
	apTask  taskPair = taskPair{workerRemainingMap: make(map[string]int, 50)}
	pc1Task taskPair = taskPair{workerRemainingMap: make(map[string]int, 50)}
	pc2Task taskPair = taskPair{workerRemainingMap: make(map[string]int, 50)}
	c1Task  taskPair = taskPair{workerRemainingMap: make(map[string]int, 50)}
	c2Task  taskPair = taskPair{workerRemainingMap: make(map[string]int, 50)}

	SchedLogPath = "/opt/ns"

	//taskWorkerRecord = map[string]taskPair{"ap": apTaskBak, "pc1": pc1TaskBak, "pc2": pc2TaskBak, "c1": c1TaskBak, "c2": c2TaskBak}

	taskTypes = map[string]taskPair{"ap": apTask, "pc1": pc1Task, "pc2": pc2Task, "c1": c1Task, "c2": c2Task}

	scheduleTask = map[string]taskPair{"seal/v0/addpiece": apTask, "seal/v0/precommit/1": pc1Task, "seal/v0/precommit/2": pc2Task, "seal/v0/commit/1": c1Task, "seal/v0/commit/2": c2Task}

	scheduleTaskMaps = make(map[string]sync.Map, 1000)
)

type taskPair struct {
	workerRemainingMap map[string]int
}

// LOTUS_MINER_PATH config path
// const LOTUS_MINER_PATH = "/home/qh/zhou_project"

func getSysPathEnv() string {
	var LotusMinerPath string
	LotusMinerPath = os.Getenv("TEST_MINER")
	return LotusMinerPath
}

func init() {
	initConfig()
}

func initConfig() {
	v := viper.New()
	//"v.SetConfigFile(getSysPathEnv() + "/worker_task_config.toml")",

	if err := v.ReadInConfig(); err != nil { // 搜索路径，并读取配置数据
		// TODO 读取配置文件错误
		logger.DebugWithFilePath(SchedLogPath+"/read_conifg_err.log", "Fatal error config file: %s \n", err)
	} else {
		for taskName, taskType := range taskTypes {
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
				taskType.workerRemainingMap[hostname] = number
			} else {
				taskType.workerRemainingMap[hostname] = defaultVaule
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

func updateWorkerFromTaskWorkerMap(tasktype, hostname string, number int) {
	logger.DebugWithFilePath(SchedLogPath+"/update_worker.log", "removing worker to map: %v\n", hostname)
	for taskname, taskPairs := range taskTypes {
		if taskname == tasktype {
			for hostnameexist := range taskPairs.workerRemainingMap {
				if hostnameexist == hostname {
					taskTypes[tasktype].workerRemainingMap[hostname] += number
				}
			}
		}
	}
}

func removeWorkerFromTaskWorkerMap(hostname string) {
	logger.DebugWithFilePath(SchedLogPath+"/remove_worker.log", "removing worker to map: %v\n", hostname)
	taskTypeList := []string{"seal/v0/addpiece", "seal/v0/precommit/1", "seal/v0/precommit/2", "seal/v0/commit/1", "seal/v0/commit/2"}
	for i := 0; i < len(taskTypeList); i++ {
		var securityMap sync.Map
		securityMap = scheduleTaskMaps[taskTypeList[i]]
		securityMap.Delete(hostname)
	}
}

func addWorkerToTaskWorkerRemaining(hostname string) {
	logger.DebugWithFilePath(SchedLogPath+"/new_worker.log", "trying to add new worker: %v\n", hostname)
	var exist = false
	for _, taskPairs := range taskTypes {
		if taskPairs.workerRemainingMap[hostname] > 0 {
			exist = true
		}
	}

	if !exist {
		logger.DebugWithFilePath(SchedLogPath+"/new_worker.log", "%v is not in current record, re-initializing record from config file...\n", hostname)
		initConfig()
	}

	logger.DebugWithFilePath(SchedLogPath+"/new_worker.log", "trying to add new worker to map: %v\n", hostname)
	for tasktype, taskWorkerPair := range taskTypes {
		var securityMap sync.Map
		srcValue := taskWorkerPair.workerRemainingMap[hostname]
		if tasktype == "ap" {
			if srcValue > 0 {
				securityMap.Store(hostname, srcValue)
				scheduleTaskMaps["seal/v0/addpiece"] = securityMap
			}
		} else if tasktype == "pc1" {
			if srcValue > 0 {
				securityMap.Store(hostname, srcValue)
				scheduleTaskMaps["seal/v0/precommit/1"] = securityMap
			}
		} else if tasktype == "pc2" {
			if srcValue > 0 {
				securityMap.Store(hostname, srcValue)
				scheduleTaskMaps["seal/v0/precommit/2"] = securityMap
			}
		} else if tasktype == "c1" {
			if srcValue > 0 {
				securityMap.Store(hostname, srcValue)
				scheduleTaskMaps["seal/v0/commit/1"] = securityMap
			}
		} else {
			if srcValue > 0 {
				securityMap.Store(hostname, srcValue)
				scheduleTaskMaps["seal/v0/commit/2"] = securityMap
			}
		}
	}
}

func parseSyncMap(key, hostname string) int {
	var securityMap sync.Map
	securityMap = scheduleTaskMaps[key]
	if temptask, ok := securityMap.Load(hostname); ok {
		currentVaule := temptask.(int)
		// currentVaule := taskHostPair[hostname]
		if currentVaule > 0 {
			return currentVaule
		}
	}
	return -1
}

func main() {
	addWorkerToTaskWorkerRemaining("idc24")
	addWorkerToTaskWorkerRemaining("idc25")
	fmt.Println(scheduleTaskMaps)
	currentValue := parseSyncMap("seal/v0/commit/1", "idc24")
	fmt.Println(currentValue)
	currentValue1 := parseSyncMap("seal/v0/precommit/1", "idc25")
	fmt.Println(currentValue1)
	removeWorkerFromTaskWorkerMap("idc24")
	currentValue2 := parseSyncMap("seal/v0/commit/1", "idc24")
	fmt.Println(currentValue2)
	fmt.Println(scheduleTaskMaps)

	fmt.Println("===========================")
	ss := make(map[string]sync.Map, 100)
	var s1 sync.Map
	s1.Store("hello", 0)
	ss["ap"] = s1
	s2 := ss["pc1"]
	if val, ok := s2.Load("sssss"); ok {
		fmt.Println(val)
	}
	fmt.Println(reflect.TypeOf(s2))

	//taskTypeList := []string{"seal/v0/addpiece", "seal/v0/precommit/1", "seal/v0/precommit/2", "seal/v0/commit/1", "seal/v0/commit/2"}
	//hostnamelist := []string{"worker-1-148-gpu","worker-1-147-gpu","worker-2-56","worker-2-57","worker_0_198_gpu","worker_0_199_gpu","worker_0_200_gpu","worker_0_201_gpu","worker_0_202_gpu","worker-31-2-gpu"}
	//for i:=0;i<len(taskTypeList);i++{
	//	tempMap := ss[taskTypeList[i]]
	//	for i := 0; i < len(hostnamelist);i++{
	//		if value,ok := tempMap.Load()
	//	}
	//}
}
