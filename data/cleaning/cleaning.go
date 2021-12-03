package cleaning

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/go-gota/gota/dataframe"
	"github.com/rocketlaunchr/dataframe-go/imports"
)

func Cleaning() {
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

	//CHECKING OTHER COLUMNS
	check := []bool{}
	oneCol := getNullIndexes("Country")
	twoCol := getNullIndexes("Year")
	threeCol := getNullIndexes("Price")
	fourCol := getNullIndexes("Fish-Consumption")
	sixCol := getNullIndexes("Captures")

	for _, v := range oneCol {
		if v == true {
			check = append(check, v)
		}
	}
	for _, v := range twoCol {
		if v == true {
			check = append(check, v)
		}
	}
	for _, v := range threeCol {
		if v == true {
			check = append(check, v)
		}
	}
	for _, v := range fourCol {
		if v == true {
			check = append(check, v)
		}
	}
	for _, v := range sixCol {
		if v == true {
			check = append(check, v)
		}
	}

	fmt.Println(check)

	//FILL NULL IDXS WITH 0
	for _, v := range idxs {
		df.UpdateRow(v, nil, map[string]interface{}{
			"Production": 0,
		})
	}

	fmt.Println("DataFrame")
	fmt.Println(df)
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
