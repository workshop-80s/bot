//go:build wireinject
// +build wireinject

package handler

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"

	uc "bot/domain/data_warehouse/article/usecase"
)

func InitScraper(storage *gorm.DB) uc.LinkScraper {
	panic(wire.Build(
		uc.LinkScraperProvider,
	))
}
