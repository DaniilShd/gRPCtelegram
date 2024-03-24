package main

import (
	"context"
	"fmt"
	"log"
	"net"

	desc "github.com/DaniilShd/gRPCtelegram/pkg/user_v1/api/user_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const grpcPort = 50051

type server struct {
	desc.UnimplementedUserV1Server
}

func (s *server) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	log.Printf("User id: %d", req.GetId())

	return &desc.GetResponse{
		Info: &desc.UserInfo{
			Id:      req.GetId(),
			Name:    "Daniil",
			IsHuman: true,
		},
	}, nil
}

func main() {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("Faild %s", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterUserV1Server(s, &server{})

	log.Printf("server listening")

	if err = s.Serve(listen); err != nil {
		log.Fatal("Error!")
	}
}
