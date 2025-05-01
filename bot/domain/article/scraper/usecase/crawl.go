package usecase

import (
	"fmt"

	"github.com/google/wire"

	"bot/domain/article/scraper/entity"
	hub "bot/domain/article/scraper/infrastructure/external/hub"

	repository "bot/domain/article/scraper/infrastructure/repository"
	repositoryI "bot/domain/article/scraper/repository"
)

type (
	Scraper struct {
		article    repositoryI.Article
		articleHub repositoryI.ArticleHub
	}
)

var ProviderScraper = wire.NewSet(
	NewScraper,
	repository.NewArticleHub,
	repository.NewArticle,
	wire.Bind(new(repositoryI.ArticleHub), new(repository.ArticleHub)),
	wire.Bind(new(repositoryI.Article), new(repository.Article)),
)

func NewScraper(
	articleHubRepository repositoryI.ArticleHub,
	articleRepository repositoryI.Article,
) Scraper {
	return Scraper{
		article:    articleRepository,
		articleHub: articleHubRepository,
	}
}

func (s Scraper) Crawl() {
	hubs := s.articleHub.Find()

	pages := []entity.Article{}
	for _, h := range hubs {
		c := h.Code()
		pages = append(pages, s.getPageUrl(c)...)
	}

	// details := []entity.Article{}
	for i, p := range pages {
		if i > 0 {
			break
		}
		fmt.Printf("%s\n", p.Url())
		d := s.getPageContent("nqs", p.Url())
		fmt.Println("details:", d.Title())
	}
}

func (s Scraper) getPageUrl(agent string) []entity.Article {
	return hub.CrawlTopPage(agent)
}

func (s Scraper) getPageContent(agent string, url string) entity.Article {
	return hub.CrawlDetail(agent, url)
}
