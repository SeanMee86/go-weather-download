package main

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"golang.org/x/net/html"
)

func getImgUrl(siteUrl string) (string, error) {
	ex, err := os.Executable()
	if err != nil {
		return "", err
	}
	wd := filepath.Dir(ex)
	resp, _ := http.Get(siteUrl)
	f, err := os.Create(wd+"/weather.html")
	if err != nil {
		return "", err
	}
	io.Copy(f, resp.Body)
	defer resp.Body.Close()
	c, err := ioutil.ReadFile(wd+"/weather.html")
	if err != nil {
		return "", err
	}
	htmlTokens := html.NewTokenizer(strings.NewReader(string(c)))
	var finUrl string
	loop: for {
		tt := htmlTokens.Next()
		switch tt {
		case html.ErrorToken:
			break loop
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
	return finUrl, nil
}