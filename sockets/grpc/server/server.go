package main

import (
	"context"
	"encoding/base64"
	"fmt"
	pb "learn_go/sockets/grpc/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

// PORT port
const PORT = "9002"

// SecretService 服务
type SecretService struct{}

// Encrypt 加密
func (s *SecretService) Encrypt(ctx context.Context, r *pb.SecretRequest) (*pb.SecretResponse, error) {
	log.Println("调用了加密服务:", r.Request)
	secret := base64.StdEncoding.EncodeToString([]byte(r.Request))
	return &pb.SecretResponse{Response: r.GetRequest() + "加密成功:" + secret}, nil
}

// Decrypt 解密
func (s *SecretService) Decrypt(ctx context.Context, r *pb.SecretRequest) (*pb.SecretResponse, error) {
	log.Println("调用了解密服务:", r.Request)
	str, err := base64.StdEncoding.DecodeString(r.Request)
	if err != nil {
		log.Fatalf("服务端解密失败:%v")
	}
	return &pb.SecretResponse{Response: r.GetRequest() + "解密成功:" + string(str)}, nil
}

func main() {
	server := grpc.NewServer()
	pb.RegisterSecretServiceServer(server, &SecretService{})

	lis, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		log.Fatalf("net.Listen err:%v", err)
	}
	addr := fmt.Sprintf("%s", lis.Addr())
	log.Printf("Listening on %v", addr)
	log.Fatal(server.Serve(lis))
}
