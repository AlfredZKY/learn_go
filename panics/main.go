package main

import (
	_ "log"
)

// panic 结构体位置 src/runtime/runtime2.go _panic 结构体
// gopanic 函数的位置 src/runtime/panic.go
func main() {
	// defer func(){
	// 	if err :=recover();err!=nil{
	// 		log.Printf("recover: %v",err)
	// 	}
	// }()
	panic("EDDYCJY.")
}
