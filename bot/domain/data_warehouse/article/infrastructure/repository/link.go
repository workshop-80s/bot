package repository

import (
	"github.com/jinzhu/gorm"

	"bot/domain/data_warehouse/article/entity"
	"bot/domain/data_warehouse/article/infrastructure/model"
)

type (
	Link struct {
		storage *gorm.DB
	}
)

func NewLink(storage *gorm.DB) Link {
	return Link{
		storage: storage,
	}
}

func (a Link) Find() []entity.Link {
	const modeUnCrawl = 10

	data := []model.Link{}
	a.storage.
		Select([]string{"id", "domain"}).
		Where("mode = ?", modeUnCrawl).
		Find(&data)

	result := []entity.Link{}
	for _, t := range data {
		result = append(result, castToLinkEntity(t))
	}

	return result
}

func (a Link) FindByOption(
	option map[string]any,
) []entity.Link {
	condition := map[string]interface{}{}
	if option["condition"] != nil {
		condition = option["condition"].(map[string]interface{})
	}

	attribute := []string{"id"}
	if option["select"] != nil {
		attribute = option["select"].([]string)
	}

	data := []model.Link{}
	a.storage.
		Select(attribute).
		Where(condition).
		Find(&data)

	result := []entity.Link{}
	for _, t := range data {
		result = append(result, castToLinkEntity(t))
	}

	return result
}

func (a Link) Create(link entity.Link) (int, error) {
	m := model.NewLink(
		0, // id
		10,
		link.HubID(),
		link.Url(),
		0,
		nil,
	)

	r := a.storage.Create(&m)
	return m.ID, r.Error
}

func castToLinkEntity(
	t model.Link,
) entity.Link {
	return entity.NewLink(
		t.ID,
		t.HubID,
		t.Url,
	)
}
