package main

import (
	"fabrzy/data/fish_consumption"
	"fabrzy/webscraper"
)

// https://go-echarts.github.io/go-echarts/docs/line

func main() {
	// webscraper.Webscraper()
	// webscraper.ReadJSON()
	completed := make(chan bool, 2)

	go func() {
		webscraper.GeneratePie()
		completed <- true
	}()

	go func() {
		fish_consumption.ConsumptionOverTime("Asia", "Afghanistan", "Africa", "Americas", "France", "India", "Japan")
		completed <- true
	}()

	<-completed
	<-completed
}
