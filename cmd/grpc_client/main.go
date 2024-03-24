package main

import (
	"context"
	"fmt"
	"log"
	"time"

	desc "github.com/DaniilShd/gRPCExample/pkg/user_v1/api/user_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	address = "localhost:50051"
	userID  = 12
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Faild %s", err)
	}
	defer conn.Close()

	c := desc.NewUserV1Client(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Get(ctx, &desc.GetRequest{Id: 16})
	if err != nil {
		log.Fatalf("Faild %s", err)
	}

	fmt.Println(r)
}
