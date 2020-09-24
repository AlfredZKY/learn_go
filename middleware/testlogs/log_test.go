package testlogs

import (
	"learn_go/logger"
	"strconv"
	"testing"
	"time"
)

var (
	nsLogPath          = "./ns/sectors_declare"
	faultLogFileName   = "_fault_sectors_"
	skipLogFileName    = "_skipped_sectors_"
	recoverLogFileName = "_recovered_sectors_"

	nsWindowPostLogPath = "/opt/ns/windowPost/"
	windowPostFail      = "_windowPostFail_"
	windowPostSuccess   = "_windowPostSuccess_"

	windowPostInfo = "windowInfo.log"
)

func sectorLog(logLevel string, sectors []uint64, index uint64, err error) {
	SectorStatusLogPath := nsLogPath + "/" + time.Now().Format("2006-01-02_15:04:05")[:10]
	timePrefix := time.Now().Format("2006-01-02_15:04:05")
	indexStr := string(strconv.Itoa(int(index)))
	indexStr = indexStr + ".log"
	if logLevel == faultLogFileName {
		logger.DebugWithFilePath(SectorStatusLogPath+"/"+timePrefix+faultLogFileName+indexStr, "len is %v data is %v \n", len(sectors), sectors)
	} else if logLevel == skipLogFileName {
		logger.DebugWithFilePath(SectorStatusLogPath+"/"+timePrefix+skipLogFileName+indexStr, "len is %v data is %v \n", len(sectors), sectors)
	} else if logLevel == recoverLogFileName {
		logger.DebugWithFilePath(SectorStatusLogPath+"/"+timePrefix+recoverLogFileName+indexStr, "len is %v data is %v \n", len(sectors), sectors)
	} else if logLevel == windowPostFail {
		SectorStatusLogPath := nsWindowPostLogPath + "/" + time.Now().Format("2006-01-02_15:04:05")[:10]
		logger.DebugWithFilePath(SectorStatusLogPath+"/"+timePrefix+windowPostFail+indexStr, "submitPost failed: deadline is %v and err is %+v \n", index, err)
	} else if logLevel == windowPostSuccess {
		SectorStatusLogPath := nsWindowPostLogPath + "/" + time.Now().Format("2006-01-02_15:04:05")[:10]
		logger.DebugWithFilePath(SectorStatusLogPath+"/"+timePrefix+windowPostSuccess+indexStr, "window post submit successfully %v \n", "")
	}
}

func TestLogs(t *testing.T) {
	var sec []uint64
	var a uint64
	var err error
	sectorLog(faultLogFileName, sec, a, err)

}
