package cmd

import (
	"fmt"
	"os"

	article "bot/domain/article/scraper/handler"
	notify "bot/domain/notify/handler"
)

func Execute() {
	switch os.Args[1] {
	case "notify":
		notify.Register()
	case "article":
		article.Register()
	default:
		fmt.Println("default")
		os.Exit(1)
	}
}
