package usecase

import (
	"fmt"

	"github.com/google/wire"

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
	// fetch article_original
	// scrape article
	// save database
	// agent := "nqs"
	// s.getListFromHub(agent)

	hubs := s.articleHub.Find()
	for _, h := range hubs {
		c := h.Code()
		s.getListFromHub(c)
	}

	s.article.Find()
}

func (s Scraper) getListFromHub(agent string) {
	p := hub.CrawlTopPage(agent)
	for _, e := range p {
		fmt.Printf("%s: %s\n", agent, e.Url())
	}
}
