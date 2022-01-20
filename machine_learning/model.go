package machine_learning

// https://github.com/sjwhitworth/golearn/blob/master/linear_models/linear_regression_test.go

import (
	"fmt"

	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/trees"
)

const (
	MAE string = "mae"
	MSE string = "mse"
)

func DecisionTree() {
	// Importing Data
	data, err := base.ParseCSVToInstances("data/cleaning/cleaned.csv", true)
	if err != nil {
		panic(err)
	}

	//Printing Data
	fmt.Println(data)

	//Model
	model := trees.NewDecisionTreeRegressor("mae", 4)

	// Training Testing Split
	trainData, testData := base.InstancesTrainTestSplit(data, 0.50)
	model.Fit(trainData)

	predictions := model.Predict(testData)

	fmt.Println(predictions)

}
