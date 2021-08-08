package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/gocolly/colly"
)

type Fact struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
}

func main() {
	allFacts := make([]Fact, 0)
	collector := colly.NewCollector(
		colly.AllowedDomains("factretriever.com", "www.factretriever.com"),
	)

	collector.OnHTML(".factsList li", func(element *colly.HTMLElement) {
		factId, err := strconv.Atoi(element.Attr("id"))
		if err != nil {
			log.Fatal("Could not get id")
		}
		factDesc := element.Text

		fact := Fact{
			ID:          factId,
			Description: factDesc,
		}

		allFacts = append(allFacts, fact)
	})

	collector.OnRequest(func(r *colly.Request) {
		log.Println("Visitingi:", r.URL.String())
	})
	collector.Visit("https://www.factretriever.com/rhino-facts")

	//enc := json.NewEncoder(os.Stdout)
	//enc.SetIndent("", " ")
	//enc.Encode(allFacts)

	writeJSON(allFacts)
}

func writeJSON(data []Fact) {
	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Fatal(err)
	}

	_ = ioutil.WriteFile("rhinofacts.json", file, 0644)
}
