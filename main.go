package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	ex, err := os.Executable()
	if err != nil {
		fmt.Println("Errors: ", err)
	}
	wd := filepath.Dir(ex)
	u, pf, c, co, err := getImgUrl(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	tmpFile := wd+"/temp-weather.png"
	d := getDate()
	err = downloadFile(baseUrl + u, tmpFile)
	if err != nil {
		fmt.Println("Errors: ", err)
	}

	i, err := os.Open(tmpFile)
	if err != nil {
		fmt.Println("Errors: ", err)
	}

	generatePdf(i, d, pf, c, co, wd, os.Args[1])
	i.Close()
	err = os.Remove(tmpFile)

	if err != nil {
		fmt.Println("Errors: ", err)
	}
}

const baseUrl = "https://forecast.weather.gov/"
