package controller

import (
	"context"
	"transponder-bot/internal/usecase"
)

var handlers =

type Commands struct {
	users *usecase.Users
	handlers map[string]func(context.Context)
}

func Run(ctx context.Context)  {



}