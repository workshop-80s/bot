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

func getDetailURL(
	wg *sync.WaitGroup,
	agents <-chan string,
	// agent string,
	details chan<- string,
) {
	defer wg.Done()

	agent := <-agents

	fmt.Println("  getDetailURL - ", agent)
	urls := hub.CrawlTopPage(agent)
	for _, u := range urls {
		details <- u.Url()
	}
}

func (s Scraper) getList() {
	const numWorkers = 20

	var wg sync.WaitGroup

	details := make(chan string)
	agents := make(chan string, 10)

	hubs := s.articleHub.Find()

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go getDetailURL(&wg, agents, details)
	}

	go func() {
		fmt.Println("Wait BEGIN")

		wg.Wait()
		close(details)
		close(agents)

		fmt.Println("Wait END")
	}()

	for _, h := range hubs {
		agents <- h.Code()
	}

	fmt.Println("   getList out BEGIN")
	r := []string{}
	for v := range details {
		r = append(r, v)
	}

	fmt.Println("   getList out: ", len(r))
	fmt.Println("   getList out END")
}

func (s Scraper) getDetail(
	ch chan entity.Article,
	agent string,
	url string,
) {
	e := hub.CrawlDetailPage(agent, url)
	ch <- e
}
