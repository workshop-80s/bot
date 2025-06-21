package hub

import (
	"bot/domain/data_warehouse/article/entity"
	"bot/domain/data_warehouse/article/infrastructure/external/hub/cafef"
	// "bot/domain/data_warehouse/article/infrastructure/external/hub/nqs"
)

type (
	Crawler interface {
		CrawlTopPage() []entity.Link
		// CrawlDetailPage(url string) entity.Article
	}
)

func NewCrawler(hub entity.Hub) Crawler {
	hubId := hub.ID()
	domain := hub.Domain()

	switch hubId {
	case 1: // cafef
		return cafef.NewHub(hubId, domain)
		// case 2:
		// 	return nqs.NewHub(hubId, domain)
	}

	panic("not found crawler")
}
