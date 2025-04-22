package hub

import (
	"bot/domain/article/scraper/entity"
	cafef "bot/domain/article/scraper/infrastructure/external/hub/cafef"
	_ "bot/domain/article/scraper/infrastructure/external/hub/nqs"
)

type (
	Crawler interface {
		CrawlTopPage() []entity.Article
		CrawlDetailPage(url string) entity.Article
	}
)

func NewCrawler(hub entity.ArticleHub) Crawler {
	hubId := hub.Id()
	domain := hub.Domain()

	switch hubId {
	case 1: // cafef
		return cafef.NewCafef(hubId, domain)
		// case 2:
		// 	return nqs.NewCrawler()
	}
	panic("not found crawler")
}
