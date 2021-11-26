package webscraper

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-gota/gota/dataframe"
	"github.com/gocolly/colly"
)

type Fishes struct {
	Data []Fish
}

type Fish struct {
	Species string `json:"species"`
	Status  string `json:"status"`
	Year    string `json:"year"`
	Region  string `json:"region"`
}

// USING GO COLLY TO WEBSCRAPE DATA
func Webscraper() {
	// REGULAR EXPRESSION TO CLEAN THE DATA
	space := regexp.MustCompile(`\s+`)

	fishStruct := make([]Fish, 0)

	collector := colly.NewCollector(
		colly.AllowedDomains("fisheries.noaa.gov", "www.fisheries.noaa.gov"),
	)

	collector.OnHTML("div.species-directory__species--8col", func(element *colly.HTMLElement) {
		species := element.DOM
		fish := Fish{
			Species: species.Find("div.species-directory__species-title--name").Text(),
			Status:  space.ReplaceAllString(strings.TrimSpace(species.Find("div.species-directory__species-status-row").Find("div.species-directory__species-status").Text()), " "),
			Year:    space.ReplaceAllString(strings.TrimSpace(species.Find("div.species-directory__species-status-row").Find("div.species-directory__species-year").Text()), " "),
			Region:  space.ReplaceAllString(strings.TrimSpace(species.Find("div.species-directory__species-status-row").Find("div.species-directory__species-region").Text()), ""),
		}
		fishStruct = append(fishStruct, fish)
	})

	collector.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting", request.URL.String())
	})

	for x := 1; x <= 5; x++ {
		collector.Visit("https://www.fisheries.noaa.gov/species-directory/threatened-endangered?title=&species_category=any&species_status=any&regions=all&items_per_page=25&page=" + strconv.Itoa(x) + "&sort=")
	}

	writeJSON(fishStruct)
}

func writeJSON(data []Fish) {
	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println("Unable to create JSON file")
		return
	}

	_ = ioutil.WriteFile("webscraper/endangeredFish.json", file, 0644)
}

func ReadJSON() Fishes {
	var fishies Fishes
	file, err := os.Open("./webscraper/endangeredFish.json")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	byteValue, _ := ioutil.ReadAll(file)
	json.Unmarshal(byteValue, &fishies.Data)
	return fishies
}

func GetDataframe() dataframe.DataFrame {
	file, err := os.Open("./webscraper/endangeredFish.json")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	df := dataframe.ReadJSON(file)
	return df
}

func generatePieItems(key string) []opts.PieData {
	items := make([]opts.PieData, 0)

	freq := map[string]int{}

	data := ReadJSON()

	if key == "Region" {
		for _, v := range data.Data {
			_, exists := freq[v.Region]

			if exists {
				freq[v.Region] += 1
			} else {
				freq[v.Region] = 1
			}
		}
		for k, v := range freq {
			if v > 3 {
				items = append(items, opts.PieData{Name: k, Value: v})
			}
		}
	}

	if key == "Year" {
		for _, v := range data.Data {
			_, exists := freq[v.Year]

			if exists {
				freq[v.Year] += 1
			} else {
				freq[v.Year] = 1
			}
		}
		for k, v := range freq {
			if len(k) < 5 {
				items = append(items, opts.PieData{Name: k, Value: v})
			}
		}
	}

	return items
}

func pieBase(key string) *charts.Pie {
	pie := charts.NewPie()
	pie.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: key + " Distribution of Endangered Fish"}),
	)

	pie.AddSeries("pie", generatePieItems(key)).
		SetSeriesOptions(charts.WithLabelOpts(
			opts.Label{
				Show:      true,
				Formatter: "{b}: {c}",
			}),
		)
	return pie
}

func GeneratePie() {
	page := components.NewPage()
	page.AddCharts(
		pieBase("Region"),
		pieBase("Year"),
	)
	http.HandleFunc("/regions", func(rw http.ResponseWriter, r *http.Request) {
		page.Render(rw)
	})
	http.ListenAndServe(":8081", nil)
}
