package repository

import (
	"fmt"

	"github.com/jinzhu/gorm"

	"bot/domain/subscribe/entity"
	"bot/domain/subscribe/infrastructure/model"
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
	fmt.Println("domain/subscribe/infrastructure/article")

	data := []model.Article{}
	a.storage.
		Select([]string{"id", "title", "url"}).
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
		t.Title,
		t.Url,
	)
}
