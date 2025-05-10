//go:build wireinject
// +build wireinject

package handler

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"

	uc "bot/domain/notify/usecase"
)

func InitialSlack(
	storage *gorm.DB,
	botToken string,
	channels map[string]string,
) uc.Slack {
	panic(wire.Build(
		uc.ProviderSlack,
	))
}
