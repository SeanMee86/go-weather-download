package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"golang.org/x/net/html"
)

type MyReader struct {}

func getImgUrl(siteUrl string) (string, string, string, string, error) {
	ex, err := os.Executable()
	if err != nil {
		return "", "", "", "", err
	}
	wd := filepath.Dir(ex)
	resp, _ := http.Get(siteUrl)
	f, err := os.Create(wd+"/weather.html")
	if err != nil {
		return "", "", "", "", err
	}
	io.Copy(f, resp.Body)
	defer resp.Body.Close()
	c, err := ioutil.ReadFile(wd+"/weather.html")
	if err != nil {
		return "", "", "", "", err
	}
	htmlTokens := html.NewTokenizer(strings.NewReader(string(c)))
	var finUrl string
	var pf string
	var city string
	var coords string
	counter := 0
	startCounter := false
	loop: for {
		tt := htmlTokens.Next()
		switch tt {
		case html.ErrorToken:
			break loop
		case html.TextToken:
			if htmlTokens.Token().Data == "Point Forecast:" {
				fmt.Println(htmlTokens.Token().Data)
				pf = htmlTokens.Token().Data
				startCounter = true
			}
			if startCounter {
				switch counter {
				case 0:
					city =  htmlTokens.Token().Data
				case 1:
					coords = htmlTokens.Token().Data
					startCounter = false
				}
				counter++
			}
		case html.StartTagToken:
			t := htmlTokens.Token()
			isImg := t.Data == "img"
			if isImg {
				if t.Attr[len(t.Attr)-1].Val == "#MouseVal" {
					finUrl = t.Attr[0].Val
				}
			}
		}
	}
	os.Remove(wd+"/weather.html")
	return finUrl, pf, city, coords, nil
}

func (m MyReader) Read(bs []byte) (n int, err error) {
	fmt.Println(string(bs))
	return 1, nil
}
