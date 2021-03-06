package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"

	"gonum.org/v1/gonum/stat"
)

func main() {
	// Open the continuous observations and predictions.
	f, err := os.Open("../data/continuous_data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Create a new CSV reader reading from the opened file.
	reader := csv.NewReader(f)

	// observed and predicted will hold the parsed observed and predicted values
	// form the continuous data file.
	var observed []float64
	var predicted []float64

	// line will track row numbers for logging.
	line := 1

	// Read in the records looking for unexpected types in the columns.
	for {

		// Read in a row. Check if we are at the end of the file.
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		// Skip the header.
		if line == 1 {
			line++
			continue
		}

		// Read in the observed and predicted values.
		observedVal, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			log.Printf("Parsing line %d failed, unexpected type\n", line)
			continue
		}

		predictedVal, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			log.Printf("Parsing line %d failed, unexpected type\n", line)
			continue
		}

		// Append the record to our slice, if it has the expected type.
		observed = append(observed, observedVal)
		predicted = append(predicted, predictedVal)
		line++
	}

	// Calculate the mean absolute error and mean squared error.
	var mAE float64
	var mSE float64
	for idx, oVal := range observed {
		mAE += math.Abs(oVal-predicted[idx]) / float64(len(observed))
		mSE += math.Pow(oVal-predicted[idx], 2) / float64(len(observed))
	}

	// Output the MAE and MSE value to standard out.
	fmt.Printf("\nMSE = %0.2f\n\n", mSE) //average of the squares of all the errors
	fmt.Printf("\nMAE = %0.2f\n", mAE)   //average of the absolute values of all the errors

	// Calculate the R^2 value.
	rSquared := stat.RSquaredFrom(observed, predicted, nil) // the proportion of the variance in the observed values that we capture in the predicted values

	// Output the R^2 value to standard out.
	fmt.Printf("\nR^2 = %0.2f\n\n", rSquared)

}
