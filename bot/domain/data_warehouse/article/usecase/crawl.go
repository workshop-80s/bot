package usecase

import (
	"github.com/google/wire"

	"bot/domain/data_warehouse/article/entity"
	"bot/domain/data_warehouse/article/infrastructure/external/hub"
	repository "bot/domain/data_warehouse/article/infrastructure/repository"
	repositoryI "bot/domain/data_warehouse/article/repository"
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

	articles := []entity.Article{}
	for _, h := range hubs {
		r := s.crawl(h)
		articles = append(articles, r...)
	}

	// save to database
	// for _, a := range articles {
	// 	s.article.Create(a)
	// }
}

func (s Scraper) crawl(h entity.ArticleHub) []entity.Article {
	crawler := hub.NewCrawler(h)
	crawler.CrawlTopPage()
	// details := crawler.CrawlTopPage()

	articles := []entity.Article{}

	// i := 0
	// for _, d := range details {
	// 	i++
	// 	if i > 1 {
	// 		break
	// 	}

	// 	url := d.Url()
	// 	if s.isCrawled(url) {
	// 		fmt.Printf("already crawled: %s\n", url)
	// 		continue
	// 	}

	// 	article := crawler.CrawlDetailPage(url)
	// 	articles = append(articles, article)
	// }

	return articles

}

func (s Scraper) isCrawled(url string) bool {
	attribute := []string{"id", "title"}
	condition := map[string]interface{}{
		"origin": url,
	}

	option := map[string]interface{}{
		"condition": condition,
		"attribute": attribute,
	}

	articles := s.article.FindByOption(option)
	return len(articles) > 0
}
