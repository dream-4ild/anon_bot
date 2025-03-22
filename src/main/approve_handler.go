package main

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"log"
	"strconv"
)

func approveHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.Message == nil {
		return
	}

	if update.Message.Chat.ID != moderChatID {
		return
	}

	messId, _ := strconv.Atoi(update.Message.Text[approveLen:])

	_, err := b.ForwardMessage(ctx, &bot.ForwardMessageParams{
		ChatID:          targetChatID,
		MessageID:       messId,
		MessageThreadID: targetThreadID,
		FromChatID:      moderChatID,
	})
	if err != nil {
		log.Printf("failed to forward message: %v", err)
		return
	}
}
