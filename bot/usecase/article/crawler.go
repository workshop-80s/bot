package article

import "fmt"


func Crawl() {
	sources := []string{
		"cafef",
		"nqs",
	}

	for _, src := range sources {
		fmt.Println(src)
	}
}