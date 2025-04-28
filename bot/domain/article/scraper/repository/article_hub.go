package repository

import (
	"bot/domain/article/scraper/entity"
)

type (
	ArticleHub interface {
		Find() []entity.ArticleHub
	}
)
