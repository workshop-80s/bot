package article

import (
	"bot/domain/article/crawler/entity"
)

type (
	Crawler interface {
		CrawlTopPage(
			link string,
		) []entity.Page

		CrawlDetail(
			link string,
		) entity.Page
	}
)
