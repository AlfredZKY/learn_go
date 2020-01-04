package main

import (
	//"runtime"
	"log"
	"math/rand"
	"runtime/pprof"
	"os"
	"time"
)

const (
	col = 10000
	row = 10000
)

func fillMatrix(m *[row][col]int) {
	s := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			m[i][j] = s.Intn(1000000)
		}
	}
}

func calculate(m *[row][col]int) {
	for i := 0; i < row; i++ {
		tmp := 0
		for j := 0; j < col; j++ {
			tmp += m[i][j]
		}
	}
}

func main() {
	// 创建输出文件
	f, err := os.Create("cpu.info")
	if err != nil {
		log.Fatal("could not create CPU profile: ", err)
	}

	// 获取系统信息
	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal("could not start CPU profile: ",err)
	}

	defer pprof.StopCPUProfile()
	
	// 主要逻辑区代码
	x := [row][col]int{}
	fillMatrix(&x)
	calculate(&x)

	f1 ,err := os.Create("mem.prof")
	if err != nil{
		log.Fatal("could not create memory profile: ",err)
	}

	// runtime.GC()

	if err := pprof.WriteHeapProfile(f1);err!=nil{
		log.Fatal("could not write mempory profile: ",err)
	}
	f1.Close()
	
	f2 ,err := os.Create("goroutine.prof")
	if err != nil{
		log.Fatal("could not create goroutine profile: ",err)
	}
	if gProf := pprof.Lookup("goroutine");gProf == nil{
		log.Fatal("could not write goroutine profile: ")
	}else{
		gProf.WriteTo(f2,0)
	}
	f2.Close()
}
