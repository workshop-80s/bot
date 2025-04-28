package repository

import (
	"bot/domain/notify/entity"
)

type (
	Article interface {
		Find() []entity.Article
	}
)
