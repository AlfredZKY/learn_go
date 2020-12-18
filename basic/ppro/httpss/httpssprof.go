package main

import (
	"math/rand"
	_ "net/http/pprof"
	"os"
	"runtime/pprof"
	"time"
)

func PProfCPUApplication() {
	f, _ := os.Create("./cpu.prof")
	pprof.StartCPUProfile(f)

	for i := 1; i < 3000; i++ {
		time.Sleep(3 * time.Millisecond)
		rand.Intn(100)
	}
	pprof.StopCPUProfile()
	f.Close()
}

func main() {
	// err := http.ListenAndServe(":9000", nil)
	// if err != nil {
	// 	panic(err)
	// }
	PProfCPUApplication()
}
