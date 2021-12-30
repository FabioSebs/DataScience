package machine_learning

// https://github.com/sjwhitworth/golearn/blob/master/linear_models/linear_regression_test.go

import (
	"encoding/csv"
	"fmt"
	"image/color"
	"log"
	"math"
	"os"
	"strconv"

	"github.com/go-gota/gota/dataframe"
	"github.com/sajari/regression"
	"github.com/sjwhitworth/golearn/base"
	linear "github.com/sjwhitworth/golearn/linear_models"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func LinearRegression() {

	// GETTING DATA FROM CSV FILE
	rawData, err := base.ParseCSVToInstances("./data/cleaning/cleaned.csv", true)
	if err != nil {
		log.Fatal(err)
	}

	// TRAIN TEST DATA
	trainData, testData := base.InstancesTrainTestSplit(rawData, 0.70)

	fmt.Println("TRAIN DATA!")
	fmt.Println(trainData)

	// Model
	model := linear.NewLinearRegression()
	//FITTING
	err = model.Fit(trainData)

	if err != nil {
		log.Fatal(err)
	}

	//PREDICT
	predictions, err := model.Predict(testData)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Attributes of Predictions \n", predictions.AllAttributes())
	fmt.Println("Attributes of Test \n", testData.AllAttributes())

	fmt.Printf("predictions \n%v", predictions)
	// analyse, err := evaluation.GetConfusionMatrix(testData, predictions)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(evaluation.GetSummary(analyse))
}

func SajariRegression() {
	//Opening CSV FILE
	f, err := os.Open("./data/cleaning/training.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	reader := csv.NewReader(f)

	reader.FieldsPerRecord = 4

	trainingData, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	var r regression.Regression
	r.SetObserved("Consumption")
	r.SetVar(0, "Year")

	for idx, record := range trainingData {
		if idx == 0 {
			continue
		}

		yVal, err := strconv.ParseFloat(record[3], 64)
		if err != nil {
			log.Fatal(err)
		}

		year, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			log.Fatal(err)
		}

		r.Train(regression.DataPoint(yVal, []float64{year}))
	}
	r.Run()
	fmt.Printf("\nRegression Formula:\n%v\n", r.Formula)

	f, err = os.Open("./data/cleaning/testing.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	reader = csv.NewReader(f)

	reader.FieldsPerRecord = 4

	testData, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	var mAE float64

	for idx, record := range testData {
		if idx == 0 {
			continue
		}
		yObserved, err := strconv.ParseFloat(record[3], 64)
		if err != nil {
			log.Fatal(err)
		}

		year, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			log.Fatal(err)
		}

		yPredicted, err := r.Predict([]float64{year})
		if err != nil {
			log.Fatal(err)
		}

		mAE += math.Abs(yObserved-yPredicted) / float64(len(testData))
	}
	fmt.Printf("MAE = %0.2f\n\n", mAE)
}

func PlotModel() {
	// Opening File
	f, err := os.Open("./data/cleaning/cleaned.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	df := dataframe.ReadCSV(f)

	yVals := df.Col("Fish-Consumption").Float()
	pts := make(plotter.XYs, df.Nrow())
	ptsPred := make(plotter.XYs, df.Nrow())

	for idx, val := range df.Col("Year").Float() {
		pts[idx].X = val
		pts[idx].Y = yVals[idx]
		ptsPred[idx].X = val
		ptsPred[idx].Y = predict(val)
	}

	p := plot.New()

	p.X.Label.Text = "Year"
	p.Y.Label.Text = "Consumption"
	p.Add(plotter.NewGrid())

	s, err := plotter.NewScatter(pts)
	if err != nil {
		log.Fatal(err)
	}
	s.GlyphStyle.Color = color.RGBA{R: 255, B: 128, A: 255}
	s.GlyphStyle.Radius = vg.Points(3)

	l, err := plotter.NewLine(ptsPred)
	if err != nil {
		log.Fatal(err)
	}
	l.LineStyle.Width = vg.Points(1)
	l.LineStyle.Dashes = []vg.Length{vg.Points(5), vg.Points(5)}
	l.LineStyle.Color = color.RGBA{B: 255, A: 255}

	p.Add(s, l)
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "regression.png"); err != nil {
		log.Fatal(err)
	}

}

func predict(year float64) float64 {
	return year*0.1413 - 264.3149
}
