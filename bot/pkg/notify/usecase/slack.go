package usecase

import (
	// "github.com/google/wire"

	repositoryI "bot/pkg/notify/repository"
)

type (
	Slack struct {
		repository repositoryI.Article
	}
)

// var ProviderFlush = wire.NewSet(
// 	NewFlush,
// 	repository.NewArticle,
// 	wire.Bind(new(repositoryI.Article), new(repository.Article)),
// )

func NewSlack(r repositoryI.Article) Slack {
	return Slack{
		repository: r,
	}
}

func (s Slack) Send() {
	// fetch DB

	s.repository.Find()
	// send mail
}
