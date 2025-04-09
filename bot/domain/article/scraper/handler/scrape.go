package handler

import (
	"flag"
	"os"

	"bot/infrastructure/storage/database"
	"bot/lib"
)

func scrape() {
	command := flag.NewFlagSet("cmd", flag.ExitOnError)

	command.Parse(os.Args[2:])

	config := lib.GetEnvConfigMap("db")

	db := database.Connect(config)
	defer func() {
		database.Disconnect(db)
	}()

	uc := InitScraper(db)
	uc.Scrape()
}
