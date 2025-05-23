// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package handler

import (
	"bot/domain/notify/infrastructure/repository"
	"bot/domain/notify/usecase"
	"github.com/jinzhu/gorm"
)

// Injectors from wire.go:

func InitialSlack(storage *gorm.DB) usecase.Slack {
	article := repository.NewArticle(storage)
	usecaseSlack := usecase.NewSlack(article)
	return usecaseSlack
}
