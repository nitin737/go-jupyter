package main

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/gocolly/colly"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQURSTUVWXYZ"

func main() {
	fmt.Println("main method...")
	c := colly.NewCollector(
		colly.AllowedDomains("www.nseindia.com"),
		colly.UserAgent("user-agent"),
		colly.AllowURLRevisit(),
	)

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", RandomString())
		fmt.Println("Visiting", r.URL)
	})
	c.OnError(func(r *colly.Response, err error) {
		log.Println("Request URL:", r.Request.URL, "failed with response:", r.Body, "\nError:", err)
	})
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited", r.Request.URL)
	})
	c.OnHTML("*", func(e *colly.HTMLElement) {
		fmt.Println(e)
		//c.Visit(e.Request.AbsoluteURL())
	})
	c.Visit("https://www.nseindia.com/market-data/live-equity-market?symbol=NIFTY%2050")

}
func RandomString() string {
	b := make([]byte, rand.Intn(10)+10)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
