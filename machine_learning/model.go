package machine_learning

// https://github.com/sjwhitworth/golearn/blob/master/linear_models/linear_regression_test.go

import (
	"fmt"
	"log"

	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/evaluation"
	linear "github.com/sjwhitworth/golearn/linear_models"
)

func LinearRegression() {

	// GETTING DATA FROM CSV FILE
	rawData, err := base.ParseCSVToInstances("./data/cleaning/cleaned.csv", true)
	if err != nil {
		log.Fatal(err)
	}

	// TRAIN TEST DATA
	trainData, testData := base.InstancesTrainTestSplit(rawData, 0.01)

	model := linear.NewLinearRegression()
	//FITTING
	err = model.Fit(trainData)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("TEST DATA!")
	fmt.Println(testData)

	//PREDICT
	predictions, err := model.Predict(testData)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Attributes of Predictions \n", predictions.AllAttributes())
	fmt.Println("Attributes of Test \n", testData.AllAttributes())

	analyse, err := evaluation.GetConfusionMatrix(testData, predictions)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(evaluation.GetSummary(analyse))

}
