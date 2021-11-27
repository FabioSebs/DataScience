package fish_prices

import (
	"encoding/csv"
	"io"
	"log"
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

// STRUCT THAT WILL BE USED FOR DATAFRAME
type FishPrices struct {
	Country     string
	CountryCode string
	Year        string
}

// OPENS AND READS CSV AND RETURNS 2D SLICE
func GetData() [][]string {
	f, err := os.Open("./data/fish_prices/PSALMUSDA.csv")
	data := [][]string{}

	if err != nil {
		log.Fatal(err)
	}

	reader := csv.NewReader(f)

	for {
		row, err := reader.Read()

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		data = append(data, row)
	}
	return data
}

// OPENS AND READS CSV AND RETURNS DATAFRAME
func GetDataframe() dataframe.DataFrame {
	// CHALLENGE : OUR CSV FILE HAS MANY YEARS AS COLUMN NAMES
	// LETS TRY TO MAKE IT ONLY ONE SINGLE COLUMN

	f, err := os.Open("./data/fish_prices/global_fish_prices.csv")

	if err != nil {
		log.Fatal(err)
	}

	years := []string{}

	reader := csv.NewReader(f)
	reader.LazyQuotes = true

	//READING ONCE ONLY GETTING ALL THE COLUMN NAMES
	row, err := reader.Read()

	if err != nil {
		log.Fatal(err)
	}

	//SLICING THE SLICE FOR ONLY THE YEARS
	row = row[4:]
	years = append(years, row...)

	// MAKING SLICE OF STRUCTS
	structList := []FishPrices{}

	for {
		row, err := reader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}
		// USES STRUCT TO PUT ALL THE YEARS INTO ONE SINGLE COLUMN WITHIN NESTED LOOP
		for x := 0; x < len(years); x++ {
			structList = append(structList, FishPrices{row[0], row[1], years[x]})
		}
	}

	// NOW THE DATAFRAME HAS ALL THE YEARS IN ONE SINGLE COLUMN BUT MISSING PRICE
	df := dataframe.LoadStructs(structList)

	// WE OPEN THE CSV FILE AGAIN AND REPEAT THE PROCESS THIS TIME FOR PRICES
	f2, err := os.Open("./data/fish_prices/global_fish_prices.csv")

	if err != nil {
		log.Fatal(err)
	}

	reader2 := csv.NewReader(f2)
	reader2.LazyQuotes = true

	prices := []string{}

	for {
		row, err := reader2.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		row = row[4:]
		// A SLICE OF ALL THE PRICES + UNWANTED YEARS AS THE FIRST 62 ELEMENTS IN THE SLICE
		prices = append(prices, row...)
	}

	// Series( slice, type, column name )
	// WE SLICE AWAY THE YEARS BY USING LEN(YEARS) AND TURN IT INTO A SERIES
	price_series := series.New(prices[len(years):], series.String, "Price")

	// WE ADD THE COLUMN
	df = df.Mutate(price_series)

	df = df.Rename("Code", "CountryCode")

	// WRITE IT AS THE CLEANED CSV FILE
	f3, err := os.Create("./data/fish_prices/clean_fishprices.csv")

	if err != nil {
		log.Fatal(err)
	}

	df.WriteCSV(f3)

	return df
}

// GO E CHART FUNCTIONS

func generateLineItems(data [][]string) []opts.LineData {
	items := make([]opts.LineData, 0)
	for i := range data {
		items = append(items, opts.LineData{Value: data[i][1]})
	}
	return items
}

func GenerateFishPrice() {
	years := []string{}

	for i := 1; i <= 30; i++ {
		years = append(years, GetData()[i][0])
	}

	line := charts.NewLine()

	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "Price of Fresh Fishes", Subtitle: "1990-2020"}),
	)

	line.SetXAxis(years).
		AddSeries("Prices", generateLineItems(GetData())).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: true}))

	f, _ := os.Create("fishprices.html")

	line.Render(f)
}

//SRC : https://data.worldbank.org/indicator/ER.FSH.CAPT.MT
