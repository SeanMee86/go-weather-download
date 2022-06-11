package main

import (
	"log"
	"os"
	"path/filepath"
)

func main() {
	ex, err := os.Executable()

	if err != nil {
		log.Fatal("Errors: ", err)
	}

	dir := filepath.Dir(ex)
	tmpFile := dir+"\\temp-weather.png"
	d := getDate()
	pd := getPdfData(os.Args[1])
	absImgUrl := baseUrl + pd.imageUrl
	downloadFile(absImgUrl, tmpFile)

	f, err := os.Open(tmpFile)

	if err != nil {
		log.Fatal("Errors: ", err)
	}

	generatePdf(f, d, pd, dir)
	f.Close()
	os.Remove(tmpFile)
}

const baseUrl = "https://forecast.weather.gov/"
