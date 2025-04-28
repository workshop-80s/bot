package file

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func Download(
	url,
	destination string,
) error {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	out, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)

	elapsed := time.Since(start)
	fmt.Printf("download took %s\n", elapsed)

	return err
}