package main

import (
	"fmt"
	// "learn_go/sockets/websockets/server"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func echo(w http.ResponseWriter, r *http.Request) {
	c, _ := upgrader.Upgrade(w, r, nil)
	defer c.Close()
	// for {
	// 	mt, _, _ := c.ReadMessage()
	// 	fmt.Println("收到：")
	// 	c.WriteMessage(mt, []byte("hello"))
	// }
	mt, _, _ := c.ReadMessage()
	fmt.Println("收到：")
	c.WriteMessage(mt, []byte("hello"))
}

func main() {
	http.HandleFunc("/echo", echo)
	log.Fatal(http.ListenAndServe(":9091", nil))
	// ws := server.NewWsServer()
	// ws.Start()
}
