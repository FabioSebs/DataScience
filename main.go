package main

import (
	"fabrzy/data/fish_catches"
	"fabrzy/data/fish_employment"
)

// https://go-echarts.github.io/go-echarts/docs/line

func main() {
	// Checking Dataframes
	// merging.GetAllDataframes()
	// // merging.GetFishConsID()

	// //Visualizations

	// webscraper.GeneratePie()

	// fish_consumption.ConsumptionOverTime("Indonesia", "China", "United States", "France")

	// fish_boats.FishBoatsOverTime()

	// fish_prices.GenerateFishPrice()

	fish_employment.EmploymentOverTime("Americas", "Asia")

	// cleaning.Cleaning()
	fish_catches.FishCatchesOverTime("Indonesia")

	// machine_learning.SajariRegression()
	// machine_learning.PlotModel()
}
