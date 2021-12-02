package cleaning

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/rocketlaunchr/dataframe-go/imports"
)

func Cleaning(){
	ctx := context.TODO()

	f, err := os.Open("./data/merging/merged.csv")

	if err != nil {
		log.Fatal(err)
	}
	
	df, err := imports.LoadFromCSV(ctx, f)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(df)
}

func getNullIndexes(csvFile *File)[]bool {
	arrbool := []bool{true}
	return arrbool
}
