package hub

import (
	"bot/domain/article/scraper/entity"
	cafef "bot/domain/article/scraper/infrastructure/external/hub/cafef"
	nqs "bot/domain/article/scraper/infrastructure/external/hub/nqs"
)

func CrawlTopPage(agent string) []entity.Article {
	switch agent {
	case "cafef":
		return cafef.CrawlTopPage()
	case "nqs":
		return nqs.CrawlTopPage()
	}
	return []entity.Article{}
}

func CrawlDetail(agent string, url string) entity.Article {
	switch agent {
	case "cafef":
		return cafef.CrawlDetail(url)
	case "nqs":
		return nqs.CrawlDetail(url)
	}
	return entity.Article{}
}
