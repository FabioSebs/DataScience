package fish_catches

import (
	"log"
	"os"

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

//https://ourworldindata.org/fish-and-overfishing#global-fish-production
