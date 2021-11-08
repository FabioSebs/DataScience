package main

import (
	"fabrzy/data/fish_boats"
	"fabrzy/data/fish_consumption"
	"fabrzy/data/fish_employment"
	"fabrzy/webscraper"
)

// https://go-echarts.github.io/go-echarts/docs/line

func main() {
	fish_employment.GetData()
	// Getting Endangered Fish from Webscraper
	// webscraper.Webscraper()
	// webscraper.ReadJSON()

	//Fish Boats Data

	completed := make(chan bool, 3)

	go func() {
		webscraper.GeneratePie()
		completed <- true
	}()

	go func() {
		fish_consumption.ConsumptionOverTime("Asia", "Afghanistan", "Africa", "Americas", "France", "India", "Japan")
		completed <- true
	}()

	go func() {
		fish_boats.FishBoatsOverTime()
		completed <- true
	}()

	<-completed
	<-completed
	<-completed
}
