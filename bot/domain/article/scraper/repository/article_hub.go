package repository

import (
	"bot/domain/article/scraper/entity"
)

type (
	ArticleHub interface {
		Find() []entity.ArticleHub
	}
)

// condition := map[string]interface{}{}
// 	if option["condition"] != nil {
// 		condition = option["condition"].(map[string]interface{})
// 	}
