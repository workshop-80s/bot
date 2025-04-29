package messenger

type (
	Messenger interface {
		Send(msg string, channel string)
	}
)
