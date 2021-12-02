package merging

import (
	"fabrzy/data/fish_boats"
	"fabrzy/data/fish_catches"
	"fabrzy/data/fish_consumption"
	"fabrzy/data/fish_prices"
	"fmt"
	"log"
	"os"

	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

// MERGING EVERYTHING TOGETHER WITH GOTA 
func GetAllDataframes() {
	// LOADING THE DATA FILES
	// df1 := GetFishBoats()  			     					//Country,Year,Total,Non-Powered Boat,Powered-Boat
	df2 := fish_consumption.GetDataframe() //Country,Code,Year,Fish
	// df3 := fish_employment.GetDataframe() 					//Country,Year,Fisheries
	df4 := fish_prices.GetDataframe() //Country, Code, Year, Price
	// df5 := webscraper.GetDataframe()      					//Species,Status,Year,Region
	df6 := fish_catches.GetDataframe() //Country,Code,Year,Production (metric tons),Captures (metric tons)

	//DROPPING UNWANTED COLUMNS
	// df1 = df1.Drop(3) // Non-Powered
	// df1 = df1.Drop(4) // Powered
	df2 = df2.Drop(1) // Code
	df4 = df4.Drop(1) // Code
	df6 = df6.Drop(1) // Code

	//RENAMING COLUMNS
	df2 = df2.Rename("Fish-Consumption", "Fish")
	df6 = df6.Rename("Production", "Production (metric tons)")
	df6 = df6.Rename("Captures", "Captures (metric tons)")
	// df1 = df1.Rename("Total-Boats", "Total")

	//JOINING DATAFRAMES HORIZONTALLY 
	// df1 = df1.InnerJoin(df2, "Country", "Year")
	// df1 = df1.InnerJoin(df3, "Year", "Country")
	// df1 = df1.InnerJoin(df4, "Year")
	// df1 = df1.InnerJoin(df5, "Year")
	df4 = df4.InnerJoin(df2, "Country", "Year")
	df4 = df4.InnerJoin(df6, "Country", "Year")
	
	//ANALYZING
	fmt.Println(df4)
	fmt.Println(df4.Dims())
	fmt.Println(df4.Names())
	fmt.Println(df4.Describe())

	//WRITING TO CSV
	f, err := os.Create("./data/merging/merged.csv")
	if err != nil {
		log.Fatal(err)
	}
	df4.WriteCSV(f)

}

// UTILITY FUNCTIONS
func GetFishBoats() dataframe.DataFrame {
	df1 := fish_boats.GetDataframe("2008")
	df2 := fish_boats.GetDataframe("2009")
	df3 := fish_boats.GetDataframe("2010")
	df4 := fish_boats.GetDataframe("2012")
	df5 := fish_boats.GetDataframe("2013")
	df6 := fish_boats.GetDataframe("2014")
	df7 := fish_boats.GetDataframe("2015")
	fish_boats_df := df1.Concat(df2)
	fish_boats_df = df3.Concat(fish_boats_df)
	fish_boats_df = df4.Concat(fish_boats_df)
	fish_boats_df = df5.Concat(fish_boats_df)
	fish_boats_df = df6.Concat(fish_boats_df)
	fish_boats_df = df7.Concat(fish_boats_df)
	return fish_boats_df
}

func GetFishConsID() dataframe.DataFrame {
	df := fish_consumption.GetDataframe()
	IDdf := df.Filter(dataframe.F{0, "Country", "==", "Indonesia"})
	return IDdf
}
