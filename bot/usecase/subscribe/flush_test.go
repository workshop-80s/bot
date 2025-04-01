package subscribe

import (
	"testing"

	"bot/domain/subscribe/entity"
	repository "bot/domain/subscribe/repository"
)

func TestFlush(t *testing.T) {
	mock := new(repository.MockArticle)

	fakeData := []entity.Article{
		entity.NewArticle(
			"title1",
			"url1",
		),
		entity.NewArticle(
			"title2",
			"url2",
		),
	}

	mock.
		On("Find").
		Return(fakeData).
		Once()

	uc := NewFlush(mock)
	uc.Flush()

	t.Run("repository.Find is called", func(t *testing.T) {
		mock.AssertNumberOfCalls(t, "Find", 1)
	})
}
