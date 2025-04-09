package repository

import (
	"bot/pkg/notify/entity"
)

type (
	Article interface {
		Find() []entity.Article
	}
)
