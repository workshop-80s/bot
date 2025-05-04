package handler

import "os"

func Register(params ...string) {
	cmd := params[0]

	switch cmd {
	case "crawl":
		scrape()
	default:
		os.Exit(1)
	}
}
