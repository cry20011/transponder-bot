package controller

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"time"
	"transponder-bot/internal/usecase"
)

type BotCommands struct {
	users    *usecase.Users
	bot      *tgbotapi.BotAPI
	handlers map[string]func(context.Context, tgbotapi.Update)
	timeout  time.Duration
}

func NewBotCommands(u *usecase.Users, token string, timeout time.Duration) (*BotCommands, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	return &BotCommands{
		users:    u,
		bot:      bot,
		timeout:  timeout,
		handlers: make(map[string]func(context.Context, tgbotapi.Update)),
	}, nil
}

func (c *BotCommands) Run(ctx context.Context) error {
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	c.addCommand("/start", func(ctx context.Context, update tgbotapi.Update) {
		err := c.users.AddUser(ctx, update.Message.From.UserName)
		if err != nil {
			return
		}

		msg := "Welcome to transponder bot."

		_, err = c.bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, msg))
		if err != nil {
			return
		}
	})

	return c.listenUpdates(ctx)
}

func (c *BotCommands) addCommand(command string, handler func(ctx context.Context, update tgbotapi.Update)) {
	c.handlers[command] = func(ctx context.Context, update tgbotapi.Update) {
		ctxWithTimeout, cancel := context.WithTimeout(ctx, c.timeout)
		defer cancel()

		handler(ctxWithTimeout, update)
	}
}

func (c *BotCommands) listenUpdates(ctx context.Context) error {
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	updates, err := c.bot.GetUpdatesChan(updateConfig)
	if err != nil {
		return err
	}

	go func() {
		for {
			select {
			case update := <-updates:
				if update.Message.IsCommand() {
					if handler, ok := c.handlers[update.Message.Text]; ok {
						go handler(ctx, update)
					} else {
						log.Printf("unknown command")
					}
				}
			case <-ctx.Done():
				return
			}
		}
	}()

	return nil
}
