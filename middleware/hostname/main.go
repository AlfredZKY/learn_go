package main

import (
	"encoding/json"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/fsnotify/fsnotify"
	"learn_go/logger"
	"log"
	"os"
	"reflect"
	"strconv"
	"sync"
	"time"

	//"github.com/filecoin-project/go-state-types/dline"
	"github.com/spf13/viper"
)

var (
	faultLogFileName      = " fault_sectors_"
	skipLogFileName       = " skip_sectors_"
	recoverLogFileName    = " recovered_sectors_"
	windowPostLogFileName = " windowpost_"
)

var (
	minerWorkerRecord taskPair = taskPair{workerRemainingMap: make(map[string]int, 50)}
	apWorkerRecord    taskPair = taskPair{workerRemainingMap: make(map[string]int, 50)}
	pc1WorkerRecord   taskPair = taskPair{workerRemainingMap: make(map[string]int, 50)}
	pc2WorkerRecord   taskPair = taskPair{workerRemainingMap: make(map[string]int, 50)}
	c1WorkerRecord    taskPair = taskPair{workerRemainingMap: make(map[string]int, 50)}
	c2WorkerRecord    taskPair = taskPair{workerRemainingMap: make(map[string]int, 50)}

	minerWorkerNew taskPair = taskPair{workerRemainingMap: make(map[string]int, 50)}
	apWorkerNew    taskPair = taskPair{workerRemainingMap: make(map[string]int, 50)}
	pc1WorkerNew   taskPair = taskPair{workerRemainingMap: make(map[string]int, 50)}
	pc2WorkerNew   taskPair = taskPair{workerRemainingMap: make(map[string]int, 50)}
	c1WorkerNew    taskPair = taskPair{workerRemainingMap: make(map[string]int, 50)}
	c2WorkerNew    taskPair = taskPair{workerRemainingMap: make(map[string]int, 50)}

	SchedLogPath     = "./logs"
	taskWorkerRecord = map[string]taskPair{"miner": minerWorkerRecord, "ap": apWorkerRecord, "pc1": pc1WorkerRecord, "pc2": pc2WorkerRecord, "c1": c1WorkerRecord, "c2": c2WorkerRecord}
	taskNewRecord    = map[string]taskPair{"miner": minerWorkerNew, "ap": apWorkerNew, "pc1": pc1WorkerNew, "pc2": pc2WorkerNew, "c1": c1WorkerNew, "c2": c2WorkerNew}
	scheduleTaskMaps = make(map[string]sync.Map, 1000)

	taskNameMapping = map[string]string{"miner": "miner", "ap": "seal/v0/addpiece", "pc1": "seal/v0/precommit/1", "pc2": "seal/v0/precommit/2", "c1": "seal/v0/commit/1", "c2": "seal/v0/commit/2"}
	Locks           sync.RWMutex
)

type tomlConfig struct {
	Miner miner
	AP ap
	PC1 pc1
	PC2 pc2
	C1 c1
	C2 c2
}

type miner struct {
	Hostnames []string
	Value []int
	Default int
	TaskName string
}

type ap struct {
	Hostnames []string
	Value []int
	Default int
	TaskName string
}

type pc1 struct {
	Hostnames []string
	Value []int
	Default int
	TaskName string
}
type pc2 struct {
	Hostnames []string
	Value []int
	Default int
	TaskName string
}

type c1 struct {
	Hostnames []string
	Value []int
	Default int
	TaskName string
}
type c2 struct {
	Hostnames []string
	Value []int
	Default int
	TaskName string
}

func parseConfigTomlNew(config *tomlConfig,temp interface{}){
	mapData := temp.(map[string]taskPair)
	for i:=0;i < len(config.Miner.Hostnames);i++{
		if config.Miner.Default  < 0 {
			mapData[config.Miner.TaskName].workerRemainingMap[config.Miner.Hostnames[i]] = config.Miner.Value[i]
		}else{
			mapData[config.Miner.TaskName].workerRemainingMap[config.Miner.Hostnames[i]] = config.Miner.Default
		}
	}
	for i:=0;i < len(config.AP.Hostnames);i++{
		if config.AP.Default  < 0 {
			mapData[config.AP.TaskName].workerRemainingMap[config.AP.Hostnames[i]] = config.AP.Value[i]
		}else{
			mapData[config.AP.TaskName].workerRemainingMap[config.AP.Hostnames[i]] = config.AP.Default
		}
	}

	for i:=0;i < len(config.PC1.Hostnames);i++{
		if config.PC1.Default  < 0 {
			mapData[config.PC1.TaskName].workerRemainingMap[config.PC1.Hostnames[i]] = config.PC1.Value[i]
		}else{
			mapData[config.PC1.TaskName].workerRemainingMap[config.PC1.Hostnames[i]] = config.PC1.Default
		}
	}

	for i:=0;i < len(config.PC2.Hostnames);i++{
		if config.PC2.Default  < 0 {
			mapData[config.PC2.TaskName].workerRemainingMap[config.PC2.Hostnames[i]] = config.PC2.Value[i]
		}else{
			mapData[config.PC2.TaskName].workerRemainingMap[config.PC2.Hostnames[i]] = config.PC2.Default
		}
	}
	for i:=0;i < len(config.C1.Hostnames);i++{
		if config.C1.Default  < 0 {
			mapData[config.C1.TaskName].workerRemainingMap[config.C1.Hostnames[i]] = config.C1.Value[i]
		}else{
			mapData[config.C1.TaskName].workerRemainingMap[config.C1.Hostnames[i]] = config.C1.Default
		}
	}
	for i:=0;i < len(config.C2.Hostnames);i++{
		if config.C2.Default  < 0 {
			mapData[config.C2.TaskName].workerRemainingMap[config.C2.Hostnames[i]] = config.C2.Value[i]
		}else{
			mapData[config.C2.TaskName].workerRemainingMap[config.C2.Hostnames[i]] = config.C2.Default
		}
	}
}

func loadTomlConfig(temp interface{}){
	var config tomlConfig
	if _,err := toml.DecodeFile("worker_task_config.toml",&config);err != nil {
		log.Fatal(err.Error())
	}
	parseConfigTomlNew(&config,temp)
}

type taskPair struct {
	workerRemainingMap map[string]int
}

// LOTUS_MINER_PATH config path
// const LOTUS_MINER_PATH = "/home/qh/zhou_project"

func sectorLog(logLevel string, sectors []int, index uint64, err error) {
	SectorStatusLogPath := SchedLogPath + "/" + time.Now().Format("2006-01-02 15:04:05")[:10]
	indexstr := string(strconv.Itoa(int(index)))
	indexstr = indexstr + ".log"
	timePrefix := time.Now().Format("2006-01-02 15:04:05")
	if logLevel == faultLogFileName {
		logger.DebugWithFilePath(SectorStatusLogPath+"/"+timePrefix+faultLogFileName+indexstr, "len is %V data is %v \n", len(sectors), sectors)
	} else if logLevel == skipLogFileName {
		logger.DebugWithFilePath(SectorStatusLogPath+"/"+timePrefix+skipLogFileName+indexstr, "len is %V data is %v \n", len(sectors), sectors)
	} else if logLevel == recoverLogFileName {
		logger.DebugWithFilePath(SectorStatusLogPath+"/"+timePrefix+recoverLogFileName+indexstr, "len is %V data is %v \n", len(sectors), sectors)
	} else if logLevel == windowPostLogFileName {
		logger.DebugWithFilePath(SectorStatusLogPath+"/"+timePrefix+recoverLogFileName+indexstr, "submitPost failed: deadline is %v and err is %+v \n", index, err)
	} else {
		logger.DebugWithFilePath(SectorStatusLogPath+"/"+timePrefix+windowPostLogFileName+indexstr, "window post submit succssfully %v \n", "")
	}
}

func getSysPathEnv() string {
	var LotusMinerPath string
	LotusMinerPath = os.Getenv("LOTUS_MINER_PATH")
	return LotusMinerPath
}

func loadTaskConfigOld() {
	logger.DebugWithFilePath(SchedLogPath+"/new_schedule.log", "Loading worker_task_config.toml %v\n", "")
	v := viper.New()
	//v.SetConfigFile(getSysPathEnv() + "./worker_task_config.toml")
	v.SetConfigFile("./worker_task_config.toml")

	if err := v.ReadInConfig(); err != nil { // 搜索路径，并读取配置数据
		// TODO 读取配置文件错误
		logger.DebugWithFilePath(SchedLogPath+"/new_schedule.log", "Fatal error config file: %s \n", err)
	}else{
		for taskName, taskType := range taskWorkerRecord {
			parseConfigToml(taskName, v, &taskType)
		}
		logger.DebugWithFilePath(SchedLogPath+"/new_schedule.log", "config content: %v\n", taskWorkerRecord)
	}
	logger.DebugWithFilePath(SchedLogPath+"/new_schedule.log", "Loading worker_task_config.toml successful! %v\n", "")
}

func notifyConfigChange(){
	v := viper.New()
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event){
		fmt.Println("config file changed:",e.Name)
	})
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
	}
}

func removeWorkerFromTaskWorkerRemaining(hostname string) {
	logger.DebugWithFilePath(SchedLogPath+"/new_schedule_remove.log", "removing worker from map: %v\n", hostname)
	taskTypeList := []string{"seal/v0/addpiece", "seal/v0/precommit/1", "seal/v0/precommit/2", "seal/v0/commit/1", "seal/v0/commit/2"}
	for i := 0; i < len(taskTypeList); i++ {
		var securityMap sync.Map
		securityMap = scheduleTaskMaps[taskTypeList[i]]
		securityMap.Delete(hostname)
	}
	logger.DebugWithFilePath(SchedLogPath+"/new_schedule_remove.log", "removing worker from map done: %v\n", hostname)
}

func checkWorkerExistence(hostname string) bool {
	logger.DebugWithFilePath(SchedLogPath+"/new_schedule.log", "checking worker existence: %v\n", hostname)
	var exist = false
	for _, taskPairs := range taskWorkerRecord {
		if taskPairs.workerRemainingMap[hostname] != 0 {
			exist = true
		}
	}
	if exist {
		logger.DebugWithFilePath(SchedLogPath+"/new_schedule.log", "This worker exists: %v\n", hostname)
	} else {
		logger.DebugWithFilePath(SchedLogPath+"/new_schedule.log", "This worker does not exist: %v\n", hostname)
	}
	return exist
}

func addWorkerToTaskWorkerRemaining(hostname string) {
	i := 0
	for !checkWorkerExistence(hostname) {
		i = i + 1
		logger.DebugWithFilePath(SchedLogPath+"/new_schedule.log", "%v is not in current record, re-initializing record from config file...\n", hostname)
		//loadTomlConfig(taskWorkerRecord)
		loadTaskConfigOld()
		logger.DebugWithFilePath(SchedLogPath+"/new_schedule.log", "loading config done! %v\n", "")
		if i > 5 {
			logger.DebugWithFilePath(SchedLogPath+"/check_config.log", "Check config file for : %v\n", hostname)
			break
		}
	}

	logger.DebugWithFilePath(SchedLogPath+"/new_schedule.log", "trying to add new worker to map: %v\n", hostname)
	for tasktype, taskWorkerPair := range taskWorkerRecord {
		//logger.DebugWithFilePath(SchedLogPath+"/new_schedule.log", "\t Inside for loop: %v\n", "")
		srcValue, _ := taskWorkerPair.workerRemainingMap[hostname]

		//logger.DebugWithFilePath(SchedLogPath+"/new_schedule.log", "\t source value: %v\n", srcValue)

		_, ok := scheduleTaskMaps[taskNameMapping[tasktype]]
		if !ok {
			var tmpSyncMap sync.Map
			tmpSyncMap.Store(123, 456)
			scheduleTaskMaps[taskNameMapping[tasktype]] = tmpSyncMap

			tmpSyncMap.Delete(123)
		}

		curSyncMap, ok := scheduleTaskMaps[taskNameMapping[tasktype]]
		if !ok {
			logger.DebugWithFilePath(SchedLogPath+"/new_schedule.log", "\t Can not get curSyncMap for %v: %v\n", tasktype, scheduleTaskMaps)
			return
		}

		if srcValue != 0 {
			curSyncMap.Store(hostname, srcValue)
		}
	}

	logger.DebugWithFilePath(SchedLogPath+"/new_schedule.log", "adding new worker to map done: %v\n", hostname)
	for tasktype, securityMap := range scheduleTaskMaps {
		logger.DebugWithFilePath(SchedLogPath+"/new_schedule.log", "Current remaining map for %v: %v\n", tasktype, stringfySyncMap(&securityMap))
	}

}

func stringfySyncMap(amap *sync.Map) string {
	m := map[string]interface{}{}
	amap.Range(func(key, value interface{}) bool {
		m[fmt.Sprint(key)] = value
		return true
	})

	b, err := json.MarshalIndent(m, "", " ")
	if err != nil {
		panic(err)
	}
	return string(b)
}

// 根据主机名 - hostname，返回还可以做任务的数目
// 如果是mer，返回-10000
// 如果是未分类的任务，返回-20000
func parseSyncMap(task, hostname string) int {
	var securityMap sync.Map

	securityMap, ok := scheduleTaskMaps[task]
	if !ok {
		// unclassifed task marker as -20000
		return -20000
	}
	if temptask, ok := securityMap.Load(hostname); ok {
		currentVaule := temptask.(int)
		return currentVaule
	}
	return 0
}

func testLog() {
	var temp []int
	index := uint64(1)
	sectorLog("", temp, index, nil)
}


// UpdateRecordConfig
func UpdateRecordConfig() {
	fmt.Println("=================")
	time.Sleep(time.Second*60)
	loadTomlConfig(taskNewRecord)
	fmt.Println(taskNewRecord)

	for taskType, taskWorkerPair := range taskNewRecord {
		for hostname, _ := range taskWorkerPair.workerRemainingMap {
			Locks.Lock()
			currentValue := parseSyncMap(taskNameMapping[taskType], hostname)
			fmt.Println(currentValue,taskNewRecord[taskType].workerRemainingMap[hostname],taskWorkerRecord[taskType].workerRemainingMap[hostname])
			newNumber := taskNewRecord[taskType].workerRemainingMap[hostname] - taskWorkerRecord[taskType].workerRemainingMap[hostname]
			fmt.Println(newNumber)

			for true{
				securityMap,ok := scheduleTaskMaps[taskNameMapping[taskType]]
				fmt.Println("test ok: ",ok)
				if ok {
					securityMap.Store(hostname, currentValue + newNumber)
					break
				}
			}
			fmt.Println("!!!! should be: ", currentValue + newNumber)
			Locks.Unlock()
		}
	}
	// TODO x'
	//taskWorkerRecord, taskNewRecord = taskNewRecord, taskWorkerRecord

	fmt.Println("======================================================")
	for taskType, securityMap := range scheduleTaskMaps {
		logger.DebugWithFilePath(SchedLogPath+"/reload_task_config.log", "Current remaining map for %v: %v\n", taskType, stringfySyncMap(&securityMap))
	}
	logger.DebugWithFilePath(SchedLogPath+"/reload_task_config.log", "\n\n Previous config map is: %v \n\n Current config map is %v\n\n", taskNewRecord, taskWorkerRecord)

}

func main() {
	//loadTomlConfig(taskWorkerRecord)
	addWorkerToTaskWorkerRemaining("xiaohong")
	//addWorkerToTaskWorkerRemaining("miner")
	//fmt.Println(scheduleTaskMaps)
	//time.Sleep(60*time.Second)
	//currentValue := parseSyncMap("seal/v0/commit/1", "xiaohong")
	//fmt.Println(currentValue)
	//currentValue1 := parseSyncMap("seal/v0/precommit/1", "xiaohong")
	//fmt.Println(currentValue1)
	//removeWorkerFromTaskWorkerRemaining("xiaohong")
	//currentValue2 := parseSyncMap("seal/v0/commit/1", "xiaohong")
	//fmt.Println(currentValue2)
	//for tasktypes,securityMap := range scheduleTaskMaps{
	//	fmt.Println( tasktypes, securityMap)
	//}
	//ss := make(map[string]sync.Map, 100)
	//var s1 sync.Map
	//s1.Store("hello", 0)
	//ss["ap"] = s1
	//s2 := ss["pc1"]
	//if val, ok := s2.Load("sssss"); ok {
	//	fmt.Println(val)
	//}
	//fmt.Println(reflect.TypeOf(s2))
	UpdateRecordConfig()
	//
	//for i:=0; i < 5; i++ {
	//	UpdateRecordConfig()
	//	time.Sleep(10*time.Second)
	//}

	//for tasktype, securityMap := range scheduleTaskMaps {
	//	fmt.Println( tasktype, stringfySyncMap(&securityMap))
	//}
	//loadTaskConfigOld()
	//time.Sleep(time.Second*60)
	//loadTaskConfigNew()
}
