package bot

import (
	"context"
	"time"
	"transponder-bot/config"
	"transponder-bot/internal/controller"
	"transponder-bot/internal/usecase"
	"transponder-bot/internal/usecase/repo"
	"transponder-bot/pkg/postgres"
)

func Run(ctx context.Context, config *config.Config) error {
	p, err := postgres.New(postgres.Options{
		Host:     "localhost",
		Port:     "5432",
		User:     "admin",
		Password: "admin",
		DbName:   "transponder_bot_db",
		SslMode:  "disable",
	})
	if err != nil {
		return err
	}

	usersRepo := repo.NewUsers(p)

	usersUsecase := usecase.NewUsers(usersRepo)

	botCommands, err := controller.NewBotCommands(usersUsecase, config.Token, time.Minute)
	if err != nil {
		return err
	}

	err = botCommands.Run(ctx)
	if err != nil {
		return err
	}

	<-ctx.Done()

	return nil
}
