package repository

import (
	"bot/domain/article/scraper/entity"
)

type (
	Article interface {
		Find() []entity.Article
	}
)
