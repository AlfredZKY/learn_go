package main

import (
	"fmt"
	"os" 
)

func main(){
	var LotusMinerPath string
	LotusMinerPath = os.Getenv("LOTUS_MINER_PATH")
	fmt.Println(LotusMinerPath)
}