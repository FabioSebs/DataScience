package main

import (
	"fabrzy/data/cleaning"
	"fabrzy/data/fish_boats"
	"fabrzy/data/fish_consumption"
	"fabrzy/data/fish_employment"
	"fabrzy/webscraper"
	"fmt"
)

// https://go-echarts.github.io/go-echarts/docs/line

func main() {
	// Checking Dataframes
	// merging.GetAllDataframes()
	// merging.GetFishConsID()

	//Visualizations

	webscraper.GeneratePie()

	fish_consumption.ConsumptionOverTime("Asia", "Afghanistan", "Africa", "Americas", "France", "India", "Japan")

	fish_boats.FishBoatsOverTime()

	// fish_prices.GenerateFishPrice()

	fish_employment.EmploymentOverTime("Africa", "Americas", "Asia", "Europe", "Oceania", "World")
	
	df := cleaning.Cleaning()
	fmt.Println(df)
}
