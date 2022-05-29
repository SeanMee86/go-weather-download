package main

import (
	"errors"
	"io"
	"net/http"
	"os"
)

func main() {
	u := "https://www.webcarpenter.com/pictures/Go-gopher-programming-language.jpg"
	fn := "gopher.jpg"
	downloadFile(u, fn)
}

func downloadFile(url string, filename string) error {
	resp, err := http.Get(url)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return errors.New("received non-200 status code")
	}

	f, err := os.Create(filename)

	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.Copy(f, resp.Body)

	if err != nil {
		return err
	}

	return nil
} 