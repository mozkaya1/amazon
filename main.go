package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

type data struct {
	Url           string  `json:"url"`
	Price         float64 `json:"price"`
	Status        int     `json:"status"`
	PriceDiscount float64 `json:"discount"`
}

func (output *data) getPrice(url string, p float64) data {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
	}
	// More human requests ...
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-Site", "none")
	req.Header.Set("Sec-Fetch-User", "?1")

	// Bot pass through ...
	req.Header.Set("sec-ch-ua", `"Not_A Brand";v="8", "Chromium";v="120", "Google Chrome";v="120"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)

	resp, err := http.DefaultClient.Do(req)
	// _, err = io.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Println(err)
	// }

	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Println(err)
	// }

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Println(err)
	}

	defer resp.Body.Close()
	price := doc.Find(".a-price-whole").First().Text()
	// fmt.Println(string(output))
	output.Status = resp.StatusCode
	output.Url = url
	if price != "" {
		output.Price, err = strconv.ParseFloat(price, 64)
	}
	if err != nil {
		log.Println(err)
	}
	if output.Price < p {
		output.PriceDiscount = (p - output.Price) / p * 100
	}
	return *output
}

func main() {
	// url := flag.String("url", "https://www.amazon.nl/-/en/gp/product/B00OYQ2YOG?smid=A17D2BRD4YMT0X", "Set url to be Checked")
	url := flag.String("url", "https://www.amazon.com/TwinGrip-Pliers-Comfort-Grip-8-inch/dp/B097C7W2YK", "Set url to be Checked")
	priceAlert := flag.Float64("set", 10, "Set Alarm Price Threshold, expected below this threshold")
	flag.Parse()
	var output data
	data := output.getPrice(*url, *priceAlert)
	json.NewEncoder(os.Stdout).Encode(data)

}
