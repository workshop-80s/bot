package repository

import (
	"github.com/jinzhu/gorm"

	"bot/domain/data_warehouse/article/entity"
	"bot/domain/data_warehouse/article/infrastructure/model"
)

type (
	ArticleHub struct {
		storage *gorm.DB
	}
)

func NewArticleHub(storage *gorm.DB) ArticleHub {
	return ArticleHub{
		storage: storage,
	}
}

func (a ArticleHub) Find() []entity.ArticleHub {
	const modeEnabled = 11

	data := []model.ArticleHub{}
	a.storage.
		Select([]string{"id", "code", "domain"}).
		Where("mode = ?", modeEnabled).
		Find(&data)

	result := []entity.ArticleHub{}
	for _, t := range data {
		result = append(result, castToArticleHubEntity(t))
	}

	return result
}

func castToArticleHubEntity(
	t model.ArticleHub,
) entity.ArticleHub {
	return entity.NewArticleHub(
		t.ID,
		t.Code,
		t.Domain,
	)
}
