package main

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"log"
)

const (
	defaultMessage = "Supported commands are:\n" +
		"..."
)

func defaultHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.Message != nil {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "Unknown command :(\n" + defaultMessage,
		})
	} else {
		log.Printf("Unknown update: %v\n", update)
	}
}
