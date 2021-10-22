package webscraper

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gocolly/colly/v2"
)

func Colly() {
	// CSV File
	f := "fish_endangered.csv"
	file, err := os.Create(f)
	if err != nil {
		log.Fatal("error!")
	}
	defer file.Close()

	// Writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	//Colly
	c := colly.NewCollector(
		colly.AllowedDomains("https://www.fisheries.noaa.gov/species-directory/threatened-endangered"),
	)

	for i := 2; i < 6; i++ {
		fmt.Printf("Scraping Page: %d\n", i)
		c.Visit("https://www.fisheries.noaa.gov/species-directory/threatened-endangered?title=&species_category=any&species_status=any&regions=all&items_per_page=25&page=" + strconv.Itoa(i) + "&sort=")
	}

	c.OnHTML(".species-directory__headers--8col", func(e *colly.HTMLElement) {
		writer.Write([]string{
			e.ChildText("a"),
		})
	})

	log.Println("Scraping complete")
}
