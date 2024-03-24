package main

import (
	"context"
	"fmt"
	"log"
	"time"

	desc "github.com/DaniilShd/gRPCtelegram/pkg/user_v1/api/user_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Faild1 %s", err)
	}
	defer conn.Close()

	c := desc.NewSendMessageToTelegramClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	message := desc.TelegramMessage{
		Id:      23,
		ChantID: int64(379017783),
		Text:    "hello!!!!!!",
	}

	var r *desc.SendResponse

	for i := 0; i < 10; i++ {
		r, err = c.Send(ctx, &desc.SendRequest{
			MessageInfo: &message,
		})
		if err != nil {
			log.Fatalf("Faild2 %s", err)
		}
		fmt.Println(r)
	}
}
