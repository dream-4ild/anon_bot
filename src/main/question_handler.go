package main

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"log"
)

func handleText(ctx context.Context, b *bot.Bot, update *models.Update) error {

	if update.Message.Text == "" {
		return nil
	}

	message, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: moderChatID,
		Text:   update.Message.Text[2:],
	})

	if err != nil {
		log.Println(err)
		return err
	}

	if err := requestApprove(ctx, b, message); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func handlerPhotos(ctx context.Context, b *bot.Bot, update *models.Update) error {
	if len(update.Message.Photo) > 0 {
		caption := update.Message.Caption[2:]
		if len(caption) > 0 {
			caption = caption[1:]
		}

		message, err := b.SendPhoto(ctx, &bot.SendPhotoParams{
			ChatID:  moderChatID,
			Photo:   &models.InputFileString{Data: update.Message.Photo[len(update.Message.Photo)-1].FileID},
			Caption: caption,
		})
		if err != nil {
			log.Printf("Ошибка отправки фото: %v", err)
			return err
		}
		if err := requestApprove(ctx, b, message); err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}

func handleDocument(ctx context.Context, b *bot.Bot, update *models.Update) error {
	if update.Message.Document != nil {
		caption := update.Message.Caption[2:]
		if len(caption) > 0 {
			caption = caption[1:]
		}

		message, err := b.SendDocument(ctx, &bot.SendDocumentParams{
			ChatID:   moderChatID,
			Document: &models.InputFileString{Data: update.Message.Document.FileID},
			Caption:  caption,
		})
		if err != nil {
			log.Printf("Ошибка отправки документа: %v", err)
			return err
		}
		if err := requestApprove(ctx, b, message); err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}

func questionHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.Message == nil {
		log.Printf("smth wrong %v\n", update)
		return
	}

	err := handleText(ctx, b, update)
	if err != nil {
		return
	}

	err = handlerPhotos(ctx, b, update)
	if err != nil {
		return
	}

	err = handleDocument(ctx, b, update)
	if err != nil {
		return
	}

}
