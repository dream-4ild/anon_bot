package main

import (
	"context"
	"fmt"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"log"
)

const (
	approveTemplate = "/approve "
	approveLen      = len(approveTemplate)
)

func requestApprove(ctx context.Context, b *bot.Bot, message *models.Message) error {
	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:    moderChatID,
		Text:      fmt.Sprintf("`%v%v`", approveTemplate, message.ID),
		ParseMode: models.ParseModeMarkdown,
	})
	return err
}

func checkTargetChat(next bot.HandlerFunc) bot.HandlerFunc {
	return func(ctx context.Context, b *bot.Bot, update *models.Update) {
		log.Printf("Check target chat: %v; target: %v", update.Message.Chat.ID, targetChatID)
		if update.Message != nil && update.Message.Chat.ID == targetChatID {
			return
		}
		next(ctx, b, update)
	}
}
