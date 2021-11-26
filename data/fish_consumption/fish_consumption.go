package fish_consumption

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

// OPENS AND READS CSV AND USES UTILITY FUNCTION TO RETURN DATA BASED ON YEAR PASSED IN TO FUNCTION
func GetCountriesFC(country string) [][]string {
	//Reading CSV File
	f, err := os.Open("./data/fish_consumption/fishconsumption.csv")
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
	return FilterSlice(data, country)
}

// OPENS AND READS CSV AND RETURNS DATAFRAME
func GetDataframe() dataframe.DataFrame {
	f, err := os.Open("./data/fish_consumption/fishconsumption.csv")

	if err != nil {
		log.Fatal(err)
	}

	df := dataframe.ReadCSV(f)

	return df
}

// UTILITY FUNCTION THAT RETURNS 2D SLICE BASED ON COUNTRY PASSED IN
func FilterSlice(slice [][]string, country string) [][]string {
	filtered := [][]string{}
	for i, v := range slice {
		if v[0] == country {
			filtered = append(filtered, slice[i])
		}
	}
	return filtered
}

// GO E CHART FUNCTIONS
func ConsumptionOverTime(country string, additional ...string) {
	if len(additional) == 0 {
		bar := charts.NewBar()
		bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
			Title:    "The Rise in Seafood Consumption in " + country,
			Subtitle: "Food supply quantity (kg/capita/yr)",
		}))
		//Setting Instance of Bar
		bar.SetXAxis(getX(GetCountriesFC(country))).AddSeries("Values", generateBarItems(GetCountriesFC(country)))

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
				Title:    "The Rise in Seafood Consumption in " + countries[i],
				Subtitle: "Food supply quantity (kg/capita/yr)",
			}))
			//Setting Instance of Bar
			bar.SetXAxis(getX(GetCountriesFC(country))).AddSeries("Values", generateBarItems(GetCountriesFC(countries[i])))

			bars = append(bars, bar)
		}
		http.HandleFunc("/foodcons", func(rw http.ResponseWriter, r *http.Request) {
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
		items = append(items, opts.BarData{Value: v[3]})
	}

	return items
}

func getX(data [][]string) []string {
	var years []string

	for _, v := range data {
		years = append(years, v[2])
	}

	return years
}
