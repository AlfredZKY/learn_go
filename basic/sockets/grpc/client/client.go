package main

import (
	"context"
	pb "learn_go/sockets/grpc/proto"
	"log"
	"strings"

	"google.golang.org/grpc"
)

const PORT = "9002"

var str = "this is gRPC"

func main() {
	conn, err := grpc.Dial(":"+PORT, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc.Dial err: %v", err)
	}
	defer conn.Close()

	client := pb.NewSecretServiceClient(conn)

	resp0, err := client.Encrypt(context.Background(), &pb.SecretRequest{
		Request: str,
	})

	if err != nil {
		log.Fatalf("client.Encrypt err: %v", err)
	}

	secretStr := strings.Split(resp0.GetResponse(), ":")[1]
	log.Printf("请求加密服务: %s", secretStr)

	resp1, err := client.Decrypt(context.Background(), &pb.SecretRequest{
		Request: secretStr,
	})

	if err != nil {
		log.Fatalf("client.Decrypt err: %v", err)
	}

	resStr := strings.Split(resp1.GetResponse(), ":")[1]
	log.Printf("请求解密服务: %s", resStr)
}
