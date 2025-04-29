package slack

import (
	"github.com/parnurzeal/gorequest"
)

const token = "xoxb-8827139872881-8821913393826-oxIBVyv1RnDiXB6VycD4CiO4"

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
