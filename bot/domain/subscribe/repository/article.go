package repository

import (
	"bot/domain/subscribe/entity"
)

type (
	Article interface {
		Find() []entity.Article
	}
)
