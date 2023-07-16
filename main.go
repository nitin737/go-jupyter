package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"

	"github.com/gocolly/colly"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQURSTUVWXYZ"

func main() {
	fName := "crypto-list.csv"
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("Cannot create file %q: %s\n", fName, err)
		return
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	//Write CSV header
	writer.Write([]string{"Name", "Market Cap", "Price", "Circulating Supply", "Volume(24)", "%1h", "%24h", "%7d"})

	c := colly.NewCollector(
		colly.AllowedDomains("coinmarketcap.com"),
		colly.UserAgent("user-agent"),
		colly.AllowURLRevisit(),
		//colly.Debugger(&debug.LogDebugger{}),
	)

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", RandomString())
	})
	c.OnError(func(r *colly.Response, err error) {
		log.Println("Request URL:", r.Request.URL, "failed with response:", r.Body, "\nError:", err)
	})
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Status: ", r.StatusCode)
	})

	c.OnHTML(".cmc-table__table-wrapper-outer tbody tr", func(e *colly.HTMLElement) {
		writer.Write([]string{
			e.ChildText("a.cmc-table__column-name--name"),
			e.ChildText(".jYSZLP"),
			e.ChildText(".iVdfNf"),
			e.ChildText(".cmc-table__cell--sort-by__circulating-supply"),
			e.ChildText(".cmc-table__cell--sort-by__volume-24-h"),
			e.ChildText(".cmc-table__cell--sort-by__percent-change-1-h"),
			e.ChildText(".cmc-table__cell--sort-by__percent-change-24-h"),
			e.ChildText(".cmc-table__cell--sort-by__percent-change-7-d"),
		})

		//c.Visit(e.Request.AbsoluteURL())
	})
	c.Visit("https://coinmarketcap.com/all/views/all/")

}
func RandomString() string {
	b := make([]byte, rand.Intn(10)+10)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
