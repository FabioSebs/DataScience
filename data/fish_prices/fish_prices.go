package fish_prices

import (
	"encoding/csv"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

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

func generateLineItems(data [][]string) []opts.LineData {
	items := make([]opts.LineData, 0)
	for i := range data {
		items = append(items, opts.LineData{Value: data[i][1]})
	}
	return items
}

func httpserver(w http.ResponseWriter, _ *http.Request) {
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

	line.Render(w)
}

func GenerateFishPrice() {
	http.HandleFunc("/fishprice", httpserver)
	http.ListenAndServe(":8081", nil)
}
