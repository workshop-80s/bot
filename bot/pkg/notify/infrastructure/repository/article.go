package repository

import (
	"github.com/jinzhu/gorm"

	"bot/pkg/notify/entity"
	"bot/pkg/notify/infrastructure/model"
)

type (
	Article struct {
		storage *gorm.DB
	}
)

func NewArticle(storage *gorm.DB) Article {
	return Article{
		storage: storage,
	}
}

func (a Article) Find() []entity.Article {
	data := []model.Article{}
	a.storage.
		Select([]string{"id", "title", "sapo"}).
		Find(&data)

	result := []entity.Article{}
	for _, t := range data {
		result = append(result, castToArticle(t))
	}

	return result
}

func castToArticle(
	t model.Article,
) entity.Article {
	return entity.NewArticle(
		t.ID,
		t.Title,
		t.Sapo,
		"",
	)
}
