package main

import (
	"fmt"
	"learn_go/highConcurrency/secmap/common"
)

// PrintInfo Print info of map
func PrintInfo(k, v interface{}) {
	fmt.Printf("index is %v,\t value is %v\n", k, v)
}

func main() {
	secMap := common.NewSynchronizedMap()
	secMap.Put('a', 97)
	secMap.Put('b', 98)
	// value := secMap.Get('a')
	// fmt.Println(value)
	secMap.Each(PrintInfo)
}
