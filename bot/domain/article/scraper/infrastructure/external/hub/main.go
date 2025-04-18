package hub

import (
	"bot/domain/article/scraper/entity"
	"strconv"
)

func CrawlTopPage(agent string) []entity.Article {
	result := []entity.Article{}

	for i := 1; i <= 5; i++ {
		p := entity.NewArticle(
			i,                               // id
			0,                               // mode
			agent+"-title-"+strconv.Itoa(i), // title
			"",                              // sapo
			"",                              // content
			"",                              // image
			agent+"-url-"+strconv.Itoa(i),   // title
			0,                               // source id
		)
		result = append(result, p)
	}
	return result
}

func CrawlDetailPage(agent, url string) entity.Article {
	return entity.NewArticle(
		0,                  // id
		0,                  // mode
		agent+"title-"+url, // title
		"",                 // sapo
		"",                 // content
		"",                 // image
		url,                // title
		0,                  // source id
	)
}

/* backend before apply channel
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

func CrawlDetailPage(agent, url string) entity.Article {
	switch agent {
	case "cafef":
		return cafef.CrawlDetailPage(url)
	case "nqs":
		return nqs.CrawlDetailPage(url)
	}
	return entity.Article{}
}
*/
