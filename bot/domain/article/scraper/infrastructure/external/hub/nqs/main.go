package nqs

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"

	"bot/domain/article/scraper/entity"
	"bot/lib/array"
	"bot/lib/text"
)

type (
	Nqs struct {
		hubId  int
		domain string
	}
)

func NewHub(
	hubId int,
	domain string,
) Nqs {
	return Nqs{
		hubId:  hubId,
		domain: domain,
	}
}

func (nqs Nqs) CrawlTopPage() []entity.Article {
	c := colly.NewCollector()

	result := []entity.Article{}

	c.OnHTML(".c-head", func(e *colly.HTMLElement) {
		e.ForEach(".b-grid", func(j int, e1 *colly.HTMLElement) {
			img := ""
			title := ""
			url := ""

			e1.ForEachWithBreak(".b-grid__img img", func(_ int, e2 *colly.HTMLElement) bool {
				img = e2.Attr("src")
				return false
			})

			e1.ForEachWithBreak(".b-grid__title a", func(_ int, e2 *colly.HTMLElement) bool {
				title = e2.Text
				url = strings.TrimLeft(e2.Attr("href"), "/")
				return false
			})

			fmt.Println("title:", title)
			fmt.Println("img:", img)
			fmt.Println("url:", url)
			fmt.Println()

			p := entity.NewArticle(
				0,         // id
				0,         // mode
				title,     // title
				"",        // sapo
				"",        // content
				img,       // image
				url,       // origin
				nqs.hubId, // source id
			)

			result = append(result, p)
		})
	})

	/*
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
	*/

	c.OnHTML(".l-main", func(e *colly.HTMLElement) {
		e.ForEach(".b-grid", func(j int, e1 *colly.HTMLElement) {
			img := ""
			title := ""
			url := ""

			e1.ForEachWithBreak(".b-grid__img img", func(_ int, e2 *colly.HTMLElement) bool {
				img = e2.Attr("src")
				return false
			})

			e1.ForEachWithBreak(".b-grid__title a", func(_ int, e2 *colly.HTMLElement) bool {
				title = e2.Text
				url = strings.TrimLeft(e2.Attr("href"), "/")
				return false
			})

			fmt.Println("title:", title)
			fmt.Println("img:", img)
			fmt.Println("url:", url)
			fmt.Println()

			p := entity.NewArticle(
				0,         // id
				0,         // mode
				title,     // title
				"",        // sapo
				"",        // content
				img,       // image
				url,       // origin
				nqs.hubId, // source id
			)

			result = append(result, p)
		})
	})
	c.Visit(nqs.domain)

	return result
}

func (nqs Nqs) CrawlDetailPage(url string) entity.Article {
	c := colly.NewCollector()

	title := ""
	content := ""
	sapo := ""
	timestamp := ""
	thumbnail := ""

	paragraph := []string{}

	c.OnHTML(".c-news-detail article", func(e *colly.HTMLElement) {
		e.ForEachWithBreak(".sc-longform-header-text", func(_ int, e1 *colly.HTMLElement) bool {
			e1.ForEachWithBreak(".sc-longform-header-title", func(_ int, e2 *colly.HTMLElement) bool {
				title = e2.Text
				return false
			})

			e1.ForEachWithBreak(".sc-longform-header-sapo", func(_ int, e2 *colly.HTMLElement) bool {
				sapo = e2.Text

				return false
			})

			e1.ForEachWithBreak(".block-sc-publish-time", func(_ int, e2 *colly.HTMLElement) bool {
				timestamp = formatDateTime(e2.Text)
				fmt.Println(e2.Text)
				return false
			})
			// class="sc-longform-header-date block-sc-publish-time"
			return false
		})

		e.DOM.Children().Each(func(i int, s *goquery.Selection) {
			tagName := goquery.NodeName(s)

			wl := []string{"figure", "p", "div"}
			if !array.Contains(wl, tagName) {
				return
			}

			if tagName == "figure" {
				img := s.Find("img")
				src, _ := img.Attr("src")
				t, _ := img.Attr("title")

				text := "<img src='" + src + "' title='" + t + "' />"
				paragraph = append(paragraph, text)
			}

			if tagName == "p" {
				text := text.Trim(text.StripTag(s.Text()))
				paragraph = append(paragraph, text)
			}
		})
	})

	c.Visit(url)
	content = strings.Join(paragraph[:len(paragraph)-1], "\n\n")
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
	write("timestamp")
	write("\n")
	write(timestamp)
	write("\n")
	write("\n")
	write("image")
	write("\n")
	write(thumbnail)
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
		0,         // id
		0,         // mode
		title,     // title
		sapo,      // sapo
		content,   // content
		thumbnail, // image
		url,       // origin
		nqs.hubId, // source id
	)
}

func formatDateTime(input string) string {
	layoutIn := "02/01/2006 15:04"     // layout for parsing
	layoutOut := "2006-01-02 15:04:05" // desired output format

	t, err := time.Parse(layoutIn, input)
	if err != nil {
		fmt.Println(err)
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
