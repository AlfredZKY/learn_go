package main

import (
	"fmt"
	"context"
	hw "learn_go/learnrpcx/genprobuf/hw"
	server "github.com/smallnest/rpcx/server"
)

// GreetImpl
type GreetImpl struct{}

func main(){
	s := server.NewServer()
	s.RegisterName("Greet",new(GreetImpl),"")
	err := s.Serve("tcp",":8972")
	if err != nil {
		panic(err)
	}
}	

func (s * GreetImpl)SayHello(ctx context.Context, args *hw.HelloRequest, reply *hw.HelloReply) (err error) {
	*reply = hw.HelloReply{
		Message:fmt.Sprintf("hello %s!",args.Name),
	}
	return nil
}

