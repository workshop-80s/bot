package handler

import (
	"fmt"
	"os"
)

func Register(params ...string) {
	cmd := params[0]

	switch cmd {
	case "slack":
		slack()
	default:
		fmt.Println("default")
		os.Exit(1)
	}
}
