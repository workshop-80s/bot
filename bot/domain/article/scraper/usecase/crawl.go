package usecase

import (
	"fmt"
	"sync"
	"time"

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
	startTime := time.Now()
	fmt.Println("Crawl START")
	s.getList()
	fmt.Println("Crawl END")
	fmt.Println("Crawl Duration: ", time.Since(startTime))
}

// https://www.codingexplorations.com/blog/interview-series-when-to-use-buffered-and-unbuffered-channels
// https://medium.com/@josueparra2892/golang-concurrency-worker-pool-fa62ffe6e438

func getPageLink(
	wg *sync.WaitGroup,
	agents <-chan string,
	details chan<- string,
) {

	agent := <-agents

	fmt.Println("  getPageLink:", agent)
	urls := hub.CrawlTopPage(agent)
	for _, u := range urls {
		fmt.Println("    getPageLink:", u.Title())
		details <- u.Url()
	}
}

func (s Scraper) getList() {
	const numWorkers = 20

	var wg sync.WaitGroup
	// var wgDetail sync.WaitGroup

	details := make(chan string, 20)
	agents := make(chan string, 5)
	// articles := make(chan entity.Article, 50)
	// register workers for top page

	fmt.Println("fetch DB")
	hubs := s.articleHub.Find()
	for _, h := range hubs {
		fmt.Println("fetch DB", h.Code())
		agents <- h.Code()
	}

	wg.Add(numWorkers)

	go func() {
		fmt.Println("Wait BEGIN")

		wg.Wait()
		// close(details)
		close(agents)
		// close(articles)

		fmt.Println("Wait END")
	}()

	// go getPageLink(&wg, agents, details)

	// register workers for detail page
	// for i := 1; i <= numWorkers; i++ {
	// 	wg.Add(1)
	// 	go getPageContent(&wg, details, articles)
	// }

	fmt.Println("getList out BEGIN")
	fmt.Println()
	for v := range details {
		fmt.Println("  detail:", v)
	}
	r := []entity.Article{}
	// for v := range articles {
	// 	fmt.Println("  detail:", v.Title())
	// 	r = append(r, v)
	// }

	fmt.Println("getList out END", len(r))
}

func getPageContent(
	wg *sync.WaitGroup,
	details <-chan string,
	articles chan<- entity.Article,
) {
	defer wg.Done()

	url := <-details
	hub.CrawlDetailPage("agent", url)
	fmt.Println("  getPageContent:", url)
	// articles <- e
}
