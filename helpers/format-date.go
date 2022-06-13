package helpers

import (
	"strings"
	"time"
)

func FormatDateForFilename() string {
	t := time.Now().Format("2006-Jan-02 Monday")
	return strings.Join(strings.Split(t, " "), "-")
}

func formatDateToTitle(date string) string {
	da := strings.Split(date, "-")
	return da[len(da)-1] + ", " + da[1] + " " + da[2] + ", " + da[0]
}