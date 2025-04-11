package nqs

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"

	"bot/domain/article/scraper/entity"
	_ "bot/lib/file"
)

const domain = "https://nguoiquansat.vn"

func CrawlTopPage() []entity.Article {
	c := colly.NewCollector()

	result := []entity.Article{}

	c.OnHTML(".c-head", func(e *colly.HTMLElement) {
		e.ForEach(".b-grid__title a", func(j int, e1 *colly.HTMLElement) {
			title := e1.Text
			url := strings.TrimLeft(e1.Attr("href"), "/")

			p := entity.NewArticle(
				0,     // id
				0,     // mode
				title, // title
				"",    // sapo
				"",    // content
				"",    // image
				url,   // origin
				0,     // source id
			)

			result = append(result, p)
		})
	})

	c.OnHTML(".l-content", func(e *colly.HTMLElement) {
		e.ForEach(".l-main .c-content-box .b-grid__content .b-grid__title a", func(j int, e1 *colly.HTMLElement) {
			title := e1.Text
			url := strings.TrimLeft(e1.Attr("href"), "/")

			p := entity.NewArticle(
				0,     // id
				0,     // mode
				title, // title
				"",    // sapo
				"",    // content
				"",    // image
				url,   // origin
				0,     // source id
			)

			result = append(result, p)
		})
	})

	c.Visit(domain)

	return result
}

func CrawlDetail(url string) entity.Article {
	c := colly.NewCollector()

	title := ""
	content := ""
	sapo := ""
	// timestamp := ""
	image := ""
	// c-news-detail
	// c.OnHTML(".c-news-detail .", func(e *colly.HTMLElement) {
	// 	e.ForEach("h1.title", func(i int, e1 *colly.HTMLElement) {
	// 		title += strings.TrimSpace(e1.Text)
	// 		fmt.Println("title:", title)
	// 	})

	// 	e.ForEach("p.dateandcat span.pdate", func(i int, e1 *colly.HTMLElement) {
	// 		timestamp += strings.TrimSpace(e1.Text)
	// 		fmt.Println("timestamp:", timestamp)
	// 	})
	// })

	// c.OnHTML(".w640", func(e *colly.HTMLElement) {
	// 	e.ForEach("h2.sapo", func(i int, e1 *colly.HTMLElement) {
	// 		caption += strings.TrimSpace(e1.Text)
	// 		fmt.Println("caption:", caption)
	// 	})
	// })

	c.OnHTML(".c-news-detail article", func(e *colly.HTMLElement) {
		e.ForEachWithBreak(".sc-longform-header-title", func(_ int, e1 *colly.HTMLElement) bool {
			title = e1.Text
			return false
		})

		e.ForEachWithBreak(".sc-longform-header-sapo", func(_ int, e1 *colly.HTMLElement) bool {
			sapo = e1.Text
			return false
		})

		// e.ForEachWithBreak(".sc-longform-header-meta", func(_ int, e1 *colly.HTMLElement) bool {
		// 	sapo = e1.Text
		// 	return false
		// })

		fmt.Println("title:", title)
		fmt.Println("sapo:", sapo)

		fmt.Println("content BEGIN")
		e.ForEach("p", func(_ int, e1 *colly.HTMLElement) {
			fmt.Println(e1.Text)

			content += strings.TrimSpace(e1.Text)
		})
		fmt.Println("content END")
	})

	c.Visit(url)

	return entity.NewArticle(
		0,       // id
		0,       // mode
		title,   // title
		sapo,    // sapo
		content, // content
		image,   // image
		url,     // origin
		0,       // source id
	)
}
