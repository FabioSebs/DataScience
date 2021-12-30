package cleaning

import (
	"context"
	"log"
	"os"

	"github.com/go-gota/gota/dataframe"
	df "github.com/rocketlaunchr/dataframe-go"
	"github.com/rocketlaunchr/dataframe-go/exports"
	"github.com/rocketlaunchr/dataframe-go/imports"
)

func Cleaning() *df.DataFrame {
	ctx := context.TODO()
	// OPENING FILE
	f, err := os.Open("./data/merging/merged.csv")
	if err != nil {
		log.Fatal(err)
	}
	// GETTING CSV INTO DATAFRAME
	df, err := imports.LoadFromCSV(ctx, f)
	if err != nil {
		log.Fatal(err)
	}

	// GETTING NULL INDEXES
	nullIdx := getNullIndexes("Production")
	idxs := []int{}

	for i, v := range nullIdx {
		if v == true {
			idxs = append(idxs, i)
		}
	}

	//FILL NULL IDXS WITH 0
	for _, v := range idxs {
		df.UpdateRow(v, nil, map[string]interface{}{
			"Production": 0,
		})
	}

	//Dropping Country
	df.RemoveSeries("Country")
	// df.RemoveSeries("Production")
	// df.RemoveSeries("Year")

	//WRITING TO A CSV FILE
	myFile, err := os.Create("./data/cleaning/cleaned.csv")
	if err != nil {
		log.Fatal(err)
	}

	err2 := exports.ExportToCSV(ctx, myFile, df)
	if err2 != nil {
		log.Fatal(err)
	}

	return df
}

// UTILITY TOOL
func getNullIndexes(col string) []bool {
	f, err := os.Open("./data/merging/merged.csv")
	if err != nil {
		log.Fatal(err)
	}
	df := dataframe.ReadCSV(f)
	ds := df.Col(col)

	arrBool := ds.IsNaN()

	return arrBool
}
