package usecase

import (
	"github.com/google/wire"

	repository "bot/domain/notify/infrastructure/repository"
	repositoryI "bot/domain/notify/repository"

	"bot/domain/notify/infrastructure/service/messenger/slack"
	messengerI "bot/domain/notify/service/messenger"
)

type (
	Slack struct {
		repository repositoryI.Article
		messenger  messengerI.Messenger
	}
)

var ProviderSlack = wire.NewSet(
	NewSlack,
	repository.NewArticle,
	wire.Bind(new(repositoryI.Article), new(repository.Article)),

	slack.NewSlack,
	wire.Bind(new(messengerI.Messenger), new(slack.Slack)),
)

func NewSlack(
	r repositoryI.Article,
	m messengerI.Messenger,
) Slack {
	return Slack{
		repository: r,
		messenger:  m,
	}
}

func (s Slack) Send() {
	articles := s.repository.Find()
	// send mail

	// slack.Send("tao test notify", "C08QB471DQ9")
	for _, a := range articles {
		msg := a.Title() + "\n" + a.Url() + ""
		s.messenger.Send(msg, "general")
	}
}
