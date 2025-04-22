package cafef

import (
	_ "bytes"
	_ "encoding/json"
	"strings"

	"github.com/gocolly/colly"

	"bot/domain/article/scraper/entity"
)

const domain = "https://cafef.vn"

type (
	Cafef struct {
		hubId  int
		domain string
	}
)

func NewCafef(
	hubId int,
	domain string,
) Cafef {
	return Cafef{
		hubId:  hubId,
		domain: domain,
	}
}

func (cff Cafef) crawlTopPageSlave(link string) []entity.Article {
	c := colly.NewCollector()

	result := []entity.Article{}

	c.OnHTML("div.listchungkhoannew div.knswli-right h3", func(e *colly.HTMLElement) {
		e.ForEach("a", func(i int, e1 *colly.HTMLElement) {
			class := e1.Attr("class")
			if class != "" {
				return
			}

			title := e1.Text
			url := domain + "/" + strings.TrimLeft(e1.Attr("href"), "/")

			p := entity.NewArticle(
				0,         // id
				0,         // mode
				title,     // title
				"",        // sapo
				"",        // content
				"",        // image
				url,       // origin
				cff.hubId, // source id
			)
			result = append(result, p)
		})
	})

	c.Visit(link)
	return result
}

func (cff Cafef) crawlTopPageMain(link string) []entity.Article {
	c := colly.NewCollector()

	result := []entity.Article{}

	c.OnHTML(".news_left", func(e *colly.HTMLElement) {
		paths := []string{
			"div.top_noibat", // tin noi bat
			"div.listchungkhoannew div.knswli-right h3", // tin chinh
		}
		for _, p := range paths {
			e.ForEach(p, func(j int, e2 *colly.HTMLElement) {
				e2.ForEach("a", func(i int, e1 *colly.HTMLElement) {
					class := e1.Attr("class")
					if class != "" {
						return
					}

					title := e1.Attr("title")
					url := cff.domain + "/" + strings.TrimLeft(e1.Attr("href"), "/")

					p := entity.NewArticle(
						0,         // id
						0,         // mode
						title,     // title
						"",        // sapo
						"",        // content
						"",        // image
						url,       // origin
						cff.hubId, // source id
					)
					result = append(result, p)
				})
			})
		}
	})

	c.Visit(link)

	return result
}

func (cff Cafef) CrawlTopPage() []entity.Article {
	result := cff.crawlTopPageMain(domain)

	r := cff.crawlTopPageSlave(domain + "/timelinehome/2.chn")
	result = append(result, r...)

	return result
}

func (cff Cafef) CrawlDetailPage(url string) entity.Article {
	c := colly.NewCollector()

	title := ""
	content := ""
	caption := ""
	timestamp := ""

	c.OnHTML(".totalcontentdetail", func(e *colly.HTMLElement) {
		e.ForEach("h1.title", func(i int, e1 *colly.HTMLElement) {
			title += strings.TrimSpace(e1.Text)
			// fmt.Println("title:", title)
		})

		e.ForEach("p.dateandcat span.pdate", func(i int, e1 *colly.HTMLElement) {
			timestamp += strings.TrimSpace(e1.Text)
			// fmt.Println("timestamp:", timestamp)
		})
	})

	c.OnHTML(".w640", func(e *colly.HTMLElement) {
		e.ForEach("h2.sapo", func(i int, e1 *colly.HTMLElement) {
			caption += strings.TrimSpace(e1.Text)
			// fmt.Println("caption:", caption)
		})
	})

	c.OnHTML("#mainContent", func(e *colly.HTMLElement) {
		e.ForEachWithBreak(".VCSortableInPreviewMode img", func(_ int, e1 *colly.HTMLElement) bool {
			// thumbnail = e1.Attr("src")
			return false
		})

		// fmt.Println("content BEGIN")
		e.ForEach(".detail-content > p", func(_ int, e1 *colly.HTMLElement) {
			// fmt.Println(e1.Text)

			content += strings.TrimSpace(e1.Text)
		})
		// fmt.Println("content END")
	})

	c.Visit(url)

	return entity.NewArticle(
		0,         // id
		0,         // mode
		title,     // title
		caption,   // sapo
		content,   // content
		"",        // image
		url,       // origin
		cff.hubId, // source id
	)
}
