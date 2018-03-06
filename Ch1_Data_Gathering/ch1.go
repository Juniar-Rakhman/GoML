package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"strconv"

	"github.com/kniren/gota/dataframe"
)

type CSVRecord struct {
	SepalLength float64
	SepalWidth  float64
	PetalLength float64
	PetalWidth  float64
	Species     string
	ParseError  error
}

func main() {

	file, err := os.Open(path.Join(path.Dir(os.Args[0]), "../data/iris.cvs"))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 0

	var csvRecord CSVRecord
	var csvData []CSVRecord

	for {
		record, err := reader.Read()
		if err == io.EOF {
			log.Println("reached EOF")
			break
		}

		for i, value := range record {

			var floatVal float64

			if i == 4 {
				if value == "" {
					log.Printf("Unexpected type in column %d\n", i)
					csvRecord.ParseError = fmt.Errorf("Empty string value")
					break
				}

				csvRecord.Species = value
			}

			if floatVal, err = strconv.ParseFloat(value, 64); err != nil {
				log.Printf("Unexpected type in column %d\n", i)
				csvRecord.ParseError = fmt.Errorf("could not parse float : %s", value)
				break
			}

			switch i {
			case 0:
				csvRecord.SepalLength = floatVal
			case 1:
				csvRecord.SepalWidth = floatVal
			case 2:
				csvRecord.PetalLength = floatVal
			case 3:
				csvRecord.PetalWidth = floatVal
			}
		}

		if csvRecord.ParseError == nil {
			csvData = append(csvData, csvRecord)
		}
	}

	csvDF := dataframe.LoadStructs(csvData)
	irisDF := dataframe.ReadCSV(file)

	fmt.Println(csvDF)

	fmt.Println(irisDF)

	//versiColorDF := irisDF.Filter(dataframe.F{
	//	Colname:    "species",
	//	Comparator: "==",
	//	Comparando: "Iris-versicolor",
	//}).Select()
}
