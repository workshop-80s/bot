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

	return []entity.Article{}
}
