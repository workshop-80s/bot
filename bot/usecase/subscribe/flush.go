package subscribe

import (
	"github.com/google/wire"

	"bot/domain/subscribe/infrastructure/repository"
	repositoryI "bot/domain/subscribe/repository"
)

type (
	Flush struct {
		repository repositoryI.Article
	}
)

var ProviderFlush = wire.NewSet(
	NewFlush,
	repository.NewArticle,
	wire.Bind(new(repositoryI.Article), new(repository.Article)),
)

func NewFlush(r repositoryI.Article) Flush {
	return Flush{
		repository: r,
	}
}

func (s Flush) Flush() {
	// fetch DB

	s.repository.Find()
	// send mail
}
