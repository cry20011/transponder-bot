package bot

import (
	"context"
	"encoding/json"
	"log"
	"transponder-bot/config"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func Run(ctx context.Context, config *config.Config) error {
	bot, err := tgbotapi.NewBotAPI(config.Token)
	if err != nil {
		return err
	}

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	updates, err := bot.GetUpdatesChan(updateConfig)

	for {
		select {
		case update := <-updates:
			if update.Message == nil {
				continue
			}

			updateJSON, err := json.Marshal(update)
			if err != nil {
				log.Printf("bad update: %v", err)

				continue
			}

			log.Printf("[%s] %s", update.Message.From.UserName, string(updateJSON))

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			_, err = bot.Send(msg)
			if err != nil {
				log.Printf("failed to send to %s: %v", update.Message.From.UserName, err)

				continue
			}
		case <-ctx.Done():
			return nil
		}
	}
}
