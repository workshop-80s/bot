package data_warehouse

import (
	"os"

	article "bot/domain/data_warehouse/article/handler"
)

func Register(params ...string) {
	domain := params[0]

	p := params[1:]
	switch domain {
	case "article":
		article.Register(p...)
	default:
		os.Exit(1)
	}
}
