package crawler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/gocolly/colly"
)

type Link struct {
	URL string `json:"URL"`
}

func GetLinks(site string) error {
	links := make([]Link, 0)
	c := colly.NewCollector()

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		href := e.Attr("href")
		link := Link{URL: href}
		links = append(links, link)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	if err := c.Visit(site); err != nil {
		log.Fatal(err)
	}

	writeJSON(links)

	return nil
}

func writeJSON(data []Link) {
	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Fatal(err)
	}

	_ = ioutil.WriteFile("links.json", file, 0644)
}
