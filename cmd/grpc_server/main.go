package main

import (
	"context"
	"fmt"
	"log"
	"net"

	desc "github.com/DaniilShd/gRPCtelegram/pkg/user_v1/api/user_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const grpcPort = 50051

type server struct {
	desc.UnimplementedSendMessageToTelegramServer
	myBot *tgbotapi.BotAPI
}

func (s *server) Send(ctx context.Context, req *desc.SendRequest) (*desc.SendResponse, error) {

	msg := tgbotapi.NewMessage(req.MessageInfo.ChantID, req.MessageInfo.Text)
	msg.ParseMode = "html"
	s.myBot.Send(msg)

	return &desc.SendResponse{
		Check: true,
	}, nil
}

func main() {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("Faild %s", err)
	}

	ch := make(chan *tgbotapi.BotAPI)

	go StartTelegramBot(ch)

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterSendMessageToTelegramServer(s, &server{
		myBot: <-ch,
	})

	log.Printf("server listening")

	if err = s.Serve(listen); err != nil {
		log.Fatal("Error!")
	}
}

func StartTelegramBot(ch chan *tgbotapi.BotAPI) {
	bot, err := tgbotapi.NewBotAPI("6996608995:AAFKAsW2XjLEElrAgaeC8qPAzZOTocmlenA")
	if err != nil {
		log.Panic(err)
	}

	ch <- bot

	bot.Debug = false

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID

			fmt.Println(update.Message.Chat.ID)

			bot.Send(msg)
		}
	}
}
