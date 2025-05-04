package repository

import (
	"bot/domain/data_warehouse/article/entity"
)

type (
	ArticleHub interface {
		Find() []entity.ArticleHub
	}
)
