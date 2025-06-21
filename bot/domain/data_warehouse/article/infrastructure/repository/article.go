package repository

import (
	"github.com/jinzhu/gorm"

	"bot/domain/data_warehouse/article/entity"
	"bot/domain/data_warehouse/article/infrastructure/model"
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

func (a Article) FindByOption(
	option map[string]interface{},
) []entity.Article {
	condition := map[string]interface{}{}
	if option["condition"] != nil {
		condition = option["condition"].(map[string]interface{})
	}

	attribute := []string{"id"}
	if option["select"] != nil {
		attribute = option["select"].([]string)
	}

	data := []model.Article{}
	a.storage.
		Select(attribute).
		Where(condition).
		Find(&data)

	result := []entity.Article{}
	for _, t := range data {
		result = append(result, castToArticleEntity(t))
	}

	return result
}

func (a Article) Create(article entity.Article) (int, error) {
	m := model.NewArticle(
		0, // id
		article.Mode(),
		article.Title(),
		article.Sapo(),
		article.Content(),
		article.Image(),
		article.PublishedAt(),
	)

	r := a.storage.Create(&m)
	return m.ID, r.Error
}

func (a Article) Update(
	article entity.Article,
	scopes []string,
) error {
	fields := make([]interface{}, len(scopes))
	for i, scope := range scopes {
		fields[i] = scope
	}

	id := article.ID()
	m := model.NewArticle(
		id,
		article.Mode(),
		article.Title(),
		article.Sapo(),
		article.Content(),
		article.Image(),
		article.PublishedAt(),
	)

	return a.storage.
		Model(&m).
		Select("", fields...). // prevent saving when don't specify scopes
		Where("id = ?", id).
		Updates(m).
		Error
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
		t.PublishedAt,
	)
}
