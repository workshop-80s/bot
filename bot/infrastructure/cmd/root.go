package cmd

import (
	"os"

	article "bot/domain/article/scraper/handler"
	notify "bot/domain/notify/handler"
)

func Execute() {
	params := os.Args[2:]

	switch os.Args[1] {
	case "notify":
		notify.Register(params...)
	case "article":
		article.Register(params...)
	default:
		os.Exit(1)
	}
}
