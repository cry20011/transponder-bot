package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"transponder-bot/config"
	"transponder-bot/internal/bot"
)

func main() {
	conf, err := config.New()
	if err != nil {
		log.Fatalf("failed to get configuration: %v", err)
	}

	ctx, stop := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)
	defer stop()

	if err := bot.Run(ctx, conf); err != nil {
		log.Fatalf("run bot failed: %v", err)
	}

	os.Exit(0)
}
