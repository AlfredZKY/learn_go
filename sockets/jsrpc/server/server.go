package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// Params export
type Params struct {
	Width, Height int
}

// Rect 矩形
type Rect struct{}

// Area 求积
func (r *Rect) Area(p Params, ret *int) error {
	*ret = p.Height * p.Width
	return nil
}

// Perimeter 求周长
func (r *Rect) Perimeter(p Params, ret *int) error {
	*ret = (p.Width + p.Height) * 2
	return nil
}
func chkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	rect := new(Rect)

	rpc.Register(rect)

	tcpAddr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:9091")
	chkError(err)
	tcpListen, err2 := net.ListenTCP("tcp", tcpAddr)
	chkError(err2)
	for {
		conn, err3 := tcpListen.Accept()
		if err3 != nil {
			continue
		}
		go jsonrpc.ServeConn(conn)
	}
}
