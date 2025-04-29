package repository

import (
	"github.com/jinzhu/gorm"

	"bot/domain/notify/entity"
	"bot/domain/notify/infrastructure/model"
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
	const modeEnabled = 11

	data := []model.Article{}
	a.storage.
		Select([]string{"id", "title", "sapo"}).
		Where("mode = ?", modeEnabled).
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
		"https://nguoiquansat.vn/gan-100-quy-dau-tu-rot-von-viet-nam-dan-tro-thanh-diem-nong-cua-dong-von-cong-nghe-cao-213283.html",
	)
}
