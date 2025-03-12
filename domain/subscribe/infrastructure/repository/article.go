package repository

import (
	"fmt"

	"github.com/jinzhu/gorm"

	"bot/domain/subscribe/infrastructure/model"
)

type (
	article struct {
		storage *gorm.DB
	}
)

func NewArticle(storage *gorm.DB) article {
	return article{
		storage: storage,
	}
}

func (article) TableName() string {
	return "profiles"
}

func (a article) Find() error {
	fmt.Println("domain/subscribe/infrastructure/article")

	data := []model.Article{}
	a.storage.
		Select([]string{"id", "title", "url"}).
		Find(&data)
	fmt.Printf("%+v", data)
	return nil
}

func (a article) ApplyTransaction(storage *gorm.DB) {
	a.storage = storage
}
