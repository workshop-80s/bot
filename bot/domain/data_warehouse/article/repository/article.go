package repository

import (
	"bot/domain/data_warehouse/article/entity"
)

type (
	Article interface {
		Find() []entity.Article
		FindByOption(map[string]any) []entity.Article
		Create(entity.Article) (int, error)
		Update(entity.Article, []string) error
	}
)
