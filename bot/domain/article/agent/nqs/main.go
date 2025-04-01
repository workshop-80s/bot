package cafef

import (
	_ "bytes"
	_ "encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/gocolly/colly"

	"bot/domain/article/crawler/entity"
	_ "bot/lib/file"
)

const domain = "https://nguoiquansat.vn"


func crawlTopPage(link string) []entity.Page {
	c := colly.NewCollector()

	// resp, _ := http.Get(link)
	// destination := "z.txt";
	// defer resp.Body.Close()
	// out, _ := os.Create(destination)
	
	// defer out.Close()

	// // Write the body to file
	// io.Copy(out, resp.Body)


	result := []entity.Page{}
	
	c.OnHTML(".c-head", func(e *colly.HTMLElement) {
		e.ForEach(".b-grid__title a", func(j int, e1 *colly.HTMLElement) {
			title := e1.Text;
			url := strings.TrimLeft(e1.Attr("href"), "/")

			fmt.Println("title:", title)
			fmt.Println("url:", url)
			fmt.Println()
				p := entity.NewPage(
					url,
					title,
					"",
					"",
				)

			result = append(result, p);
		});
	})

	c.OnHTML(".l-content", func(e *colly.HTMLElement) {
		e.ForEach(".l-main .c-content-box .b-grid__content .b-grid__title a", func(j int, e1 *colly.HTMLElement) {
			title := e1.Text
			url := strings.TrimLeft(e1.Attr("href"), "/")

			fmt.Println("title:", title)
			fmt.Println("url:", url)
			fmt.Println()
				p := entity.NewPage(
					url,
					title,
					"",
					"",
				)

			result = append(result, p);
		});
	})

	c.Visit(link)

	return result
}

func CrawlTopPage(link string) []entity.Page {
	result := crawlTopPage(link);

	return result;
}



func CrawlDetail(link string) entity.Page {
	// resp, _ := http.Get(link)
	// destination := "z.html";
	// defer resp.Body.Close()
	// out, _ := os.Create(destination)
	
	// defer out.Close()

	// // Write the body to file
	// io.Copy(out, resp.Body)

	c := colly.NewCollector()

	title := ""
	content := ""
	sapo  := ""
	// timestamp := ""
	thumbnail := ""
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
	
	c.Visit(link)
	
	return entity.NewPage(
		link,
		title,
		content,
		thumbnail,
	)
}

// func Crawl() {
// 	started := time.Now()
// 	CrawlTop();

// 	elapsed := time.Since(started)
// 	fmt.Printf("total took %s", elapsed);
// }

// crawl top-page
func Crawl() {
	started := time.Now()
	// r := CrawlTopPage(domain);

	link := "https://nguoiquansat.vn/my-ngung-vien-tro-ukraine-chau-au-co-du-suc-ganh-vac-202587.html"
	CrawlDetail(link);

	// for _, e := range r {
	// 	fmt.Println("urlxxxxxx: ", e.Url())
	// }

	elapsed := time.Since(started)
	// fmt.Println("total items: ", len(r))
	fmt.Printf("total took %s", elapsed);
}