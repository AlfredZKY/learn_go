package main

import (
	"fmt"
	"learn_go/middleware/jiankong"
	"learn_go/middleware/readconfig"
)

// CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go
func main() {
	// p,err := readconfig.ReadConf("/home/zky/go/src/middleware/config.toml")
	// if err != nil{
	// 	fmt.Printf("%v",err)
	// }
	// fmt.Printf("Person %s\n",p.MonitorUnit.LocalIP)
	//LocalIP := p.MonitorUnit.LocalIP
	jiankong.Send(readconfig.LocalIP, "this is a messgae for testing")
	allocated, ubytes := 123, 456
	newValue := fmt.Sprintf("too much data in sector: %d > %d", allocated, ubytes)
	fmt.Println(newValue)
}
