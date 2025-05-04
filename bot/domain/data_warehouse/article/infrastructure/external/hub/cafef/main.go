package cafef

import (
	_ "bytes"
	_ "encoding/json"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"

	"bot/domain/data_warehouse/article/entity"
	"bot/lib/array"
	"bot/lib/text"
)

type (
	Cafef struct {
		hubId  int
		domain string
	}
)

func NewHub(
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
			url := cff.domain + "/" + strings.TrimLeft(e1.Attr("href"), "/")

			p := entity.NewArticle(
				0,         // id
				0,         // mode
				title,     // title
				"",        // sapo
				"",        // content
				"",        // image
				"",        // publishedAt
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
						"",        // publishedAt
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
	domain := cff.domain

	result := cff.crawlTopPageMain(domain)

	r := cff.crawlTopPageSlave(domain + "/timelinehome/2.chn")
	result = append(result, r...)

	return result
}

func (cff Cafef) CrawlDetailPage(url string) entity.Article {
	c := colly.NewCollector()
	url = "https://cafef.vn/hoa-phat-no-vay-gan-90000-ty-tien-mat-xuong-thap-nhat-4-nam-18825050123344295.chn"
	// https://cafef.vn/gia-cuoc-van-tai-bien-tang-manh-cong-ty-so-huu-doi-tau-container-lon-nhat-viet-nam-bao-lai-tang-bang-lan-co-phieu-boc-dau-len-dinh-lich-su-188250501232158774.chn
	url = "https://cafef.vn/dau-thang-5-gui-tiet-kiem-tai-agribank-vietcombank-bidv-vietinbank-huong-lai-suat-cao-nhat-bao-nhieu-188250503075526502.chn"
	title := ""
	image := ""
	content := ""
	sapo := ""
	publishedAt := ""

	c.OnHTML(".totalcontentdetail", func(e *colly.HTMLElement) {
		e.ForEachWithBreak("h1", func(i int, e1 *colly.HTMLElement) bool {
			title = strings.TrimSpace(e1.Text)
			return false
		})

		e.ForEachWithBreak("p.dateandcat span.pdate", func(i int, e1 *colly.HTMLElement) bool {
			publishedAt = formatDateTime(strings.TrimSpace(e1.Text))
			return false
		})

		e.ForEachWithBreak(".w640", func(_ int, e1 *colly.HTMLElement) bool {
			e1.DOM.Children().Each(func(i int, s *goquery.Selection) {
				c, _ := s.Attr("class")
				wl := []string{"sapo", "media VCSortableInPreviewMode"}
				if !array.Contains(wl, c) {
					return
				}

				if c == "media VCSortableInPreviewMode" {
					image, _ = s.Find("img").Attr("src")
				}

				if c == "sapo" {
					sapo = strings.TrimSpace(s.Text())
				}
			})

			return false
		})
	})

	c.OnHTML("#mainContent .detail-content", func(e *colly.HTMLElement) {
		e.DOM.Children().Each(func(i int, s *goquery.Selection) {
			tagName := s.Nodes[0].Data

			wl := []string{"figure", "p"}
			if !array.Contains(wl, tagName) {
				return
			}

			if tagName == "figure" {
				img := s.Find("img")
				src, _ := img.Attr("src")
				t, _ := img.Attr("title")

				content += "<img src='" + src + "' title='" + t + "' />"
			}

			if tagName == "p" {
				content += text.Trim(text.StripTag(s.Text()))
			}

			content += "\n"
		})
	})

	c.Visit(url)
	write("url")
	write("\n")
	write(url)
	write("\n")
	write("\n")
	write("title")
	write("\n")
	write(title)
	write("\n")
	write("\n")
	write("publishedAt")
	write("\n")
	write(publishedAt)
	write("\n")
	write("\n")
	write("image")
	write("\n")
	write(image)
	write("\n")
	write("\n")
	write("sapo")
	write("\n")
	write(sapo)
	write("\n")
	write("\n")
	write("content")
	write("\n")
	write(content)
	write("\n")

	return entity.NewArticle(
		0,           // id
		0,           // mode
		title,       // title
		sapo,        // sapo
		content,     // content
		image,       // image
		publishedAt, // publishedAt
		url,         // origin
		cff.hubId,   // source id
	)
}

func formatDateTime(input string) string {
	layoutIn := "02-01-2006 - 03:04 PM" // layout for parsing
	layoutOut := "2006-01-02 15:04:05"  // desired output format

	t, err := time.Parse(layoutIn, input)
	if err != nil {
		return ""
	}

	return t.Format(layoutOut)
}

func write(text string) {
	// Open file with append and write-only flags. Create it if it doesn't exist.
	file, err := os.OpenFile("zzz.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Write additional content to the file
	_, err = file.WriteString(text)
	if err != nil {
		panic(err)
	}
}
