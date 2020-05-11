package test1

import (
	"errors"
	"os"
	"runtime/pprof"
)

var (
	profileName = "cpuprofile.out"
)

// ShowCPUInfo 显示cpu的信息
func ShowCPUInfo() {
	// f, err := common.CreateFile("", profileName)
	// if err != nil {
	// 	fmt.Printf("CPU profile creation error:%v\n", err)
	// 	return
	// }
	// defer f.Close()

	// if err := StartCPUProfile(f); err != nil {
	// 	fmt.Printf("CPU profile start error:%v\n", err)
	// 	return
	// }

	// if err = common.Execute(op.StartCPUProfile()){}

}

// StartCPUProfile 开始显示cpu信息到文件
func StartCPUProfile(f *os.File) error {
	if f == nil {
		return errors.New("nil file")
	}
	return pprof.StartCPUProfile(f)
}

// StopCPUProfile 停止显示cpu信息
func StopCPUProfile(f *os.File) {
	pprof.StopCPUProfile()
}
