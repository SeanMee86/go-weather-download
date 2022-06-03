package main

import (
	"fmt"
	"os"
)

func main() {
	u := "https://forecast.weather.gov/meteograms/Plotter.php?lat=37.4353&lon=-122.0712&wfo=MTR&zcode=CAZ508&gset=18&gdiff=3&unit=0&tinfo=PY8&ahour=0&pcmd=10000010100000000000000000000000000000000000000000000000000&lg=en&indu=1!1!1!&dd=&bw=&hrspan=48&pqpfhr=6&psnwhr=6"
	tmpFile := "temp-weather.png"
	d := getDate()
	err := downloadFile(u, tmpFile)
	if err != nil {
		fmt.Println("Errors: ", err)
	}

	i, err := os.Open(tmpFile)
	if err != nil {
		fmt.Println("Errors: ", err)
	}

	generatePdf(d)
	i.Close()
	err = os.Remove(tmpFile)

	if err != nil {
		fmt.Println("Errors: ", err)
	}
}
