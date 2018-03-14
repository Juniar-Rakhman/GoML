package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gonum/stat"
	"github.com/kniren/gota/dataframe"
)

func main() {
	// Open the CSV file.
	irisFile, err := os.Open("/home/jrakhman/go/src/github.com/Juniar-Rakhman/GoML/data/iris.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer irisFile.Close()

	// Create a dataframe from the CSV file.
	irisDF := dataframe.ReadCSV(irisFile)

	// Get the float values from the "sepal_length" column as
	// we will be looking at the measures for this variable.
	sepalLength := irisDF.Col("sepal_length").Float()

	// Calculate the Mean of the variable.
	meanVal := stat.Mean(sepalLength, nil)

	// Calculate the Mode of the variable.
	modeVal, modeCount := stat.Mode(sepalLength, nil)

	// Calculate the Median of the variable.
	//medianVal, err := stats.Median(sepalLength)

	if err != nil {
		log.Fatal(err)
	}
	// Output the results to standard out.
	fmt.Printf("\nSepal Length Summary Statistics:\n")
	fmt.Printf("Mean value: %0.2f\n", meanVal)
	fmt.Printf("Mode value: %0.2f\n", modeVal)
	fmt.Printf("Mode count: %d\n", int(modeCount))
	//fmt.Printf("Median value: %0.2f\n\n", medianVal)
}
