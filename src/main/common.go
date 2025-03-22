package main

import (
	"context"
	"fmt"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
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
