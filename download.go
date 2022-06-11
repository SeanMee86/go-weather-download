package main

import (
	"errors"
	"io"
	"log"
	"net/http"
	"os"
)

func downloadFile(url string, filename string) {
	resp, err := http.Get(url)

	if err != nil {
		log.Fatal("Errors: ", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatal(errors.New("received non-200 status code"))
	}

	f, err := os.Create(filename)

	if err != nil {
		log.Fatal("Errors", err)
	}
	defer f.Close()

	_, err = io.Copy(f, resp.Body)

	if err != nil {
		log.Fatal("Errors", err)
	}
}