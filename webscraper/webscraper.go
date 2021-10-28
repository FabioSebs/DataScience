package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

type Fish struct {
	Species string `json:"species"`
	Status  string `json:"status"`
	Year    string `json:"year"`
	Region  string `json:"region"`
}

func main() {
	fishStruct := make([]Fish, 0)
	defer fmt.Println(fishStruct)

	collector := colly.NewCollector(
		colly.AllowedDomains("fisheries.noaa.gov", "www.fisheries.noaa.gov"),
	)

	collector.OnHTML("div.species-directory__species--8col", func(element *colly.HTMLElement) {
		species := element.DOM
		fish := Fish{
			Species: species.Find("div.species-directory__species-title--name").Text(),
			Status:  strings.TrimSuffix(strings.TrimSpace(species.Find("div.species-directory__species-status-row").Find("div.species-directory__species-status").Text()), "\n"),
			Year:    strings.TrimSuffix(strings.TrimSpace(species.Find("div.species-directory__species-status-row").Find("div.species-directory__species-year").Text()), "\n"),
			Region:  strings.TrimSuffix(strings.TrimSpace(species.Find("div.species-directory__species-status-row").Find("div.species-directory__species-region").Text()), "\n"),
		}
		fishStruct = append(fishStruct, fish)
	})

	collector.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting", request.URL.String())
	})

	for x := 1; x <= 5; x++ {
		collector.Visit("https://www.fisheries.noaa.gov/species-directory/threatened-endangered?title=&species_category=any&species_status=any&regions=all&items_per_page=25&page=" + strconv.Itoa(x) + "&sort=")
	}

	writeJSON(fishStruct)
}

func writeJSON(data []Fish) {
	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println("Unable to create JSON file")
		return
	}

	_ = ioutil.WriteFile("endangeredFish.json", file, 0644)
}
