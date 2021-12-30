package fish_catches

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-gota/gota/dataframe"
)

func GetDataframe() dataframe.DataFrame {
	f, err := os.Open("./data/fish_catches/captures.csv")

	if err != nil {
		log.Fatal(err)
	}

	df := dataframe.ReadCSV(f)

	return df
}

// GO ECHART FUNCTIONS

// OPENS AND READS CSV FILE AND RETURNS AS 2D SLICE DEPENDING ON YEAR YOU PASS IN
func GetData(Country string) [][]string {
	f, err := os.Open("./data/fish_catches/captures.csv")
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
		if row[0] == Country {
			data = append(data, row)
		}
	}
	return data
}

func FishCatchesOverTime(Country string) {
	bar := charts.NewBar()
	bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    "Total Fish Catches Over Time in " + Country,
		Subtitle: "Catches in metric tons",
	}),
		charts.WithColorsOpts(opts.Colors{opts.HSLColor(168, 80, 20)}),
		charts.WithDataZoomOpts(opts.DataZoom{}),
	)
	bar.SetXAxis(getX(GetData(Country))).AddSeries("Values", generateBarItems(GetData(Country)))

	f, _ := os.Create("fishcatches.html")

	bar.Render(f)
}

func getX(data [][]string) []int {
	var years []int
	for _, val := range data {
		year, _ := strconv.Atoi(val[2])
		years = append(years, year)
	}
	return years
}

func generateBarItems(data [][]string) []opts.BarData {
	items := make([]opts.BarData, 0)

	total := []int{}

	for _, v := range data {
		catchesFloat, _ := strconv.ParseFloat(v[4], 64)
		catches := int(catchesFloat)
		total = append(total, catches)
	}

	for _, v := range total {
		fmt.Println(v)
		items = append(items, opts.BarData{Value: v})
	}
	return items
}

//https://ourworldindata.org/fish-and-overfishing#global-fish-production
