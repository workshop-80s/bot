package repository

import (
	"github.com/jinzhu/gorm"

	"bot/domain/article/scraper/entity"
	"bot/domain/article/scraper/infrastructure/model"
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
	data := []model.ArticleHub{}
	a.storage.
		Select([]string{"id", "code"}).
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
	)
}
