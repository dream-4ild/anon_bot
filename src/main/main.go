package main

import (
	"context"
	"github.com/go-telegram/bot"
	"os"
	"os/signal"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithDefaultHandler(defaultHandler),
	}

	b, err := bot.New(os.Getenv("ANON_BOT_TOKEN"), opts...)
	if err != nil {
		panic(err)
	}

	b.RegisterHandler(bot.HandlerTypeMessageText, "/q ", bot.MatchTypePrefix, questionHandler)
	b.RegisterHandler(bot.HandlerTypePhotoCaption, "/q", bot.MatchTypePrefix, questionHandler)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/approve", bot.MatchTypePrefix, approveHandler)

	b.Start(ctx)
}
