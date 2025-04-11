package usecase

import (
	"testing"

	"bot/domain/article/scraper/entity"
	repository "bot/domain/article/scraper/repository"
)

func TestCrawl(t *testing.T) {
	mockArticle := new(repository.MockArticle)
	mockArticleHub := new(repository.MockArticleHub)

	fakeArticleData := []entity.Article{
		entity.NewArticle(
			0,
			0,
			"title1",
			"sapo1",
			"content1",
			"image1",
			"origin1",
			1,
		),
		entity.NewArticle(
			0,
			0,
			"title2",
			"sapo2",
			"content2",
			"image2",
			"origin2",
			1,
		),
	}

	fakeArticleHubData := []entity.ArticleHub{}

	mockArticle.
		On("Find").
		Return(fakeArticleData).
		Once()

	mockArticleHub.
		On("Find").
		Return(fakeArticleHubData).
		Once()

	uc := NewScraper(
		mockArticleHub,
		mockArticle,
	)

	uc.Crawl()

	t.Run("repository.Find is called", func(t *testing.T) {
		mockArticle.AssertNumberOfCalls(t, "Find", 1)
	})
}
