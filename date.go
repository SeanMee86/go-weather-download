package main

import (
	"strings"
	"time"
)

func getDate() string {
	t := time.Now().Format("2006-Jan-02 Monday")
	return strings.Join(strings.Split(t, " "), "-")
}