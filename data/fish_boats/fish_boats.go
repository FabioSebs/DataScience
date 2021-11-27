package fish_boats

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"sort"
	"strconv"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-gota/gota/dataframe"
)

// OPENS AND READS CSV FILE AND RETURNS AS 2D SLICE DEPENDING ON YEAR YOU PASS IN
func GetData(year string) [][]string {
	f, err := os.Open("./data/fish_boats/" + year + ".csv")
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

// OPENS CSV AND RETURNS IT AS DATAFRAME
func GetDataframe(year string) dataframe.DataFrame {
	f, err := os.Open("./data/fish_boats/" + year + ".csv")

	if err != nil {
		log.Fatal(err)
	}

	df := dataframe.ReadCSV(f)
	return df
}

// LOOKS THROUGH CSV AND RETURNS TOTAL BOATS DEPENDING ON COUNTRY STRING YOU PASS IN
func GetTotalBoats(country string, data [][]string) int {
	for _, v := range data {
		if v[0] == country {
			res, _ := strconv.Atoi(v[2])
			return res
		}
	}
	return 0
}

// GETS THE DATA OF ALL YEARS FROM ALL THE CSVS IN LOCAL DIRECTORY
func GetAllFiles() map[int][][]string {
	years := []int{2008, 2009, 2010, 2012, 2013, 2014, 2015}
	data := map[int][][]string{}
	for _, v := range years {
		data[v] = GetData(strconv.Itoa(v))
	}
	return data
}

// GO ECHART FUNCTIONS
func FishBoatsOverTime() {
	bar := charts.NewBar()
	bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    "Total Fish Boats Over Time in Indonesia",
		Subtitle: "Non Powered and Powered Boat",
	}),
		charts.WithColorsOpts(opts.Colors{opts.HSLColor(168, 50, 40)}),
		charts.WithDataZoomOpts(opts.DataZoom{}),
	)

	bar.SetXAxis(getX(GetAllFiles())).
		AddSeries("Values", generateBarItems(GetAllFiles()))

	f, _ := os.Create("fishboats.html")

	bar.Render(f)
}

func getX(data map[int][][]string) []int {
	var years []int
	for k, _ := range data {
		years = append(years, k)
	}
	sort.Ints(years)
	return years
}

func generateBarItems(data map[int][][]string) []opts.BarData {
	items := make([]opts.BarData, 0)

	total := []int{}

	for _, v := range data {
		boats, err := strconv.Atoi(v[1][2])

		if err != nil {
			log.Fatal(err)
		}

		total = append(total, boats)
	}

	sort.Ints(total)

	for _, v := range total {
		items = append(items, opts.BarData{Value: v})
	}
	return items
}
