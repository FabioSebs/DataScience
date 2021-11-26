package fish_boats

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-gota/gota/dataframe"
)

// OPENS AND READS CSV FILE AND RETURNS AS 2D SLICE
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

func GetDataframe(year string) dataframe.DataFrame {
	f, err := os.Open("./data/fish_boats/" + year + ".csv")

	if err != nil {
		log.Fatal(err)
	}

	df := dataframe.ReadCSV(f)
	return df
}

func GetTotalBoats(country string, data [][]string) int {
	for _, v := range data {
		if v[0] == country {
			res, _ := strconv.Atoi(v[2])
			return res
		}
	}
	return 0
}

func GetAllFiles() map[int][][]string {
	years := []int{2008, 2009, 2010, 2012, 2013, 2014, 2015}
	data := map[int][][]string{}
	for _, v := range years {
		data[v] = GetData(strconv.Itoa(v))
	}
	return data
}

func FishBoatsOverTime() {
	bar := charts.NewBar()
	bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    "Total Fish Boats Over Time in Indonesia",
		Subtitle: "Non Powered and Powered Boat",
	}))

	bar.SetXAxis(getX(GetAllFiles())).AddSeries("Values", generateBarItems(GetAllFiles()))

	http.HandleFunc("/fishboats", func(rw http.ResponseWriter, r *http.Request) {
		bar.Render(rw)
	})
	http.ListenAndServe(":8081", nil)
	return
}

func getX(data map[int][][]string) []int {
	var years []int
	for k, _ := range data {
		fmt.Println(k)
		years = append(years, k)
	}
	return years
}

func generateBarItems(data map[int][][]string) []opts.BarData {
	items := make([]opts.BarData, 0)

	for _, v := range data {
		items = append(items, opts.BarData{Value: v[1][2]})
	}
	fmt.Println(items)
	return items
}
