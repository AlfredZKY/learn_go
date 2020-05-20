package main

import (
	"context"
	"fmt"
	hw "learn_go/learnrpcx/genprobuf/hw"
)

func main() {
	xclinet := hw.NewXClientForGreet("127.0.0.1:8972")
	client := hw.NewGreetClient(xclinet)

	args := &hw.HelloRequest{
		Name: "1",
	}

	reply, err := client.SayHello(context.Background(), args)
	if err != nil {
		panic(err)
	}
	fmt.Println("reply:", reply.Message)
}
