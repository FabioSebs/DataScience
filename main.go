package main

import (
	"fabrzy/data/fish_boats"
	"fabrzy/data/fish_consumption"
	"fabrzy/data/fish_employment"
	"fabrzy/data/fish_prices"
	"fabrzy/data/merging"
	"fabrzy/webscraper"
)

// https://go-echarts.github.io/go-echarts/docs/line

func main() {
	// Checking Dataframes
	merging.GetAllDataframes()
	merging.GetFishConsID()

	//Visualizations

	completed := make(chan bool, 5)

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

	go func() {
		fish_prices.GenerateFishPrice()
		completed <- true
	}()

	go func() {
		fish_employment.EmploymentOverTime("Africa", "Americas", "Asia", "Europe", "Oceania", "World")
		completed <- true
	}()

	<-completed
	<-completed
	<-completed
	<-completed
	<-completed
}
