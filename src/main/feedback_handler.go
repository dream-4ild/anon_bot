package main

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"log"
)

func feedbackHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.Message == nil {
		log.Printf("smth wrong %v\n", update)
		return
	}

	if len(update.Message.Text) < 3 {
		return
	}

	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: moderChatID,
		Text:   "Feedback for YOU\n" + update.Message.Text[3:],
	})

	if err != nil {
		log.Println(err)
		return
	}
}
