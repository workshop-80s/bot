package repository

import (
	"bot/domain/data_warehouse/article/entity"
)

type (
	Article interface {
		Find() []entity.Article
		FindByOption(map[string]interface{}) []entity.Article
		Create(entity.Article) (int, error)
		Update(entity.Article, []string) error
	}
)
