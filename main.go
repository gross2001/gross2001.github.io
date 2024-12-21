package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithDefaultHandler(handler),
	}

	b, err := bot.New("5882171886:AAFrDE_VMBK73c-ME12nNaV0jQaKVRbxrPw", opts...)
	if err != nil {
		panic(err)
	}
	//	b.RegisterHandler(bot.HandlerTypeMessageText, "/hello", bot.MatchTypeExact, helloHandler)
	b.Start(ctx)
}

func handler(ctx context.Context, b *bot.Bot, update *models.Update) {
	//	ticker := time.NewTicker(5 * time.Second)
	//	for {
	//		time.Sleep(5 * time.Second)
	//		<-ticker.C
	log.Println(update)
	if update != nil && update.Message != nil {
		fmt.Printf("%+v\n", update.Message)
	}

	if update != nil && update.Message != nil && update.Message.Location != nil {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   fmt.Sprintf("%f %f", update.Message.Location.Latitude, update.Message.Location.Longitude),
		})
		//		}
	}
}
