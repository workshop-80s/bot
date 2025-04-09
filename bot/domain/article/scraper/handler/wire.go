//go:build wireinject
// +build wireinject

package handler

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"

	uc "bot/domain/article/scraper/usecase"
)

func InitScraper(storage *gorm.DB) uc.Scraper {
	panic(wire.Build(
		uc.ProviderScraper,
	))
}
