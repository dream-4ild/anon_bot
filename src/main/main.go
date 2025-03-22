package main

import (
	"context"
	"github.com/go-telegram/bot"
	"log"
	"os"
	"os/signal"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	log.Printf("acsdlkklv")

	opts := []bot.Option{
		bot.WithMiddlewares(checkTargetChat),
		bot.WithDefaultHandler(defaultHandler),
	}

	b, err := bot.New(os.Getenv("ANON_BOT_TOKEN"), opts...)
	if err != nil {
		panic(err)
	}

	b.RegisterHandler(bot.HandlerTypeMessageText, "/q ", bot.MatchTypePrefix, questionHandler)
	b.RegisterHandler(bot.HandlerTypePhotoCaption, "/q", bot.MatchTypePrefix, questionHandler)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/approve", bot.MatchTypePrefix, approveHandler)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/f", bot.MatchTypePrefix, feedbackHandler)

	b.Start(ctx)
}
