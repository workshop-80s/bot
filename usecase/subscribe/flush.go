package subscribe

import (
	"bot/domain/subscribe/repository"
	"fmt"
)

type (
	subscribeFlush struct {
		repository repository.Article
	}
)

func NewSubscribeFlush(r repository.Article) subscribeFlush {
	return subscribeFlush{
		repository: r,
	}
}

func (s subscribeFlush) Flush() {
	fmt.Println("BEGIN usecase")
	s.repository.Find()

	fmt.Println("END usecase")

	// fetch DB

	// send mail
}
