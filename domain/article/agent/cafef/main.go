package cafef

import (
	_ "bytes"
	_ "encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gocolly/colly"

	"bot/domain/article/crawler/entity"
	"bot/lib/file"
	_ "bot/lib/file"
)

const domain = "https://cafef.vn"

func crawlTopPageSlave(link string) []entity.Page {
	c := colly.NewCollector() 

	result := []entity.Page{}
	
	c.OnHTML("div.listchungkhoannew div.knswli-right h3", func(e *colly.HTMLElement) {
		e.ForEach("a", func(i int, e1 *colly.HTMLElement) {
			class := e1.Attr("class");
			if (class != "") {
				return;
			}

			// title := e1.Attr("title")
			title := e1.Text
			url := domain + "/" + strings.TrimLeft(e1.Attr("href"), "/")

			p := entity.NewPage(
				url,
				title,
				"",
				"",
			)

			result = append(result, p);
		});
	})
	
	c.Visit(link);
	return result
}

func crawlTopPageMain(link string) []entity.Page {
	c := colly.NewCollector()

	result := []entity.Page{}
	
	c.OnHTML(".news_left", func(e *colly.HTMLElement) {
		paths := []string{
			"div.top_noibat", // tin noi bat
			"div.listchungkhoannew div.knswli-right h3", // tin chinh
		}
		for _, p := range paths {
			e.ForEach(p, func(j int, e2 *colly.HTMLElement) {
				e2.ForEach("a", func(i int, e1 *colly.HTMLElement) {
					class := e1.Attr("class");
					if (class != "") {
						return;
					}

					title := e1.Attr("title")
					url := domain + "/" + strings.TrimLeft(e1.Attr("href"), "/")

					p := entity.NewPage(
						url,
						title,
						"",
						"",
					)

					result = append(result, p);
				});
			});
		}
	})
	
	c.Visit(link)

	return result
}

func CrawlTopPage(link string) []entity.Page {
	result := crawlTopPageMain(link);

	r := crawlTopPageSlave(link + "/timelinehome/2.chn");
	result = append(result, r...);
	
	return result;
}


func CrawlDetail(link string) entity.Page {
	c := colly.NewCollector()

	title := ""
	content := ""
	caption := ""
	timestamp := ""
	thumbnail := ""

	c.OnHTML(".totalcontentdetail", func(e *colly.HTMLElement) {
		e.ForEach("h1.title", func(i int, e1 *colly.HTMLElement) {
			title += strings.TrimSpace(e1.Text)
			fmt.Println("title:", title)
		})

		e.ForEach("p.dateandcat span.pdate", func(i int, e1 *colly.HTMLElement) {
			timestamp += strings.TrimSpace(e1.Text)
			fmt.Println("timestamp:", timestamp)
		})
	})
	
	c.OnHTML(".w640", func(e *colly.HTMLElement) {
		e.ForEach("h2.sapo", func(i int, e1 *colly.HTMLElement) {
			caption += strings.TrimSpace(e1.Text)
			fmt.Println("caption:", caption)
		})
	})

	c.OnHTML("#mainContent", func(e *colly.HTMLElement) {
		e.ForEachWithBreak(".VCSortableInPreviewMode img", func(_ int, e1 *colly.HTMLElement) bool {
			thumbnail = e1.Attr("src");
			return false
		})

		fmt.Println("content BEGIN")
		e.ForEach(".detail-content > p", func(_ int, e1 *colly.HTMLElement) {
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

func downloadThumbnail(p entity.Page) {
	thumbnail := p.Thumbnail()

	path, err := os.Getwd()
    if err != nil {
        panic(err)
    }

	destination := path + "/tmp/img/" + filepath.Base(thumbnail);

	file.Download(thumbnail, destination)
}

func Crawl() {
	link := "https://cafef.vn/loai-nuoc-uong-co-chua-hang-ty-hat-vi-nhua-nguoi-viet-rat-me-dang-uong-hang-ngay-188250303105421555.chn"

	started := time.Now()
	CrawlDetail(link);

	elapsed := time.Since(started)
	fmt.Printf("total took %s", elapsed);
}

// crawl top-page
func CrawlTop() {
	started := time.Now()
	r := CrawlTopPage(domain);

	for _, e := range r {
		fmt.Println("url: ", e.Url())
	}

	elapsed := time.Since(started)
	fmt.Println("total items: ", len(r))
	fmt.Printf("total took %s", elapsed);
}