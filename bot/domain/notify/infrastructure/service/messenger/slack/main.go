package slack

import (
	artisan "bot/infrastructure/messenger/slack"
)

type (
	Slack struct {
		artisan  artisan.Slack
		channels map[string]string
	}
)

func NewSlack(token string, channels map[string]string) Slack {
	a := artisan.NewSlack(token)

	return Slack{
		artisan:  a,
		channels: channels,
	}
}

func (s Slack) Send(msg, channel string) {
	ch := s.channels[channel]
	s.artisan.Send(msg, ch)
}
