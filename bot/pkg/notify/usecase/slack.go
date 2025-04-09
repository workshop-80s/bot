package usecase

import (
	"github.com/google/wire"

	repository "bot/pkg/notify/infrastructure/repository"
	repositoryI "bot/pkg/notify/repository"
)

type (
	Slack struct {
		repository repositoryI.Article
	}
)

var ProviderSlack = wire.NewSet(
	NewSlack,
	repository.NewArticle,
	wire.Bind(new(repositoryI.Article), new(repository.Article)),
)

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
