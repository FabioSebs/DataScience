package fish_employment

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func GetData() {
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
	fmt.Println(data)
}
