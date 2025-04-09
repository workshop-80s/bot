package usecase

import (
	"fmt"

	"github.com/google/wire"

	cafef "bot/domain/article/scraper/infrastructure/external/hub/cafef"

	repository "bot/domain/article/scraper/infrastructure/repository"
	repositoryI "bot/domain/article/scraper/repository"
)

type (
	Scraper struct {
		repository repositoryI.Article
	}
)

var ProviderScraper = wire.NewSet(
	NewScraper,
	repository.NewArticle,
	wire.Bind(new(repositoryI.Article), new(repository.Article)),
)

func NewScraper(r repositoryI.Article) Scraper {
	return Scraper{
		repository: r,
	}
}

func (s Scraper) getListFromHub() {
	// const domain = "https://nguoiquansat.vn"
	const domain = "https://cafef.vn"

	p := cafef.CrawlTopPage(domain)
	for _, e := range p {
		// Code to be executed for each element
		fmt.Println(e.Title())
		fmt.Println(e.Url())
	}
}

func (s Scraper) Scrape() {
	// fetch article_original
	// scrape article
	// save database
	s.getListFromHub()

	fmt.Println("scraper")

}
