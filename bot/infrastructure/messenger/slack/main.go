package slack

import (
	"github.com/parnurzeal/gorequest"
)

type (
	Slack struct {
		botToken string
	}
)

func NewSlack(token string) Slack {
	return Slack{
		botToken: token,
	}
}

func (s Slack) Send(
	msg string,
	channel string,
) {
	url := "https://slack.com/api/chat.postMessage"

	payload := map[string]string{
		"channel": channel,
		"text":    msg,
	}
	request := gorequest.New()
	request.
		Post(url).
		Set("Authorization", "Bearer "+token).
		Send(payload).
		End()
}
