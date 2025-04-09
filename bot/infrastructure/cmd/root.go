package cmd

import (
	"fmt"
	"os"

	notify "bot/domain/notify/handler"
)

func Execute() {
	switch os.Args[1] {
	case "notify":
		notify.Register()
	case "article":
		fmt.Println("article")
	default:
		fmt.Println("default")
		os.Exit(1)
	}
}
