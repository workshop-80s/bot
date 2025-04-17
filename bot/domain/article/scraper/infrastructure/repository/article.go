package repository

import (
	"github.com/jinzhu/gorm"

	"bot/domain/article/scraper/entity"
	"bot/domain/article/scraper/infrastructure/model"
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
		result = append(result, castToArticleEntity(t))
	}

	return result
}

func castToArticleEntity(
	t model.Article,
) entity.Article {
	return entity.NewArticle(
		t.ID,
		t.Mode,
		t.Title,
		t.Sapo,
		t.Content,
		t.Image,
		t.Origin,
		t.ArticleHubID,
	)
}
