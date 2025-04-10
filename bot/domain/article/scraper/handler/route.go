package handler

func Register(params ...string) {
	cmd := params[0]
	switch cmd {
	case "crawl":
		scrape()
	default:
		scrape()
	}
}
