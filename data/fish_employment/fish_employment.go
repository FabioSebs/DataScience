package fish_employment

import (
	"encoding/csv"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-gota/gota/dataframe"
)

func GetEmploymentData(country string) [][]string {
	f, err := os.Open("./data/fish_employment/employed-fisheries-aquaculture.csv")
	if err != nil {
		log.Fatal(err)
	}
	reader := csv.NewReader(f)
	data := [][]string{}
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
	// fmt.Println(data)
	return FilterSlice(data, country)
}

func GetDataframe() dataframe.DataFrame {
	f, err := os.Open("./data/fish_employment/employed-fisheries-aquaculture.csv")
	if err != nil {
		log.Fatal(err)
	}
	df := dataframe.ReadCSV(f)
	return df
}

func FilterSlice(datset [][]string, country string) [][]string {
	filtered := [][]string{}
	for _, v := range datset {
		if v[0] == country {
			filtered = append(filtered, v)
		}
	}
	return filtered
}

func EmploymentOverTime(country string, additional ...string) {
	if len(additional) == 0 {
		bar := charts.NewBar()
		bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
			Title: "The Rise in Employment -" + country,
		}))
		//Setting Instance of Bar
		bar.SetXAxis(getX(GetEmploymentData(country))).AddSeries("Values", generateBarItems(GetEmploymentData(country)))

		e, _ := os.Create("foodConsumption" + country + ".html")
		bar.Render(e)
		return
	} else {
		countries := []string{country}
		bars := []*charts.Bar{}
		for _, v := range additional {
			countries = append(countries, v)
		}
		for i := 0; i < len(countries); i++ {
			bar := charts.NewBar()
			bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
				Title: "The Rise in Employment -" + countries[i],
			}))
			//Setting Instance of Bar
			bar.SetXAxis(getX(GetEmploymentData(country))).AddSeries("Values", generateBarItems(GetEmploymentData(countries[i])))

			bars = append(bars, bar)
		}
		http.HandleFunc("/employment", func(rw http.ResponseWriter, r *http.Request) {
			for _, v := range bars {
				v.Render(rw)
			}
		})
		http.ListenAndServe(":8081", nil)
		return
	}

}

func generateBarItems(data [][]string) []opts.BarData {
	items := make([]opts.BarData, 0)

	for _, v := range data {
		items = append(items, opts.BarData{Value: v[2]})
	}

	return items
}

func getX(data [][]string) []string {
	var years []string

	for _, v := range data {
		years = append(years, v[1])
	}

	return years
}
