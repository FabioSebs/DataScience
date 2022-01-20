package fish_prices

import (
	"encoding/csv"
	"io"
	"log"
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-gota/gota/dataframe"
)

// OPENS AND READS CSV AND RETURNS 2D SLICE
func GetData() [][]string {
	f, err := os.Open("./data/fish_prices/prices.csv")
	x := 0
	data := [][]string{}

	if err != nil {
		log.Fatal(err)
	}

	reader := csv.NewReader(f)
	reader.LazyQuotes = true

	for {
		row, err := reader.Read()

		// Just to skip the column names
		if x == 0 {
			x++
			continue
		}

		// End of File Error
		if err == io.EOF {
			break
		}

		// Error Handle
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

	f, err := os.Open("./data/fish_prices/prices.csv")

	if err != nil {
		log.Fatal(err)
	}

	df := dataframe.ReadCSV(f)

	df.Filter()
	return df
}

// GO E CHART FUNCTIONS

func generateLineItems(data [][]string) []opts.LineData {
	items := make([]opts.LineData, 0)
	for _, val := range data {
		items = append(items, opts.LineData{Value: val[1]})
	}
	return items
}

func getYears(data [][]string) []string {
	years := []string{}
	for _, val := range data {
		years = append(years, val[0])
	}
	return years
}

func GenerateFishPrice() {
	years := getYears(GetData())
	line := charts.NewLine()

	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "Prices of Fresh Fishes in the World", Subtitle: "1990-2021"}),
	)

	line.SetXAxis(years).
		AddSeries("Captures", generateLineItems(GetData())).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: true}))

	f, _ := os.Create("fishprices.html")

	line.Render(f)
}

//SRC : https://fred.stlouisfed.org/series/PSALMUSDQ
