package main

import (
	"fmt"
	"log"
	"net/rpc/jsonrpc"
)

// Params export
type Params struct {
	Width, Height int
}

func main() {
	rpc, err := jsonrpc.Dial("tcp", "127.0.0.1:9091")
	if err != nil {
		log.Fatal(err)
	}

	ret := 0

	err2 := rpc.Call("Rect.Area", Params{50, 100}, &ret)
	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println(ret)

	err3 := rpc.Call("Rect.Perimeter", Params{50, 100}, &ret)
	if err2 != nil {
		log.Fatal(err3)
	}
	fmt.Println(ret)
}
