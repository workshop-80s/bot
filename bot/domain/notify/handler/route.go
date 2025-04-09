package handler

import (
	"fmt"
	"os"
)

func Register() {
	switch os.Args[2] {
	case "slack":
		slack()
	default:
		fmt.Println("default")
		os.Exit(1)
	}
}
