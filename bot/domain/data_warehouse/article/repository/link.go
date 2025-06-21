package repository

import (
	"bot/domain/data_warehouse/article/entity"
)

type (
	Link interface {
		Find() []entity.Link
		FindByOption(map[string]any) []entity.Link
		Create(entity.Link) (int, error)
		// Update(entity.Link, []string) error
	}
)
