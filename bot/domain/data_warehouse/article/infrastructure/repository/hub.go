package repository

import (
	"github.com/jinzhu/gorm"

	"bot/domain/data_warehouse/article/entity"
	"bot/domain/data_warehouse/article/infrastructure/model"
)

type (
	Hub struct {
		storage *gorm.DB
	}
)

func NewHub(storage *gorm.DB) Hub {
	return Hub{
		storage: storage,
	}
}

func (a Hub) Find() []entity.Hub {
	const modeEnabled = 11

	data := []model.Hub{}
	a.storage.
		Select([]string{"id", "domain"}).
		Where("mode = ?", modeEnabled).
		Find(&data)

	result := []entity.Hub{}
	for _, t := range data {
		result = append(result, castToHubEntity(t))
	}

	return result
}

func castToHubEntity(
	t model.Hub,
) entity.Hub {
	return entity.NewHub(
		t.ID,
		t.Domain,
	)
}
