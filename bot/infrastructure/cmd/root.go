package cmd

import (
	"os"

	// article "bot/domain/article/scraper/handler"
	dw "bot/domain/data_warehouse"
	notify "bot/domain/notify/handler"
)

func Execute() {
	params := os.Args[2:]
	switch os.Args[1] {
	case "notify":
		notify.Register(params...)
	// case "article":
	// 	article.Register(params...)
	case "dw": // data warehouse
		dw.Register(params...)
	default:
		os.Exit(1)
	}
}
