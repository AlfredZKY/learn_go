package logger

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"
)

type Logger struct {
	FilePath string
	LogLevel int
	TestInt  int

	loggerLock sync.Mutex
}

var ColorClose = "\033[0m"
var logLevelColor = map[int]string{
	DebugLevel: "\033[32m",
	InfoLevel:  "\033[33m",
	ErrorLevel: "\033[31m",
}

var logLevelPrefix = map[int]string{
	DebugLevel: "DEBUG",
	InfoLevel:  "INFO",
	ErrorLevel: "ERROR",
}

const (
	DebugLevel = iota
	InfoLevel
	ErrorLevel
)

var logger Logger

func init() {
	logger = Logger{
		FilePath:   "./test.log",
		LogLevel:   DebugLevel,
		loggerLock: sync.Mutex{},
	}
}

func SetUp(logLevel int, filePath string) {
	logger.LogLevel = logLevel
	logger.FilePath = filePath
}

func Debug(format string, args ...interface{}) {
	DebugWithFilePath("", format, args...)
}

func DebugWithFilePath(filePath string, format string, args ...interface{}) {
	logger.PrintLog(DebugLevel, filePath, format, args...)
}

func Info(format string, args ...interface{}) {
	InfoWithFilePath("", format, args...)
}

func InfoWithFilePath(filePath string, format string, args ...interface{}) {
	logger.PrintLog(InfoLevel, filePath, format, args...)
}

func Error(format string, args ...interface{}) {
	ErrorWithFilePath("", format, args...)
}

func ErrorWithFilePath(filePath string, format string, args ...interface{}) {
	logger.PrintLog(ErrorLevel, filePath, format, args...)
}

func (l *Logger) PrintLog(logLevel int, filePath string, format string, args ...interface{}) {
	if l.LogLevel > logLevel {
		return
	}

	l.loggerLock.Lock()
	defer l.loggerLock.Unlock()

	curTime := time.Now()
	msg := fmt.Sprintf(format, args...)
	codeFile, codeLine := GetCallerInfo()
	logMsg := fmt.Sprintf("%s%s %s %d %s%s", l.getLogLevelColor(logLevel), l.getLogPrefix(logLevel, curTime), codeFile, codeLine, msg, ColorClose)

	_, _ = os.Stderr.WriteString(logMsg)

	if filePath == "" {
		filePath = l.FilePath
	}
	l.WriteLogMsgToFile(filePath, logMsg)
}

func (l *Logger) WriteLogMsgToFile(filePath, loggerStr string) {
	CheckFileDir(filePath)

	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("file open failed, file path: %v, err: %v\n", l.FilePath, err)
	}
	defer file.Close()

	data := []byte(loggerStr)
	_, err = file.Write(data)
	if err != nil {
		fmt.Printf("file write data failed, file path: %v, err: %v\n", l.FilePath, err)
	}
}

func (l *Logger) getLogLevelColor(logLevel int) string {
	logLevelColor, ok := logLevelColor[logLevel]
	if !ok {
		return ColorClose
	}

	return logLevelColor
}

func (l *Logger) getTimePrefix(curTime time.Time) string {
	return curTime.Format("2006-01-02 15:04:05")
}

func (l *Logger) getLogLevelPrefix(logLevel int) string {
	logLevelPrefix, ok := logLevelPrefix[logLevel]
	if !ok {
		return ""
	}
	return logLevelPrefix
}

func (l *Logger) getLogPrefix(logLevel int, curTime time.Time) string {
	return fmt.Sprintf("%s %s", l.getTimePrefix(curTime), l.getLogLevelPrefix(logLevel))
}

func GetCallerInfo() (string, int) {
	_, codeFile, codeLine, _ := runtime.Caller(3)
	return codeFile, codeLine
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}

func GetFileDirPath(path string) string {
	lastIndex := strings.LastIndex(path, "/")
	if lastIndex == -1 {
		return ""
	}

	if lastIndex == 0 {
		return "/"
	}

	return path[0:lastIndex]
}

func CheckFileDir(filePath string) {
	dirPath := GetFileDirPath(filePath)

	if len(dirPath) == 0 || strings.Compare(dirPath, "/") == 0 {
		return
	}

	ok, _ := PathExists(dirPath)
	if !ok {
		if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
			fmt.Printf("create dir failed, path: %v, err: %v", dirPath, err)
		}
	}
}
