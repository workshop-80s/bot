package usecase

import (
	"github.com/google/wire"

	"bot/domain/data_warehouse/article/entity"
	"bot/domain/data_warehouse/article/infrastructure/external/hub"
	repository "bot/domain/data_warehouse/article/infrastructure/repository"
	repositoryI "bot/domain/data_warehouse/article/repository"
)

type (
	LinkScraper struct {
		hub  repositoryI.Hub
		link repositoryI.Link
	}
)

var LinkScraperProvider = wire.NewSet(
	NewLinkScraper,
	repository.NewHub,
	repository.NewLink,
	wire.Bind(new(repositoryI.Hub), new(repository.Hub)),
	wire.Bind(new(repositoryI.Link), new(repository.Link)),
)

func NewLinkScraper(
	hubRepository repositoryI.Hub,
	linkRepository repositoryI.Link,
) LinkScraper {
	return LinkScraper{
		hub:  hubRepository,
		link: linkRepository,
	}
}

func (s LinkScraper) Crawl() {
	hubs := s.hub.Find()

	links := []entity.Link{}

	for _, h := range hubs {
		r := s.crawl(h)
		links = append(links, r...)
	}

	// save to database
	for _, a := range links {
		s.link.Create(a)
	}
}

func (s LinkScraper) crawl(h entity.Hub) []entity.Link {
	crawler := hub.NewCrawler(h)
	links := crawler.CrawlTopPage()

	newLinks := []entity.Link{}
	for _, t := range links {
		url := t.Url()
		if s.isCrawled(url) {
			continue
		}
		newLinks = append(newLinks, t)
	}

	return newLinks
}

func (s LinkScraper) isCrawled(url string) bool {
	attribute := []string{"id", "title"}
	condition := map[string]interface{}{
		"url": url,
	}

	option := map[string]interface{}{
		"condition": condition,
		"attribute": attribute,
	}

	links := s.link.FindByOption(option)
	return len(links) > 0
}
